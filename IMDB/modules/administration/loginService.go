package administration

import (
	"fynd/IMDB/models"
	"fynd/IMDB/modules/datastore"
)

// loginService gives you person object if userid and password is correct.
func loginService(userObj models.User) (models.User, bool, string) {
	userObjFromDB, err := datastore.GetUser(userObj.Username)
	if err != nil {
		models.Logger.Println("loginService : ", err)
		return userObjFromDB, false, err.Error()
	}
	if userObjFromDB.Password != userObj.Password {
		return userObjFromDB, false, "authentication failed"
	}
	return userObjFromDB, true, "Welcome!"
}
