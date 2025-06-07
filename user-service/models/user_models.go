package models

import "time"

type RegisterUserRequest struct {
	Name     string `json:"name" bson:"name"`
	Mobile   string `json:"mobile" bson:"mobile"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type RegistertUserInsertion struct {
	Name      string    `json:"name" bson:"name"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type RegisterUserResponse struct {
	UserId  string `json:"userId" bson:"userId"`
	Message string `json:"message"`
}
