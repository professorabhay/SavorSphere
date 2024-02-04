package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Name			*string					`json:"first_name" validate:"required,min=4,max=50"`
	Password		*string					`json:"Password" validate:"required,min=6"`
	Email			*string					`json:"email" validate:"email,required"`
	Token			*string					`json:"token"`
	User_type		*string					`json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token	*string					`json:"refresh_token"`
	Created_at		time.Time				`json:"created_at"`  //just like timestamps in NodeJs
	Updated_at		time.Time				`json:"updated_at"`
	User_id			string					`json:"user_id"`
}
