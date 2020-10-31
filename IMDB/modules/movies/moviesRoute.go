package movies

import (
	"encoding/json"
	"errors"
	"fynd/IMDB/middleware"
	"fynd/IMDB/models"
	"net/http"
	"strings"

	"github.com/pquerna/ffjson/ffjson"
)

//Init method
func Init() {
	http.HandleFunc("/movies/", movies)
}

func movies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case models.DeletRequest:
		Delete(w, r)
	case models.GetRequest:
		Get(w, r)
	case models.PostRequest:
		Post(w, r)
	case models.PatchRequest:
		Patch(w, r)
	default:
		http.Error(w, errors.New("method not found").Error(), http.StatusNotFound)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	movie := models.Movie{}
	err := middleware.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&movie)
	if err != nil {
		models.Logger.Println("movie Bind : ", err)
		http.Error(w, "movie obj Bind Error", http.StatusInternalServerError)
		return
	}
	_, err = AddMovieService(movie)
	if err != nil {
		return
	}
	models.Logger.Println("movie added successfully")
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	result, err := getMovieService()
	if err != nil {
		return
	}
	b, _ := ffjson.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	models.Logger.Println("fetched movie successful")
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	err := middleware.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/movies/")
	if strings.Contains(id, "/") {
		http.Error(w, errors.New("method not found").Error(), 404)
		return
	}
	err = deleteMovieService(id)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusNoContent)
	models.Logger.Println("movie removed successfully")
	return
}

func Patch(w http.ResponseWriter, r *http.Request) {
	movie := models.UpdateMovie{}
	err := middleware.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/movies/")
	if strings.Contains(id, "/") {
		http.Error(w, errors.New("method not found").Error(), 404)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&movie)
	if err != nil {
		models.Logger.Println("movie Bind : ", err)
		http.Error(w, "movie obj Bind Error", http.StatusInternalServerError)
		return
	}

	err = patchMovieService(id, movie)
	if err != nil {
		return
	}
	models.Logger.Println("movie edit successful")
	return
}
