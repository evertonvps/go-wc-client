package rest

import (
	"time"
)

type ApiConfig struct {
	API             bool
	APIPrefix       string
	Version         string
	Timeout         time.Duration
	VerifySSL       bool
	QueryStringAuth string
	OauthTimestamp  time.Time
	ConsumerKey     string
	ConsumerSecret  string
}
