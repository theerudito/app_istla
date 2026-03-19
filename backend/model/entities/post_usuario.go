package entities

type PostUsuario struct {
	PostUserId          int    `json:"post_user_id"`
	Descripcion         string `json:"descripcion"`
	UsuarioId           int    `json:"usuario_id"`
	File                []byte `json:"file"`
	UsuarioCreacion     string `json:"usuario_creacion"`
	UsuarioModificacion string `json:"usuario_modificacion"`
}
