package dto

type LogErroDTO struct {
	LogErroId   uint   `json:"logErroId"`
	Mensaje     string `json:"mensaje"`
	TablaNombre string `json:"tabla_nombre"`
	Fecha       string `json:"fecha"`
}
