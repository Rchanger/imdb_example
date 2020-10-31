package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

//Customer ... for customer profile
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	EmailID  string `json:"emailID" bson:"emailID"`
	Role     string `json:"role" bson:"role"`
}

//Customer ... for customer profile
type Movie struct {
	Popularity float64       `json:"popularity" bson:"99popularity"`
	Director   string        `json:"director" bson:"director"`
	Name       string        `json:"name" bson:"name"`
	Score      float64       `json:"score" bson:"imdb_score"`
	Genre      []string      `json:"genre" bson:"genre"`
	Id         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
}

type UpdateMovie struct {
	Popularity *float64 `json:"popularity,omitempty" bson:"99popularity,omitempty"`
	Score      *float64 `json:"score,omitempty" bson:"imdb_score,omitempty"`
}

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	LoginID string `json:"loginId,omitempty"`
	Role    string `json:"role,omitempty"`
	jwt.StandardClaims
}
