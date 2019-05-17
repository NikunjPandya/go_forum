package models

import (
	u "awesomeProject/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Answers struct {
	gorm.Model
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
	UserID     uint   `json:"user_id"`
}

func (answer *Answers) Validate() (map[string]interface{}, bool) {
	if answer.Answer == "" {
		return u.Message(false, "Empty answer not allowed!"), false
	}

	if answer.QuestionID == 0 {
		return u.Message(false, "Question ID missing!"), false
	}

	if answer.UserID == 0 {
		return u.Message(false, "User ID missing!"), false
	}

	return u.Message(true, "Answer Validated"), true
}

func (answer *Answers) Create() map[string]interface{} {

	if resp, ok := answer.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(answer).Error

	if err != nil {
		fmt.Println("Error", answer.ID)
		return u.Message(false, err.Error())
	}

	resp := u.Message(true, "Success")
	resp["answer"] = answer
	return resp
}

func GetAnswersByUserID(user_id uint) []*Answers {
	answers := make([]*Answers, 0)
	err := GetDB().Table("answers").Where("user_id = ?", user_id).Find(&answers).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return answers
}

func GetAnswersByQuestionID(question_id uint) []*Answers {
	answers := make([]*Answers, 0)
	err := GetDB().Table("answers").Where("question_id = ?", question_id).Find(&answers).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return answers
}
