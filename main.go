package main

import (
	"awesomeProject/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/question/post", controllers.PostQuestion).Methods("POST")
	router.HandleFunc("/api/question/byuser", controllers.FetchQuestionsByUser).Methods("GET")
	router.HandleFunc("/api/question/all", controllers.FetchAllQuestions).Methods("GET")

	router.HandleFunc("/api/answer/post", controllers.PostAnswer).Methods("POST")
	router.HandleFunc("/api/answer/byuser", controllers.GetAnswerByUserID).Methods("GET")
	router.HandleFunc("/api/answer/byquestion", controllers.FetchAnswersByQuestionID).Methods("GET")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)

	if err != nil {
		fmt.Println(err)
	}

}
