package security

type AuthToken interface {
	GenerateToken(id, role string) (string, error)
}
