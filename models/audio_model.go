package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Audio struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Description string             `json:"description,omitempty"`
	URL         string             `json:"url,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
}
