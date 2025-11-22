package entities

type Role struct {
	BaseEntity
	Name string `db:"name" json:"name"`
}

func NewRole(name string) *Role {
	return &Role{
		BaseEntity: *NewBase(),
		Name:       name,
	}
}

func (role *Role) GetName() string {
	return role.Name
}
