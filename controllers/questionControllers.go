package controllers

import (
	"awesomeProject/models"
	u "awesomeProject/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PostQuestion = func(w http.ResponseWriter, r *http.Request) {
	question := &models.Question{}
	err := json.NewDecoder(r.Body).Decode(question)

	if err != nil {
		u.Response(w, u.Message(false, "Malformed Request"))
		return
	}

	resp := question.Create()
	u.Response(w, resp)
}

var FetchQuestionsByUser = func(w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()
	user_id_str, ok := vals["user_id"]

	if !ok {
		fmt.Println("Error", user_id_str)
		u.Response(w, u.Message(false, "User ID is missing"))
		return
	}

	user_id, err := strconv.Atoi(user_id_str[0])

	if err != nil {
		fmt.Println("Error", user_id)
		u.Response(w, u.Message(false, "User ID Missing"))
		return
	}

	data := models.GetQuestionsByUserId(uint(user_id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Response(w, resp)
}

var FetchAllQuestions = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetAllQuestions()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Response(w, resp)
}