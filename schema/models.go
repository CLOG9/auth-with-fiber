package schema

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

// BeforeCreate hook to generate CUID before creating a record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	return nil
}

// User model
type User struct {
	ID       string `gorm:"type:text;primaryKey;uniqueIndex"`
	Username string `gorm:"uniqueIndex"                      json:"username"`
	Email    string `gorm:"uniqueIndex"                      json:"email"`
	Password string `                                        json:"password"`
	Type     string `                                        json:"type"`
	Image    string `                                        json:"image"`
	Token_id string `                                        json:"token_id"`
}
