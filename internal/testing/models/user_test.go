package models

import (
	"fmt"
	"noir.github.com/aws_billing/internal/models"
	"testing"
)

func TestUser(t *testing.T) {
	models.DB = models.CreateDB()
	var user models.User

	//user.Username = "test@highcloud.cn"
	//user.Password = "000000"
	//_, err := user.Create()
	//if err != nil {
	//	fmt.Println(err)
	//}

	user.Find("pppq@qq.com")
	fmt.Println(user.ID)
	//fmt.Println(user.Total())
	//fmt.Println(&user)

	//user.RealDelete(1)

	//fmt.Println(user.Match("test@highcloud.cn", "o1ron"))
}

func TestUserList(t *testing.T) {
	models.DB = models.CreateDB()
	list := make([]map[string]interface{}, 0)
	err := models.DB.Model(&models.User{}).Select("users.id", "users.username", "any_value(login_logs.created_at) as created_at", "users.is_admin", "users.status").Joins("left join login_logs on login_logs.username = users.username").Group("users.username").Scan(&list).Error
	fmt.Println(list, err)
}
