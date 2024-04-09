package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:unique`
	Password     string `json:"password"` // Remember to hash passwords before storing
	Username     string `json:"username"`
	GcpProjectID string `json:"gcpProjectID"`
	GCPFile      []byte `json:"gcpFile"`
	AzClientID   string `json:"azClientID"`
	AzSecret     string `json:"azSecret"`
	AzTenantID   string `json:"azTenantID"`
	IbmApiKey    string `json:"ibmApiKey"`
}
