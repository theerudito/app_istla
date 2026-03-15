package dto

type PostUsuarioDTO struct {
	PostUserId uint `json:"post_user_id"`

	Usuario   string `json:"usuario"`
	UsuarioId string `json:"usuario_id"`

	StorageId string `json:"id_storage"`
	Url       string `json:"url"`

	UsuarioCreacion     string `json:"usuario_creacion"`
	UsuarioModificacion string `json:"usuario_modificacion"`
	FechaCreacion       string `json:"fecha_creacion"`
	FechaModificacion   string `json:"fecha_modificacion"`
}
