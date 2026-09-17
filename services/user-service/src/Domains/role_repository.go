package domains

type RoleRepository interface {
	FindRoleById(id int) (string, error)
}
