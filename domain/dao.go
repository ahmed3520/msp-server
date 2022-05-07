package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmed3520/msp-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"image":    user.Image,
		"password": user.Password,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}
func CreateEvent(event *Event) (*Event, *utils.RestErr) {
	eventsC := db.Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := eventsC.InsertOne(ctx, bson.M{
		"name":          event.Name,
		"image":         event.Image,
		"thumbnail":     event.Thumbnail,
		"speakersevent": event.SpeakersEvent,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert event to the database..")
		return nil, restErr
	}
	event.ID = result.InsertedID.(primitive.ObjectID)
	return event, nil

}
func CreateCommite(committe *Committe) (*Committe, *utils.RestErr) {
	committesC := db.Collection("committes")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := committesC.InsertOne(ctx, bson.M{
		"name":            committe.Name,
		"image":           committe.Image,
		"description":     committe.Description,
		"thumbnail":       committe.Thumbnail,
		"NumberOfMembers": committe.NumberOfMembers,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert event to the database..")
		return nil, restErr
	}
	committe.ID = result.InsertedID.(primitive.ObjectID)
	return committe, nil
}

func Find(email string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}
func FindEvent(id primitive.ObjectID) (*Event, *utils.RestErr) {
	fmt.Println("id event", id)
	var event Event
	eventsC := db.Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := eventsC.FindOne(ctx, bson.M{"_id": id}).Decode(&event)
	if err != nil {
		restErr := utils.NotFound("Event not found..")
		return nil, restErr
	}
	return &event, nil
}
func FindCommitte(name string) (*Committe, *utils.RestErr) {
	fmt.Println("id event", name)
	var committe Committe
	eventsC := db.Collection("committes")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := eventsC.FindOne(ctx, bson.M{"name": name}).Decode(&committe)
	if err != nil {
		restErr := utils.NotFound("Committe not found..")
		return nil, restErr
	}
	return &committe, nil
}
func GetAllCommitte() ([]primitive.M, *utils.RestErr) {
	CommittesC := db.Collection("committes")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cur, err := CommittesC.Find(ctx, bson.D{{}})

	if err != nil {
		restErr := utils.NotFound("Commites not found..")
		return nil, restErr
	}
	var results []primitive.M
	for cur.Next(ctx) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			restErr := utils.NotFound("commite not found..")
			return nil, restErr
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}
	cur.Close(context.Background())
	return results, nil

}
func GetAllEvents() ([]primitive.M, *utils.RestErr) {
	eventsC := db.Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cur, err := eventsC.Find(ctx, bson.D{{}})

	if err != nil {
		restErr := utils.NotFound("Events not found..")
		return nil, restErr
	}
	var results []primitive.M
	for cur.Next(ctx) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			restErr := utils.NotFound("Events not found..")
			return nil, restErr
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}
	cur.Close(context.Background())
	return results, nil

}
func Delete(email string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		restErr := utils.NotFound("faild to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return restErr
	}
	return nil
}
func DeleteEvent(id primitive.ObjectID) *utils.RestErr {
	eventsC := db.Collection("events")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := eventsC.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		restErr := utils.NotFound("Failed to delete..")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := utils.NotFound("Event not found..")
		return restErr
	}
	return nil

}

func Update(email string, field string, value string) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}

func UpdateCommitte(name string, field string, value string) (*Committe, *utils.RestErr) {
	committesC := db.Collection("committes")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := committesC.UpdateOne(ctx, bson.M{"name": name}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("committe not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := FindCommitte(name)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
