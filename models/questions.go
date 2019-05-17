package models

import (
	u "awesomeProject/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Question string `json:"question"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
}

func (question *Question) Validate() (map[string]interface{}, bool) {

	if question.Question == "" {
		return u.Message(false, "Empty question not allowed!"), false
	}

	if question.UserID == 0 {
		return u.Message(false, "User ID Missing!"), false
	}

	return u.Message(true, "Question validated!"), true
}

func (question *Question) Create() map[string]interface{} {

	if resp, ok := question.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(question).Error

	if err != nil {
		fmt.Println("Error ", question.ID)
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "Success")
	resp["question"] = question
	return resp
}

func GetQuestionsByUserId(user_id uint) []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Where("user_id = ?", user_id).Find(&questions).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return questions
}

func GetAllQuestions() []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Find(&questions).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return questions
}

func GetAllUnansweredQuestions() []*Question {
	questions := make([]*Question, 0)
	err := GetDB().Table("questions").Joins("left join answers on answers.question_id = questions.id").
		Where("answers.id IS NULL").Find(&questions).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return questions
}
