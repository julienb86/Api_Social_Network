package models

import "time"

type Post struct {
	Id        *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title     *string   `json:"title"`
	Content   *string   `json:"content"`
	UserId    *string   `json:"userid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
