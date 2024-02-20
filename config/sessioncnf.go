package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Storeinstance struct {
	St *session.Store
}

var Store Storeinstance

func SessCnf() {
	sessions := session.New(session.Config{
		Storage: Rds(),
	})

	Store.St = sessions
}
