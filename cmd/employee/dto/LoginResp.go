package dto

import (
	"net/http"
	"time"
)

type LoginResponse struct {
	Name    string
	Value   string
	Expires time.Time
	Cookie  http.Cookie
}
