package entities

type Usuario struct {
	UsuarioId      uint   `json:"id_usuario"`
	Identificacion string `json:"identificacion"`
	Nombres        string `json:"nombres"`
	Apellidos      string `json:"apellidos"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PerfilId       uint   `json:"id_perfil"`
}
