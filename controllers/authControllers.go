package controllers

import (
	"awesomeProject/models"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		u.Response(w, u.Message(false, "Malformed request"))
		return
	}

	resp := user.Create()
	u.Response(w, resp)
}
