package model

// User is representation of ACCOUNTING.User_Ref
type User struct {
	UserID                             int
	PwdHash, Name, SecondName, Surname string
}
