package modules

import (
	"fynd/IMDB/modules/administration"
	"fynd/IMDB/modules/movies"
)

//Init ..
func Init() {
	administration.Init()
	movies.Init()
}
