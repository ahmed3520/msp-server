package services

import (
	"github.com/ahmed3520/msp-server/domain"
	"github.com/ahmed3520/msp-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func CreateCommite(committe *domain.Committe) (*domain.Committe, *utils.RestErr) {
	committe, restErr := domain.CreateCommite(committe)
	if restErr != nil {
		return nil, restErr
	}
	return committe, nil
}

func FindUser(email string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Find(email)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}
func FindEvent(id primitive.ObjectID) (*domain.Event, *utils.RestErr) {
	event, restErr := domain.FindEvent(id)
	if restErr != nil {
		return nil, restErr
	}
	return event, nil
}
func FindCommitte(name string) (*domain.Committe, *utils.RestErr) {
	committe, restErr := domain.FindCommitte(name)
	if restErr != nil {
		return nil, restErr
	}
	return committe, nil
}
func GetAllEvents() ([]primitive.M, *utils.RestErr) {
	events, restErr := domain.GetAllEvents()
	if restErr != nil {
		return nil, restErr
	}
	return events, nil
}
func GetAllCommitte() ([]primitive.M, *utils.RestErr) {
	committe, restErr := domain.GetAllCommitte()
	if restErr != nil {
		return nil, restErr
	}
	return committe, nil
}
func DeleteUser(email string) *utils.RestErr {
	restErr := domain.Delete(email)
	if restErr != nil {
		return restErr
	}
	return nil
}
func DeleteEvent(id primitive.ObjectID) *utils.RestErr {
	restErr := domain.DeleteEvent(id)
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

func UpdateCommitte(name string, field string, value string) (*domain.Committe, *utils.RestErr) {
	committe, restErr := domain.UpdateCommitte(name, field, value)
	if restErr != nil {
		return nil, restErr
	}
	return committe, nil
}
