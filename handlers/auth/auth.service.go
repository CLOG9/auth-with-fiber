package handlers

import (
	"testfiber/config"
	"testfiber/schema"
)

func createUser(body *UserByEmail) error {
	user := &schema.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Image:    body.Image,
		Type:     body.Type,
		Token_id: body.Token_id,
	}

	if err := config.DB.Db.Create(user).Error; err != nil {
		return err
	}
	return nil

}
func GetUserByEmail(email string) (*UserByEmail, error) {
	user := &UserByEmail{}

	if err := config.DB.Db.Model(&schema.User{}).Where("email = ?", email).Select("username, email, password, type").First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetRedisSession(key string) ([]byte, error) {
	rdb := config.Rds()
	value, err := rdb.Get(key)
	if err != nil {
		return nil, err
	}
	return value, nil
}
