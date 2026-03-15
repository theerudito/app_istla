package dto

type LogAccionDTO struct {
	LodAccionId uint   `json:"id_accion_log"`
	Accion      string `json:"accion"`
	TablaNombre string `json:"tabla_nombre"`
	Descripcion string `json:"descripcion"`
	RegistroId  string `json:"id_registro"`
	Fecha       string `json:"fecha"`
}
