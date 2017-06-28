package models

type User struct {
	ID                  uint `gorm:"primary_key;AUTO_INCREMENT"`
	Email               string
	FirstName           string
	LastName            string
	Password            string
	DeclarationStatuses *[]DeclarationStatus
	UserProjects        *[]Project
	Declarations        *[]Declaration
}
