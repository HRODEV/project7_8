package models

type Receipt struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	ImagePath string
	Data      string `gorm:"type:text"`
}

func (receipt *Receipt) IsValid() bool {
	return true
}
