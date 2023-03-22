package config

import (
	"time"
)

const (
	HTTP_WRITE_TIMEOUT = time.Second * 15
	HTTP_READ_TIMEOUT  = time.Second * 15
	HTTP_IDLE_TIMEOUT  = time.Second * 60
)
