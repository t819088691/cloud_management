package models

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
