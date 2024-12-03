package models

type Admin struct {
	Base     `json:"base,omitempty"`
	Username string `json:"username,omitempty" gorm:"type:VARCHAR(20)"`
	Password string `json:"password,omitempty" gorm:"type:VARCHAR"`
}
