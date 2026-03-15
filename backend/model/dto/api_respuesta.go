package dto

type APIRespuesta[T any] struct {
	Codigo    int    `json:"codigo"`
	Mensaje   string `json:"mensaje"`
	Resultado T      `json:"resultado"`
}

type APIRespuestaAcciones struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type APIRespuestaLogin struct {
	UsuarioId int    `json:"usuario_id"`
	Token     string `json:"token"`
}

type APIRespuestaRegister struct {
	UsuarioId int    `json:"usuario_id"`
	Token     string `json:"token"`
}
