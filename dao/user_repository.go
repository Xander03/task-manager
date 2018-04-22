package dao

import (
	. "../model"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type UserRepository struct {
	Repository Repository
}

func (r *UserRepository) AddUser(user User) error {
	user.ID = bson.NewObjectId()
	if err := r.Repository.getUsersCollection().Insert(user); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.Repository.getUsersCollection().Find(nil).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserById(id string) (User, error) {
	var user User

	if !bson.IsObjectIdHex(id) {
		return user, errors.New("get user by id dao: id wrong id")
	}

	oid := bson.ObjectIdHex(id)
	if err := r.Repository.getUsersCollection().FindId(oid).One(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user User) error {
	if err := r.Repository.getUsersCollection().UpdateId(user.ID, user); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("delete user dao wrong id")
	}

	oid := bson.ObjectIdHex(id)
	if err := r.Repository.getUsersCollection().RemoveId(oid); err != nil{
		return err
	}
	return nil
}