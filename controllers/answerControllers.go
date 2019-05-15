package controllers

import (
	"awesomeProject/models"
	u "awesomeProject/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PostAnswer = func(w http.ResponseWriter, r *http.Request) {
	answer := &models.Answers{}
	err := json.NewDecoder(r.Body).Decode(answer)

	if err != nil{
		u.Response(w, u.Message(false, "Malformed Request"))
		return
	}

	resp := answer.Create()
	u.Response(w, resp)
}

var GetAnswerByUserID = func(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	user_id_str, ok := vals ["user_id"]

	if !ok || len(user_id_str) == 0{
		fmt.Println("Error", user_id_str)
		u.Response(w, u.Message(false, "User ID is Missing"))
		return
	}

	user_id, err := strconv.Atoi(user_id_str[0])

	if err != nil {
		fmt.Println("Error", user_id)
		u.Response(w, u.Message(false, "User ID Missing"))
		return
	}

	data := models.GetAnswersByUserID(uint(user_id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Response(w, resp)
}

var FetchAnswersByQuestionID = func(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	question_id_str, ok := vals ["question_id"]

	if !ok || len(question_id_str) == 0 {
		fmt.Println("Error", question_id_str)
		u.Response(w, u.Message(false, "Question ID is Missing"))
		return
	}

	question_id, err := strconv.Atoi(question_id_str[0])

	if err != nil {
		fmt.Println("Error", question_id)
		u.Response(w, u.Message(false, "Question ID Missing"))
		return
	}

	data := models.GetAnswersByQuestionID(uint(question_id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Response(w, resp)
}