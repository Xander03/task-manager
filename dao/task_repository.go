package dao

import (
	. "../model"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type TaskRepository struct {
	Repository Repository
}

func (r *TaskRepository) AddTask(task Task) error {
	task.ID = bson.NewObjectId()
	if err := r.Repository.getTasksCollection().Insert(task); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTasks(userId string) ([]Task, error) {
	var tasks []Task
	if !bson.IsObjectIdHex(userId) {
		return tasks, errors.New("get all tasks dao wrong user id")
	}

	oUserId := bson.ObjectIdHex(userId)

	if err := r.Repository.getTasksCollection().Find(bson.M{"user_id": oUserId}).All(&tasks); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTaskById(id string) (Task, error) {
	var task Task

	if !bson.IsObjectIdHex(id) {
		err := errors.New("get task by id dao wrong id")
		return task, err
	}
	oid := bson.ObjectIdHex(id)

	if err := r.Repository.getTasksCollection().FindId(oid).One(&task); err != nil {
		return task, err
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(task Task) error {
	if err := r.Repository.getTasksCollection().UpdateId(task.ID, task); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("delete task dao wrong id")
	}

	oid := bson.ObjectIdHex(id)

	if err := r.Repository.getTasksCollection().RemoveId(oid); err != nil {
		return err
	}

	return nil
}