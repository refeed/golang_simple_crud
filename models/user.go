package models

import (
	"context"
	"golangSimpleCrud/auth"
	"golangSimpleCrud/contracts"
	"golangSimpleCrud/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Role         string `json:"role,omitempty" bson:"role,omitempty"` // TODO: Use enum instead of string
	Username     string `json:"username,omitempty" bson:"_id,omitempty"`
	PasswordHash string `json:"passwordHash,omitempty" bson:"passwordHash,omitempty"`
}

func GetUserCollection() *mongo.Collection {
	return db.GetDB().Collection("users")
}

func GetAllUsers() []contracts.GetUserRes {
	ctx := context.TODO()
	// TODO: We might want only do the passwordHash filtering in the controllers
	//	     instead of here, just like controllers.Users.GetOne() does
	opts := options.Find().SetProjection(bson.M{"passwordHash": 0})
	cur, err := GetUserCollection().Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Println(err.Error())
	}
	defer cur.Close(ctx)

	users := []contracts.GetUserRes{}
	for cur.Next(ctx) {
		var user contracts.GetUserRes

		err := cur.Decode(&user)
		if err != nil {
			log.Println(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func GetUserById(id string) (User, error) {
	var user User
	ctx := context.TODO()
	err := GetUserCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func DeleteUserById(id string) (bool, error) {
	ctx := context.TODO()
	res, err := GetUserCollection().DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount == 0 {
		return false, err
	}
	return true, err
}

func CreateUser(userForm contracts.CreateUserReq) error {
	ctx := context.TODO()
	user := User{
		Name:         userForm.Name,
		Role:         userForm.Role,
		Username:     userForm.Username,
		PasswordHash: auth.HashPassword(userForm.Password),
	}
	_, err := GetUserCollection().InsertOne(ctx, user)
	return err
}

func UpdateUser(id string, updateUserForm contracts.UpdateUserReq) error {
	ctx := context.TODO()

	newUser := User{
		Name:         updateUserForm.Name,
		Role:         updateUserForm.Role,
		PasswordHash: auth.HashPassword(updateUserForm.Password),
	}

	_, err := GetUserCollection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": newUser})
	return err
}

func IsUserPasswordMatch(username string, password string) bool {
	user, err := GetUserById(username)
	if err != nil {
		log.Printf("IsUserPasswordMatch(): %v", err.Error())
		return false
	}
	return user.PasswordHash == auth.HashPassword(password)
}
