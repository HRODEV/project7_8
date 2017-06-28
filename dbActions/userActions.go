package dbActions

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/HRODEV/project7_8/models"
)

func GetUserByID(id int , db *gorm.DB) (user models.User){
	db.Find(&user);
	return
}