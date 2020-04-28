package service

import (
	"accounting/accounting/data/model"
	"accounting/accounting/data/repository"
	"accounting/accounting/util"
	"encoding/json"
	"errors"
	"regexp"
)

const pwdHashPattern = "([A-Z0-9]{8}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{12})"

var userRepository repository.UserRepository

func init() {
	userRepository = *repository.NewUserRepository()
}

// GetUsers returns all users from database
func GetUsers(request GetRequest) util.JSONObject {
	request.ValidateBodyIsEmpty()
	return toJSON(userRepository.GetUsers())
}

// CreateUser creates new user and returns him with assigned UserId
func CreateUser(request PostRequest) util.JSONObject {
	user := &model.User{}
	json.Unmarshal(request.Body, user)

	if err := validateUser(*user); err != nil {
		return errorToJSON(err)
	}

	return toJSON(userRepository.InsertUser(user))
}

// UpdateUser updates some properties of User
func UpdateUser(request PutRequest) util.JSONObject {
	user := &model.User{}
	json.Unmarshal(request.Body, user)

	return toJSON(userRepository.UpdateUser(user))
}

// DeleteUser deletes user from database
func DeleteUser(request DeleteRequest) util.JSONObject {
	user := &model.User{}
	json.Unmarshal(request.Body, user)

	return toJSON(userRepository.DeleteUser(user))
}

func validateUser(user model.User) error {
	if match, _ := regexp.MatchString("(?i)"+pwdHashPattern, user.PwdHash); !match {
		return errors.New("Password hash must match pattern " + pwdHashPattern)
	}

	return nil
}
