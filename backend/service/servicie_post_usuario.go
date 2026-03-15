package service

import (
	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
)

type IPostUsuario interface {
	Get_PostUser_By_UserId(id uint) dto.APIRespuesta[[]*dto.APIRespuestaRegister]
	Create_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones
	Update_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones
	Delete_PostUser(id uint) *dto.APIRespuestaAcciones
}
