package models

import (
	"gorm.io/gorm"
	"log"
	"noir.github.com/aws_billing/internal/modules/utils"
	"time"
)

const PasswordSaltLength = 6

// User model
type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Status    Status `json:"status" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Salt      string `json:"salt"`
	IsAdmin   int    `json:"is_admin"`
	Username  string `json:"username" gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) Create() (insertId int, err error) {
	user.Status = Enabled
	user.Salt = user.generateSalt()
	user.Password = user.encryptPassword(user.Password, user.Salt)
	if err := DB.Create(user).Error; err == nil {
		insertId = user.ID
	} else {
		log.Println(err)
	}
	return
}

func (user *User) Update(id int, data CommonMap) {
	DB.Model(user).Where("id = ?", id).Updates(data)
}

func (user *User) Delete(id int) {
	DB.Model(user).Delete(user, id)
}

func (user *User) Disable(id int) {
	user.Update(id, CommonMap{"status": Disabled})
}

func (user *User) Enable(id int) {
	user.Update(id, CommonMap{"status": Enabled})
}

func (user *User) Find(username string) {
	DB.Model(user).Where("username = ?", username).First(user)
}

func (user *User) List() ([]map[string]interface{}, error) {
	//select u.id, u.username, ll.created_at from users as u left join login_logs ll on u.username = ll.username group by u.username ;
	list := make([]map[string]interface{}, 0)
	err := DB.Model(user).Select("users.id", "users.username", "any_value(login_logs.created_at) as created_at", "users.is_admin", "users.status").Joins("left join login_logs on login_logs.username = users.username").Group("users.username").Scan(&list).Error
	return list, err
}

func (user *User) Total() int64 {
	result := DB.Find(user)
	return result.RowsAffected
}

func (user *User) Match(username, password string) bool {
	DB.Where("username = ?", username).First(user)
	hashPassword := user.encryptPassword(password, user.Salt)
	return hashPassword == user.Password
}

func (user *User) RealDelete(id int) {
	DB.Unscoped().Delete(user, id)
}

func (user *User) encryptPassword(password, salt string) string {
	return utils.MD5(password + salt)
}

func (user *User) generateSalt() string {
	return utils.RandString(PasswordSaltLength)
}
