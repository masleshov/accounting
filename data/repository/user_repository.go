package repository

import (
	"accounting/accounting/data/model"
)

// UserRepository represents type which can manage data related with users
type UserRepository struct {
	*Repository
}

// NewUserRepository creates new instance of UserRepository type
func NewUserRepository() *UserRepository {
	return &UserRepository{
		Repository: NewRepository(),
	}
}

// GetUsers returns all users from database
func (ur *UserRepository) GetUsers() DBResult {
	users := []model.User{}

	var rows, err = ur.db.ExecSelect(selectAllUsersQuery(), nil)
	if rows != nil && err == nil {
		for rows.Next() {
			user := new(model.User)
			err = rows.Scan(&user.UserID, &user.PwdHash, &user.Name, &user.SecondName, &user.Surname)
			if err != nil {
				break
			}

			users = append(users, *user)
		}
	}

	return newDBResult(users, err)
}

// InsertUser inserts user to database and changes their UserId
func (ur *UserRepository) InsertUser(user *model.User) DBResult {
	args := []interface{}{user.PwdHash, user.Name, user.SecondName, user.Surname}
	row, err := ur.db.ExecCRUD(insertUserQuery(), args)
	if err == nil {
		row.Scan(&user.UserID)
	}

	return newDBResult(user, err)
}

// UpdateUser updates user in database
func (ur *UserRepository) UpdateUser(user *model.User) DBResult {
	args := []interface{}{user.UserID, user.PwdHash, user.Name, user.SecondName, user.Surname}
	_, err := ur.db.ExecCRUD(updateUserQuery(), args)

	return newDBResult(user, err)
}

// DeleteUser deletes user from database
func (ur *UserRepository) DeleteUser(user *model.User) DBResult {
	args := []interface{}{user.UserID}
	_, err := ur.db.ExecCRUD(deleteUserQuery(), args)

	return newDBResult(user, err)
}

func selectAllUsersQuery() string {
	return "Select User_ID " +
		", Pwd_Hash " +
		", Name " +
		", Second_Name " +
		", Surname " +
		"From ACCOUNTING.User_Ref"
}

func insertUserQuery() string {
	return "Insert into ACCOUNTING.User_Ref(Pwd_Hash, Name, Second_Name, Surname) " +
		"Values($1, $2, $3, $4) " +
		"Returning User_ID"
}

func updateUserQuery() string {
	return "Update ACCOUNTING.User_Ref " +
		"Set Pwd_Hash = $2 " +
		"  , Name = $3 " +
		"  , Second_Name = $4 " +
		"  , Surname = $5 " +
		"Where User_ID = $1"
}

func deleteUserQuery() string {
	return "Delete From ACCOUNTING.User_Ref Where User_ID = $1"
}
