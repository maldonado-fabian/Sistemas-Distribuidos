package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastname"`
	Rut      string             `json:"rut"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
