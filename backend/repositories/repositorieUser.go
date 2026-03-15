package repositories

import (
	"database/sql"

	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type repositorieUser struct {
	db *sql.DB
}

func NewRepositorieUser(db *sql.DB) service.IUser {
	return &repositorieUser{db: db}
}

func (r repositorieUser) Login(obj dto.UsuarioLoginDTO) dto.APIRespuesta[*dto.APIRespuestaLogin] {
	//TODO implement me
	panic("implement me")
}

func (r repositorieUser) Register(obj entities.Usuario) dto.APIRespuesta[*dto.APIRespuestaRegister] {
	//TODO implement me
	panic("implement me")
}
