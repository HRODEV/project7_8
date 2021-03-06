package models

import "errors"

type User struct {
	ID                  uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Email               string `gorm:"unique_index"`
	FirstName           string
	LastName            string
	Password            string
	DeclarationStatuses *[]DeclarationStatus
	UserProjects        *[]Project
	Declarations        *[]Declaration
}

func (u *User) IsValid() (bool, error) {
	if u == nil {
		return false, errors.New("Something went really wrong...")
	}

	if len(u.Email) < 5 {
		return false, errors.New("Email address is not valid")
	}

	if len(u.Password) < 8 {
		return false, errors.New("Password must be at least 8 characters")
	}

	return true, nil
}
