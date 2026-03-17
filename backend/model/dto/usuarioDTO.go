package dto

type UsuarioDTO struct {
	UsuarioId      int    `json:"id_usuario"`
	Identificacion string `json:"identificacion"`
	Nombres        string `json:"nombres"`
	Apellidos      string `json:"apellidos"`
	Email          string `json:"email"`
	PerfilId       int    `json:"id_perfil"`
	Perfil         string `json:"perfil"`
}

type UsuarioLoginDTO struct {
	Identificacion string `json:"identificacion"`
	Password       string `json:"password"`
}
