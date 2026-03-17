package repositories

import (
	"database/sql"
	"strings"

	"github.com/theerudito/istla/helpers"
	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type repositorieUser struct {
	db *sql.DB
}

func NewRepositorieUser(db *sql.DB) service.IUser {
	return &repositorieUser{db: db}
}

func (r repositorieUser) Login(obj dto.UsuarioLoginDTO) *dto.APIRespuestaAcciones {

	var (
		err        error
		passwordDB string
		token      string
	)

	err = r.db.QueryRow(`
		SELECT password
		FROM usuarios
		WHERE identificacion = $1`,
		obj.Identificacion).Scan(&passwordDB)

	if err != nil {
		return &dto.APIRespuestaAcciones{Codigo: 404, Mensaje: "usuario no existe"}
	}

	passwordDB, err = helpers.DesencriptarDato(passwordDB)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error desencriptando la contraseña "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error desencriptando la contraseña"}
	}

	if passwordDB != obj.Password {
		return &dto.APIRespuestaAcciones{Codigo: 401, Mensaje: "contraseña incorrecta"}
	}

	token, err = helpers.GenerateToken(obj.Identificacion)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error generando token "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error generando token"}
	}

	return &dto.APIRespuestaAcciones{Codigo: 200, Mensaje: token}
}

func (r repositorieUser) Register(obj entities.Usuario) *dto.APIRespuestaAcciones {

	var (
		err               error
		existe, usuarioId int
		tx                *sql.Tx
		password, token   string
	)

	tx, err = r.db.Begin()
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error iniciando la transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error iniciando la transaccion"}
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = tx.QueryRow(`
		SELECT COUNT(*) 
		FROM usuarios 
		WHERE identificacion = $1`,
		strings.ToUpper(obj.Identificacion)).Scan(&existe)

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error ejecutando la consulta "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error ejecutando la consulta"}
	}

	if existe > 0 {
		return &dto.APIRespuestaAcciones{Codigo: 409, Mensaje: "ya existe un usuario registrado con la identificacion " + obj.Identificacion}
	}

	password, err = helpers.EncriptarDato(obj.Password)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error encriptando la contraseña "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error encriptando la contraseña"}
	}

	queryInsert := `
	INSERT INTO usuarios(identificacion, nombres, apellidos, email, password, id_perfil)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id_usuario;`

	err = tx.QueryRow(
		queryInsert,
		strings.ToUpper(obj.Identificacion),
		strings.ToUpper(obj.Nombres),
		strings.ToUpper(obj.Apellidos),
		strings.ToLower(obj.Email),
		password,
		obj.PerfilId,
	).Scan(&usuarioId)

	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error insertando el registro "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error insertando el registro"}
	}

	err = tx.Commit()
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error finalizando la transaccion "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error finalizando la transaccion"}
	}

	err = helpers.InsertLogs(r.db, "INSERT", "usuarios", usuarioId, "registro creado correctamente")
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error insertando auditoria "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error insertando auditoria"}
	}

	token, err = helpers.GenerateToken(obj.Identificacion)
	if err != nil {
		_ = helpers.InsertLogsError(r.db, "usuario", "error generando token "+err.Error())
		return &dto.APIRespuestaAcciones{Codigo: 500, Mensaje: "error generando token"}
	}

	return &dto.APIRespuestaAcciones{Codigo: 200, Mensaje: token}
}
