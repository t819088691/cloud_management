package models

import (
	"log"
	"strings"
)

type Platform struct {
	PlatformID   int    `json:"platform_id" gorm:"not null"`
	PlatformName string `json:"platform_name" gorm:"not null"`
	AccountID    int    `json:"account_id" gorm:"not null"`
}

func (p *Platform) Create(platformName string, userID int) (insertId int, err error) {

	p.PlatformName = strings.ToLower(platformName)
	switch p.PlatformName {
	case "gcp":
		p.PlatformID = 1
	case "aws":
		p.PlatformID = 2
	case "azure":
		p.PlatformID = 3
	default:
		log.Fatalln("无效的云平台")
	}

	p.AccountID = userID
	if err := DB.Create(p).Error; err == nil {
		insertId = p.PlatformID
	} else {
		log.Fatalln(err)
	}
	return
}

func (p *Platform) List() ([]Platform, error) {
	list := make([]Platform, 0)
	err := DB.Find(&list).Error
	return list, err
}

func (p *Platform) Delete(userID int, platformName string) {
	DB.Where("account_id = ? and platform_name = ?", userID, platformName).Delete(p)
}

type GCP struct {
	PlatformID       int    `json:"platform_id"`
	ProjectNumber    string `json:"project_number"`
	ProjectID        string `json:"project_id"`
	ProjectName      string `json:"project_name"`
	BillingAccount   string `json:"billing_account"`
	BillingAccountID string `json:"billing_account_id"`
	Organizations    string `json:"organizations"`
	EmailAddress     string `json:"email_address"`
	PhoneNumbs       string `json:"phone_numbs"`
	Contacts         string `json:"contacts"`
}

func (gcp *GCP) Create() (insertId int, err error) {
	if err := DB.Create(gcp).Error; err == nil {
		insertId = gcp.PlatformID
	} else {
		log.Fatalln(err)
	}
	return
}

func (gcp *GCP) Update() {

}
