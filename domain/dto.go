package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Image    string             `json:"image" bson:"image,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}
type Event struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Image         string             `json:"image" bson:"image,omitempty"`
	Date          string             `json:"date" bson:"date,omitempty"`
	Location      string             `json:"location" bson:"location,omitempty"`
	Description   string             `json:"description" bson:"description,omitempty"`
	Thumbnail     string             `json:"thumbnail" bson:"thumbnail,omitempty"`
	SpeakersEvent []User             `json:"speakersevent" bson:"speakersevent,omitempty"`
}

type Committe struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	Image           string             `json:"image" bson:"image,omitempty"`
	Description     string             `json:"description" bson:"description,omitempty"`
	Thumbnail       string             `json:"thumbnail" bson:"thumbnail,omitempty"`
	NumberOfMembers string             `json:"numberofmembers" bson:"numberofmembers,omitempty"`
}
