package repositories

import (
	"database/sql"

	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type repositoriePostUser struct {
	db *sql.DB
}

func NewRepositoriePostUser(db *sql.DB) service.IPostUsuario {
	return &repositoriePostUser{db: db}
}

func (r repositoriePostUser) Get_PostUser_By_UserId(id uint) dto.APIRespuesta[[]*dto.APIRespuestaRegister] {
	//TODO implement me
	panic("implement me")
}

func (r repositoriePostUser) Create_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones {
	//TODO implement me
	panic("implement me")
}

func (r repositoriePostUser) Update_PostUser(obj entities.PostUsuario) *dto.APIRespuestaAcciones {
	//TODO implement me
	panic("implement me")
}

func (r repositoriePostUser) Delete_PostUser(id uint) *dto.APIRespuestaAcciones {
	//TODO implement me
	panic("implement me")
}
