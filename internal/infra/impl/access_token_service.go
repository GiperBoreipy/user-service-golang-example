package impl

import (
	"time"

	"user_service/internal/application/data_objects"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtAccessTokenServiceImpl struct {
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewJwtAccessTokenServiceImpl(accessSecret string, refrestSecret string, accessExpire time.Duration, refreshExpire time.Duration) JwtAccessTokenServiceImpl {
	return JwtAccessTokenServiceImpl{
		accessSecret:  accessSecret,
		refreshSecret: refrestSecret,
		accessTTL:     accessExpire,
		refreshTTL:    refreshExpire,
	}
}

type tokenClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s JwtAccessTokenServiceImpl) CreateAccessToken(userId uuid.UUID) (data_objects.UserAuthToken, error) {
	token, err := s.generateToken(userId, s.accessSecret, s.accessTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s JwtAccessTokenServiceImpl) CreateRefreshToken(userId uuid.UUID) (data_objects.UserAuthToken, error) {
	token, err := s.generateToken(userId, s.refreshSecret, s.refreshTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s JwtAccessTokenServiceImpl) GetUserId(authToken data_objects.UserAuthToken) (uuid.UUID, error) {
	userID, err := s.parseToken(string(authToken), s.accessSecret)
	if err == nil {
		return userID, nil
	}

	// Если не получилось — пробуем refresh
	userID, err = s.parseToken(string(authToken), s.refreshSecret)
	if err == nil {
		return userID, nil
	}

	return uuid.Nil, jwt.ErrInvalidKey
}

func (s JwtAccessTokenServiceImpl) generateToken(userId uuid.UUID, secret string, ttl time.Duration) (data_objects.UserAuthToken, error) {
	now := time.Now()

	claims := tokenClaims{
		UserId: userId.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return data_objects.UserAuthToken(signed), nil
}

func (s JwtAccessTokenServiceImpl) parseToken(tokenString string, secret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(secret), nil
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return uuid.Nil, jwt.ErrInvalidKey
	}

	return uuid.Parse(claims.UserId)
}
