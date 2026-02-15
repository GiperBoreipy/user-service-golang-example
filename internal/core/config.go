package core

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
}

func LoadConfig() Config {
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if accessSecret == "" {
		accessSecret = "access-secret"
	}

	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	if refreshSecret == "" {
		refreshSecret = "refresh-secret"
	}

	accessTTL := 15 * time.Minute
	if v := os.Getenv("ACCESS_TOKEN_TTL"); v != "" {
		if minutes, err := strconv.Atoi(v); err == nil {
			accessTTL = time.Duration(minutes) * time.Minute
		}
	}

	refreshTTL := 30 * 24 * time.Hour
	if v := os.Getenv("REFRESH_TOKEN_TTL"); v != "" {
		if hours, err := strconv.Atoi(v); err == nil {
			refreshTTL = time.Duration(hours) * time.Hour
		}
	}

	return Config{
		AccessTokenSecret:  accessSecret,
		RefreshTokenSecret: refreshSecret,
		AccessTokenTTL:     accessTTL,
		RefreshTokenTTL:    refreshTTL,
	}
}
