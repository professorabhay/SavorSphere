package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required,max=20"`
	Email    string             `json:"email" binding:"required,email"`
	Password string             `json:"password" binding:"required,min=6"`
}
