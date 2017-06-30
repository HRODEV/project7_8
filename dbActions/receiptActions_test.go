package dbActions

import (
	"github.com/HRODEV/project7_8/models"
	"testing"
)

func TestGetReceiptById(t *testing.T) {
	// Create declaration to receive
	newReceipt := models.Receipt{ImagePath: "/path/to/declaration"}
	CreateReceipt(&newReceipt, db)

	var foundReceipt models.Receipt

	GetReceiptById(newReceipt.ID, &foundReceipt, db)

	if newReceipt != foundReceipt {
		t.Errorf("the received receipt should be %+v but was %+v", newReceipt, foundReceipt)
	}
}