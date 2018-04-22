package model

import "gopkg.in/mgo.v2/bson"

type Task struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	UserID bson.ObjectId `bson:"user_id" json:"user_id"`
	Description string `bson:"desc" json:"desc"`
	IsDone bool `bson:"is_done" json:"is_done"`
}

type Tasks []Task
