package interactors

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) (bool, error)
}
