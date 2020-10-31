package datastore

import (
	"fmt"
	"fynd/IMDB/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(username string) (models.User, error) {
	dbStatus, err := mgo.Dial(models.Config.MongoServer)
	if err != nil {
		models.Logger.Println("Error Connecting mongo server", err)
		return models.User{}, err
	}
	cd := dbStatus.DB(models.DBNAME).C(models.COLLECTION_USERS)
	result := models.User{}
	err = cd.Find(bson.M{"username": username}).Select(bson.M{"password": 1, "role": 1, "username": 1}).One(&result)
	if err != nil {
		models.Logger.Println("Error Finding user ", err)
		return models.User{}, err
	}
	return result, nil
}

func GetAllMovies() ([]models.Movie, error) {
	dbStatus, err := mgo.Dial(models.Config.MongoServer)
	if err != nil {
		models.Logger.Println("Error Connecting mongo server", err)
		return []models.Movie{}, err
	}
	cd := dbStatus.DB(models.DBNAME).C(models.COLLECTION_MOVIES)
	result := []models.Movie{}
	err = cd.Find(bson.M{}).All(&result)
	if err != nil {
		models.Logger.Println("Error Finding movies ", err)
		return []models.Movie{}, err
	}
	return result, nil
}

func AddMovie(movie models.Movie) (string, error) {
	dbStatus, err := mgo.Dial(models.Config.MongoServer)
	if err != nil {
		models.Logger.Println("Error Connecting mongo server", err)
		return "Error Connecting mongo server", err
	}
	cd := dbStatus.DB(models.DBNAME).C(models.COLLECTION_MOVIES)
	err = cd.Insert(movie)
	if err != nil {
		models.Logger.Println("Error Adding movie: ", err)
		return "Movie adding failed", err
	}
	return "Movie added successfully", nil
}

func DeleteMovie(id string) error {
	dbStatus, err := mgo.Dial(models.Config.MongoServer)
	if err != nil {
		models.Logger.Println("Error Connecting mongo server", err)
		return err
	}
	cd := dbStatus.DB(models.DBNAME).C(models.COLLECTION_MOVIES)
	err = cd.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Println("Error Removing movie: ", err)
		models.Logger.Println("Error Removing movie: ", err)
		return err
	}
	return nil
}

func PatchMovie(id string, movie models.UpdateMovie) error {
	dbStatus, err := mgo.Dial(models.Config.MongoServer)
	if err != nil {
		models.Logger.Println("Error Connecting mongo server", err)
		return err
	}
	cd := dbStatus.DB(models.DBNAME).C(models.COLLECTION_MOVIES)
	err = cd.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": movie})
	if err != nil {
		fmt.Println("Error updating movie: ", err)
		models.Logger.Println("Error updating movie: ", err)
		return err
	}
	return nil
}
