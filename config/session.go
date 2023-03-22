package config

import (
	"time"
)

const (
	SESSION_ID             = "GOSECRETSESSID"
	SESSION_NAME           = "GOSESSID"
	SESSION_AUTHENTICATION = "AUTHENTICATION"
	SESSION_DURATION       = int(time.Duration(time.Hour * 1))
)
