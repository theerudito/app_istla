package dto

type PostUsuarioDTO struct {
	PostUserId int `json:"post_user_id"`

	Descripcion string `json:"descripcion"`

	Usuario   string `json:"usuario"`
	UsuarioId int    `json:"usuario_id"`

	StorageId int    `json:"id_storage"`
	Url       string `json:"url"`

	UsuarioCreacion     string `json:"usuario_creacion"`
	UsuarioModificacion string `json:"usuario_modificacion"`
	FechaCreacion       string `json:"fecha_creacion"`
	FechaModificacion   string `json:"fecha_modificacion"`
}
