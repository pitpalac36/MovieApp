package main

import (
	"backend/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("authenticate")
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:password`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// dummy token
	app.writeJSON(w, http.StatusAccepted, "jwt")
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) Movie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	movie, err := app.DB.GetMovieById(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movie)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) MovieForEdit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	movie, err := app.DB.MovieForEdit(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var payload = struct {
		Movie *models.Movie `json:"movie"`
	}{
		movie,
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AddMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addMovie")
	var requestPayload struct {
		ID          int    `json:id`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		RunTime     string `json:"runtime"`
		Image       string `json:"image"`
		MPAARating  string `json:"mpaa_rating"`
		Description string `json:"description"`
		Genre       string `json:"genre"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	runtime, err := strconv.Atoi(requestPayload.RunTime)
	if err != nil {
		app.errorJSON(w, errors.New("invalid runtime value"), http.StatusBadRequest)
		return
	}

	movie := models.Movie{
		Title:       requestPayload.Title,
		ReleaseDate: requestPayload.ReleaseDate,
		RunTime:     runtime,
		Image:       requestPayload.Image,
		MPAARating:  requestPayload.MPAARating,
		Description: requestPayload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	genre_id, err := strconv.Atoi(requestPayload.Genre)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = app.DB.InsertMovie(genre_id, &movie)
	if err != nil {
		app.errorJSON(w, errors.New("error when adding movie"), http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusCreated, movie)
}

func (app *application) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.DB.DeleteMovie(id)
	if err != nil {
		app.errorJSON(w, errors.New("error when deleting movie"), http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, id)
}
