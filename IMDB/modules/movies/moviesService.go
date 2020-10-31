package movies

import (
	"fynd/IMDB/models"
	"fynd/IMDB/modules/datastore"
)

// getMovieService gives you movies object.
func getMovieService() ([]models.Movie, error) {
	return datastore.GetAllMovies()

}

func AddMovieService(movie models.Movie) (string, error) {
	return datastore.AddMovie(movie)

}

func deleteMovieService(id string) error {
	return datastore.DeleteMovie(id)

}

func patchMovieService(id string, movie models.UpdateMovie) error {
	return datastore.PatchMovie(id, movie)

}
