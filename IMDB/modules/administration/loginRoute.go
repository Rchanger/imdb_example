package administration

import (
	"encoding/json"
	"errors"
	"fynd/IMDB/models"
	"fynd/IMDB/modules/common"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

//Init method
func Init() {
	http.HandleFunc("/auth/login", login)
}

//login
func login(w http.ResponseWriter, r *http.Request) {
	models.Logger.Println("login route")
	if r.Method != models.PostRequest {
		http.Error(w, errors.New("method not found").Error(), 404)
		return
	}
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		models.Logger.Println("login Bind : ", err)
		http.Error(w, "login obj Bind Error", 500)
		return
	}
	result, status, _ := loginService(user)
	if !status {
		models.Logger.Println("error occured while fetching token from for use ", result)
		http.Error(w, "Authentication Failed", http.StatusUnauthorized)
		return
	}
	t, err := common.GenerateToken(result)
	if err != nil {
		models.Logger.Println("ValidateCredentialsRoute GenerateToken : ", err)
		http.Error(w, "Generattion of Authentication Token Failed", http.StatusInternalServerError)
		return
	}
	result.Password = ""
	result.Role = ""
	w.Header().Set("authorization", t)
	b, _ := ffjson.Marshal(result)
	w.Write(b)
	return
}
