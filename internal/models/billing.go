package models

import "time"

type Billing struct {
	BillingID           int       `json:"billing_id,omitempty" gorm:"primaryKey"`
	BillingAccountName  string    `json:"billing_account_name,omitempty"`
	BillingAccountID    int       `json:"billing_account_id,omitempty"`
	CustomerName        string    `json:"customer_name,omitempty"`
	ProductCode         string    `json:"product_code,omitempty"`
	ProductName         string    `json:"product_name,omitempty"`
	MeterType           string    `json:"meter_type,omitempty"`
	MeterName           string    `json:"meter_name,omitempty"`
	UsageStartDate      time.Time `json:"usage_start_date"`
	UsageEndDate        time.Time `json:"usage_end_date"`
	Quantity            float64   `json:"quantity,omitempty"`
	BillingCurrencyCode string    `json:"billing_currency_code,omitempty" gorm:"USD"`
	TotalBillingCost    float64   `json:"total_billing_cost,omitempty"`
	PricingCurrencyCode string    `json:"pricing_currency_code,omitempty" gorm:"USD"`
	TotalPricingCost    float64   `json:"total_pricing_cost,omitempty"`
}
