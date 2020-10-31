package main

import (
	"fmt"
	"fynd/IMDB/models"
	"fynd/IMDB/modules"

	// "fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

func main() {
	models.Logger.Println("Server Started")
	err := setconfig()
	if err != nil {
		panic(err)
	}
	modules.Init()
	fmt.Println("starting server on localhost:", models.Config.AppPort)
	models.Logger.Println("starting server on localhost:", models.Config.AppPort)
	err = http.ListenAndServe(":"+models.Config.AppPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//setconfig - reads config file and assigns to config model
func setconfig() error {
	_, err := toml.DecodeFile(models.GetConfigFilePath(), &models.Config)
	if err != nil {
		models.Logger.Println("err", err)
		return err
	}

	return nil
}
