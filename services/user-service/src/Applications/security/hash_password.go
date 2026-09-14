package security

type PasswordHasher interface {
	Hash(password string) (string, error)
	CompareHashPassword(password, hashPassword string) error
}
