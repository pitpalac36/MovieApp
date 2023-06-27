package repository

import (
	"backend/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			m.id, genre_id, genre, title, release_date, runtime,
			mpaa_rating, description, coalesce(image, ''),
			m.created_at, m.updated_at
		from
			movies m 
			inner join movies_genres mg 
				on m.id = mg.movie_id 
			inner join genres g
				on mg.genre_id = g.id
		order by
			title
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		var genre models.Genre
		err := rows.Scan(
			&movie.ID,
			&genre.ID,
			&genre.Genre,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movie.Genre = &genre
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from users where email = $1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *PostgresDBRepo) GetMovieById(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from movies where id = $1`

	var movie models.Movie
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.ReleaseDate,
		&movie.RunTime,
		&movie.MPAARating,
		&movie.Description,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	query = `select g.id, g.genre from movies_genres mg
		left join genres g on (mg.genre_id = g.id)
		where mg.movie_id = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, err
		}
		movie.Genre = &g
	}
	return &movie, err
}

func (m *PostgresDBRepo) MovieForEdit(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from movies where id = $1`

	var movie models.Movie
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.ReleaseDate,
		&movie.RunTime,
		&movie.MPAARating,
		&movie.Description,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	query = `select g.id, g.genre from movies_genres mg
		left join genres g on (mg.genre_id = g.id)
		where mg.movie_id = $1
		order by g.genre`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, err
		}
		movie.Genre = &g
	}
	return &movie, err
}

func (m *PostgresDBRepo) InsertMovie(genre_id int, movie *models.Movie) error {
	formated_date := strings.SplitAfter(movie.ReleaseDate, " ")[0]

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Add Movie
	query := `insert into movies (title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	movie_id := 0
	err := m.DB.QueryRowContext(ctx, query, movie.Title, formated_date, movie.RunTime,
		movie.MPAARating, movie.Description, movie.Image, movie.CreatedAt, movie.UpdatedAt).Scan(&movie_id)

	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	// Insert new row into movies_genres
	query = `insert into movies_genres (movie_id, genre_id)
		values ($1, $2)`

	_, err = m.DB.ExecContext(ctx, query, movie_id, genre_id)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) DeleteMovie(id int) error {
	fmt.Println("Deleted")
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `delete from movies where id = $1`
	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}
