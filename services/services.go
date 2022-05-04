package services

import (
	"github.com/ahmed3520/msp-server/domain"
	"github.com/ahmed3520/msp-server/utils"
)

func CreateUser(user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}
func CreateEvent(event *domain.Event) (*domain.Event, *utils.RestErr) {
	event, restErr := domain.CreateEvent(event)
	if restErr != nil {
		return nil, restErr
	}
	return event, nil
}

func FindUser(email string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Find(email)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func DeleteUser(email string) *utils.RestErr {
	restErr := domain.Delete(email)
	if restErr != nil {
		return restErr
	}
	return nil
}

func UpdateUser(email string, field string, value string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Update(email, field, value)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}
