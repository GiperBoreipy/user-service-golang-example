package data_objects

type UserAuthToken string

const UserAuthTokenContextKey UserAuthToken = "user_auth_token"

type AccessToken struct {
	AccessToken  string
	RefreshToken string
}
