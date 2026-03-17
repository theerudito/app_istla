package repositories

import (
	"database/sql"
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/theerudito/istla/helpers"
	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type repositoriePostUser struct {
	db *sql.DB
}

func NewRepositoriePostUser(db *sql.DB) service.IPostUsuario {
	return &repositoriePostUser{db: db}
}

func (r repositoriePostUser) Get_PostUser_By_UserId(id uint) dto.APIRespuesta[[]*dto.PostUsuarioDTO] {
	var posts []*dto.PostUsuarioDTO

	rows, err := r.db.Query(`
	SELECT
		pu.id_post_usuario,
		pu.descripcion,
		pu.usuario_creacion,
		pu.usuario_modificacion,
		pu.fecha_creacion,
		pu.fecha_modificacion,
		COALESCE(u.id_usuario, 0) AS usuario_id,
		COALESCE(u.nombres || ' ' || u.apellidos, '') AS usuario,
		COALESCE(s.id_storage, 0) AS storage,
		COALESCE(s.url, '') AS url
	FROM post_usuario AS pu
	LEFT JOIN usuarios AS u ON pu.id_usuario = u.id_usuario
	LEFT JOIN storage AS s ON pu.id_storage = s.id_storage
	WHERE 
		pu.id_usuario = $1
	ORDER BY 
    	pu.id_post_usuario`, id)

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post user", "error ejecutando la consulta "+err.Error())
		return dto.APIRespuesta[[]*dto.PostUsuarioDTO]{Codigo: 500, Mensaje: "error ejecutando la consulta", Resultado: []*dto.PostUsuarioDTO{}}
	}

	defer rows.Close()

	for rows.Next() {

		post := &dto.PostUsuarioDTO{}

		err := rows.Scan(
			&post.PostUserId,
			&post.Descripcion,
			&post.UsuarioCreacion,
			&post.UsuarioModificacion,
			&post.FechaCreacion,
			&post.FechaModificacion,
			&post.UsuarioId,
			&post.Usuario,
			&post.StorageId,
			&post.Url)

		if err != nil {
			return dto.APIRespuesta[[]*dto.PostUsuarioDTO]{Codigo: 500, Mensaje: "error leyendo los registros", Resultado: []*dto.PostUsuarioDTO{}}
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return dto.APIRespuesta[[]*dto.PostUsuarioDTO]{Codigo: 404, Mensaje: "no se encontraron registros", Resultado: []*dto.PostUsuarioDTO{}}
	}

	return dto.APIRespuesta[[]*dto.PostUsuarioDTO]{Codigo: 200, Mensaje: "OK", Resultado: posts}

}

func (r repositoriePostUser) Create_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones {
	var (
		err                error
		tx                 *sql.Tx
		url, extension     string
		uuidImage          = uuid.New().String()
		postId, id, existe int
		storageId          *int
	)

	tx, err = r.db.Begin()

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post user", "error iniciando la transacion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error iniciando la transacion"}
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	qUser := `SELECT COUNT(*) FROM usuarios WHERE id_usuario = $1`
	err = r.db.QueryRow(qUser, obj.UsuarioId).Scan(&existe)
	if existe == 0 {
		return &dto.APIRespuestaAcciones{Codigo: 404, Mensaje: "el usuario no existe"}
	}

	extension, err = helpers.ExtencionFile(obj.File)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "storage", "error obteniendo extension del archivo "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error obteniendo extension del archivo"}
	}

	url, err = helpers.SaveImageToDirectory(obj.File, uuidImage, extension, os.Getenv("PDF"))
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "storage", "error guardando el archivo "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error guardando el archivo"}
	}

	objStorage := dto.StorageItemDTO{
		TX:        tx,
		FileName:  uuidImage,
		Extension: extension,
		Url:       url,
		Option:    "INSERT",
	}

	id, err = helpers.StorageManager(objStorage)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "storage", err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: err.Error()}
	}

	storageId = &id

	err = tx.QueryRow(`
	INSERT INTO post_usuario (
	descripcion, id_usuario, id_storage, usuario_creacion, usuario_modificacion) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id_post_usuario`,
		strings.ToUpper(obj.Descripcion),
		obj.UsuarioId,
		storageId,
		obj.UsuarioCreacion,
		obj.UsuarioCreacion).Scan(&postId)

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error insertando el registro "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error insertando el registro"}
	}

	err = helpers.InsertLogs(r.db, "INSERT", "post usuario", postId, "registro creado correctamente")
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "logs", "error al insertar la auditoria "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error al insertar la auditoria"}
	}

	err = tx.Commit()

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error al finalizar la transacion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error al finalizar la transacion"}
	}

	return &dto.APIRespuestaAcciones{Codigo: 200, Mensaje: "registro creado correctamente"}

}

