package dao

import (
	"log"

	"github.com/yzbtdiy/uni-update-server/models"
)

// 从数据库中查找是否存在用户和密码匹配的行
func GetUserInfoFromDB(username string, password string) (user *models.DbUserTable) {
	result := db.Where("username=? and password=?", username, password).First(&user)
	if result.RowsAffected == 1 {
		return user
	} else {
		hasAdmin := db.Where("username = admin").First(&user)
		if hasAdmin.RowsAffected == 0 {
			log.Println("User admin not found, will add it")
			CreateDefaultUser()
		}
		return nil
	}
}

func CreateDefaultUser() {
	addAdmin := models.DbUserTable{
		Id:       1,
		Username: "admin",
		Password: "admin",
	}
	result := db.Create(&addAdmin)
	if result.RowsAffected != 0 {
		log.Println("Add default user admin with password admin")
	}
}
