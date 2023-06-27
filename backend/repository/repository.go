package repository

import (
	"backend/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetMovieById(id int) (*models.Movie, error)
	MovieForEdit(id int) (*models.Movie, error)
	InsertMovie(genre int, movie *models.Movie) error
	DeleteMovie(id int) error
}