func (r repositoriePostUser) Update_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones {

	var (
		err            error
		tx             *sql.Tx
		existeUser     int
		storageId      int
		fileName       string
		url, extension string
		uuidImage      = uuid.New().String()
	)

	tx, err = r.db.Begin()
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error iniciando transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error iniciando transaccion"}
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	queryPost := `SELECT id_storage FROM post_usuario WHERE id_post_usuario = $1`
	err = tx.QueryRow(queryPost, obj.PostUserId).Scan(&storageId)

	if errors.Is(err, sql.ErrNoRows) {
		return &dto.APIRespuestaAcciones{Codigo: 404, Mensaje: "el post no existe"}
	}

	if err != nil {
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error consultando el post"}
	}

	queryUser := `SELECT COUNT(*) FROM usuarios WHERE id_usuario = $1`
	err = tx.QueryRow(queryUser, obj.UsuarioId).Scan(&existeUser)

	if existeUser == 0 {
		return &dto.APIRespuestaAcciones{Codigo: 404, Mensaje: "el usuario no existe"}
	}

	if obj.File != nil {

		queryStorage := `SELECT nombre || extencion AS oldFile FROM storage WHERE id_storage = $1`
		err = tx.QueryRow(queryStorage, storageId).Scan(&fileName)

		extension, err = helpers.ExtencionFile(obj.File)
		if err != nil {
			return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error obteniendo extension"}
		}

		url, err = helpers.SaveImageToDirectory(obj.File, uuidImage, extension, os.Getenv("PDF"))
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", "error guardando archivo "+err.Error())
			return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error guardando archivo"}
		}

		objStorage := dto.StorageItemDTO{
			TX:        tx,
			StorageId: storageId,
			FileName:  uuidImage,
			Extension: extension,
			Url:       url,
			Option:    "UPDATE",
		}

		_, err = helpers.StorageManager(objStorage)
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", err.Error())
			return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: err.Error()}
		}
	}

	update := `
	UPDATE post_usuario
	SET
		descripcion 			= $1,
		id_usuario 				= $2,
		usuario_modificacion 	= $3,
		fecha_modificacion 		= NOW()
	WHERE id_post_usuario 		= $4`

	_, err = tx.Exec(
		update,
		strings.ToUpper(obj.Descripcion),
		obj.UsuarioId,
		obj.UsuarioCreacion,
		obj.PostUserId,
	)

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error actualizando "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error actualizando registro"}
	}

	err = helpers.InsertLogs(r.db, "UPDATE", "post usuario", obj.PostUserId, "registro actualizado")
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error guardando auditoria "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error guardando auditoria"}
	}

	err = tx.Commit()

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error finalizando transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error finalizando transaccion"}
	}

	if fileName != "" {
		path := os.Getenv("PDF") + "/" + fileName
		err = helpers.DeleteImageFromDirectory(path)
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", "error borrando imagen anterior "+err.Error())
			return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error borrando imagen anterior"}
		}
	}

	return &dto.APIRespuestaAcciones{Codigo: 200, Mensaje: "registro actualizado correctamente"}
}

func (r repositoriePostUser) Delete_PostUser(id uint) *dto.APIRespuestaAcciones {

	var (
		err       error
		tx        *sql.Tx
		storageId sql.NullInt64
		fileName  string
	)

	// iniciar transacción
	tx, err = r.db.Begin()
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error iniciando transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error iniciando transaccion"}
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// 1️⃣ Obtener el id_storage del post (si tiene)
	query := `SELECT id_storage FROM post_usuario WHERE id_post_usuario = $1`
	err = tx.QueryRow(query, id).Scan(&storageId)

	if errors.Is(err, sql.ErrNoRows) {
		return &dto.APIRespuestaAcciones{Codigo: 404, Mensaje: "el post no existe"}
	}

	if err != nil {
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error consultando el post"}
	}

	// 2️⃣ Borrar el post primero
	deletePostQuery := `DELETE FROM post_usuario WHERE id_post_usuario = $1`
	_, err = tx.Exec(deletePostQuery, id)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error eliminando el post "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error eliminando el registro"}
	}

	// 3️⃣ Registrar auditoría
	err = helpers.InsertLogs(r.db, "DELETE", "post usuario", int(id), "registro eliminado")
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error guardando auditoria "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error guardando auditoria"}
	}

	// 4️⃣ Si existía storage, eliminarlo en BD y obtener el nombre del archivo
	if storageId.Valid {

		queryStorage := `SELECT nombre || extencion AS oldFile FROM storage WHERE id_storage = $1`
		err = tx.QueryRow(queryStorage, storageId.Int64).Scan(&fileName)
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", "error consultando storage "+err.Error())
		}

		objStorage := dto.StorageItemDTO{
			TX:        tx,
			StorageId: int(storageId.Int64),
			Option:    "DELETE",
		}

		_, err = helpers.StorageManager(objStorage)
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", "error eliminando el storage "+err.Error())
		}
	}

	// 5️⃣ Commit de la transacción
	err = tx.Commit()
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "post usuario", "error finalizando transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error finalizando transaccion"}
	}

	// 6️⃣ Borrar el archivo físico después del commit
	if fileName != "" {
		path := os.Getenv("PDF") + "/" + fileName
		err = helpers.DeleteImageFromDirectory(path)
		if err != nil {
			_ = helpers.InsertLogsError(r.db, "storage", "error borrando la imagen "+err.Error())
		}
	}

	return &dto.APIRespuestaAcciones{Codigo: 200, Mensaje: "registro eliminado correctamente"}
}
