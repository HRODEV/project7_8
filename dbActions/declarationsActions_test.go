package dbActions

import (
	"github.com/HRODEV/project7_8/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func TestGetDeclarations(t *testing.T) {
	// Create user for the declaration
	user := models.User{Email: "test@test.com", FirstName: "Test", LastName: "User", Password: "Password"}
	db.Create(&user)

	// Create declaration to receive
	newDeclaration := models.Declaration{ID: 0, Title: "declaration", TotalPrice: 10, VATPrice: 2.10, Date: "2017-07-09T12:32:00", Description: "description", ReceiptID: 0, UserID: user.ID, ProjectID: 0}
	CreateDeclaration(&newDeclaration, db)

	var declaration []models.Declaration

	GetDeclarationsForUser(user.ID, &declaration, db)

	if len(declaration) == 0 {
		t.Error("at least one declartion should be received")
	}
}

func TestGetDeclarationById(t *testing.T) {
	// Create declaration to receive
	newDeclaration := models.Declaration{ID: 0, Title: "declaration", TotalPrice: 10, VATPrice: 2.10, Date: "2017-07-09T12:32:00", Description: "description", ReceiptID: 0, UserID: 0, ProjectID: 0}
	CreateDeclaration(&newDeclaration, db)

	var foundDeclartion models.Declaration

	GetDeclarationById(newDeclaration.ID, &foundDeclartion, db)

	if newDeclaration != foundDeclartion {
		t.Errorf("the received declartion should be %+v but was %+v", newDeclaration, foundDeclartion)
	}
}

func TestCreateDeclaration(t *testing.T) {
	newDeclaration := models.Declaration{ID: 0, Title: "declaration", TotalPrice: 10, VATPrice: 2.10, Date: "2017-07-09T12:32:00", Description: "description", ReceiptID: 0, UserID: 0, ProjectID: 0}
	CreateDeclaration(&newDeclaration, db)

	var lastDeclartion models.Declaration
	db.First(&lastDeclartion, newDeclaration.ID)

	if lastDeclartion != newDeclaration {
		t.Errorf("the last declartion should be %+v but was %+v", newDeclaration, lastDeclartion)
	}
}

func TestUpdateDeclarationById(t *testing.T) {
	// Create declaration to modify
	newDeclaration := models.Declaration{ID: 0, Title: "declaration", TotalPrice: 10, VATPrice: 2.10, Date: "2017-07-09T12:32:00", Description: "description", ReceiptID: 0, UserID: 0, ProjectID: 0}
	CreateDeclaration(&newDeclaration, db)

	// Update declaration
	UpdateDeclarationById(newDeclaration.ID, &models.Declaration{Title: "updated declaration"}, db)

	var lastDeclartion models.Declaration
	db.First(&lastDeclartion, newDeclaration.ID)

	if lastDeclartion.Title != "updated declaration" {
		t.Errorf("the new declartion title should be %s but was %s", newDeclaration.Title, lastDeclartion.Title)
	}
}

func TestDeleteDeclarationById(t *testing.T) {
	// Create declaration to modify
	newDeclaration := models.Declaration{ID: 0, Title: "declaration", TotalPrice: 10, VATPrice: 2.10, Date: "2017-07-09T12:32:00", Description: "description", ReceiptID: 0, UserID: 0, ProjectID: 0}
	CreateDeclaration(&newDeclaration, db)

	// Delete declaration
	DeleteDeclarationById(newDeclaration.ID, db)

	var lastDeclartion models.Declaration
	db.First(&lastDeclartion, newDeclaration.ID)

	if lastDeclartion == newDeclaration {
		t.Error("the last declaration should be deleted")
	}
}
