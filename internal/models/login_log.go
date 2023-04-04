package models

import "time"

// 用户登录日志

type LoginLog struct {
	ID        int    `json:"id"`
	Username  string `json:"username" gorm:"not null"`
	IP        string `json:"ip" gorm:"not null"`
	CreatedAt time.Time
}

func (log *LoginLog) Create() (insertId int, err error) {
	err = DB.Create(log).Error
	if err == nil {
		insertId = log.ID
	}
	return
}

func (log *LoginLog) List() ([]LoginLog, error) {
	list := make([]LoginLog, 0)
	err := DB.Find(&list).Error
	return list, err
}

func (log *LoginLog) Total() int64 {
	result := DB.Find(log)
	return result.RowsAffected
}
