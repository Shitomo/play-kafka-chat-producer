package model

// ファーストネーム
type FirstName string

func (f FirstName) Valid() *ValidationError {
	if f == "" {
		return NewValidationError("firstname", "empty")
	}
	return nil
}

// ラストネーム
type LastName string

func (l LastName) Valid() *ValidationError {
	if l == "" {
		return NewValidationError("lastname", "empty")
	}
	return nil
}

// 誕生日
type BirthDay Datetime

// ユーザー
type User struct {
	FirstName FirstName
	LastName  LastName
	BirthDay  BirthDay
}

// ユーザーを作成
func NewUser(firstName FirstName, lastName LastName, birthDay BirthDay) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDay:  birthDay,
	}
}
