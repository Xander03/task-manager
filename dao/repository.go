package dao

import "gopkg.in/mgo.v2"

const SERVER = "localhost:27017"
const DB_NAME = "taskmanager"
const USERS_COLLECTION = "users"
const TASKS_COLLECTION = "tasks"

type Repository struct{}

func (r *Repository) getUsersCollection() *mgo.Collection {
	session := createSession()
	return session.DB(DB_NAME).C(USERS_COLLECTION)
}

func (r *Repository) getTasksCollection() *mgo.Collection {
	session := createSession()
	return session.DB(DB_NAME).C(TASKS_COLLECTION)
}

func createSession() *mgo.Session {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		panic("Couldn't create db connection")
	}
	return session
}