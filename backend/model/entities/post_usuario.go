package entities

import "time"

type PostUsuario struct {
	PostUserId          uint      `json:"post_user_id"`
	Descripcion         string    `json:"descripcion"`
	UsuarioId           string    `json:"usuario_id"`
	StorageId           string    `json:"id_storage"`
	UsuarioCreacion     string    `json:"usuario_creacion"`
	UsuarioModificacion string    `json:"usuario_modificacion"`
	FechaCreacion       time.Time `json:"fecha_creacion"`
	FechaModificacion   time.Time `json:"fecha_modificacion"`
}
