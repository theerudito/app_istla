package service

import (
	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
)

type IUser interface {
	Login(obj dto.UsuarioLoginDTO) dto.APIRespuesta[*dto.APIRespuestaLogin]
	Register(obj entities.Usuario) dto.APIRespuesta[*dto.APIRespuestaRegister]
}
