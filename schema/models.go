package schema

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

// CustomCUIDField defines a custom field type for CUID IDs
type CustomCUIDField struct {
	ID string `gorm:"type:text;primaryKey"`
}

// BeforeCreate hook to generate CUID before creating a record
func (c *CustomCUIDField) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = cuid.New()
	return nil
}

// User model
type User struct {
	CustomCUIDField        // Embed the custom CUID field
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}
