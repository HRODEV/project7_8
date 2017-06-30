package dbActions

import (
	"errors"
	"github.com/HRODEV/project7_8/models"
	"github.com/jinzhu/gorm"
)

func GetReceiptById(id uint, receipt *models.Receipt, db *gorm.DB) {
	db.First(&receipt, id)
}

func CreateReceipt(receipt *models.Receipt, db *gorm.DB) error {
	if receipt.IsValid() {
		db.Create(&receipt)
		return nil
	} else {
		return errors.New("Receipt Struct not valid")
	}
}
