package entity

type Role string

const (
	RoleManager    Role = "manager"
	RoleTechnician Role = "technician"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleManager, RoleTechnician:
		return true
	default:
		return false
	}
}
