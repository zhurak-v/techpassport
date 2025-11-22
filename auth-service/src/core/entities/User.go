package entities

type User struct {
	BaseEntity
	Email    string    `db:"email" json:"email"`
	Roles    []Role    `db:"-" json:"roles"`
	Accounts []Account `db:"-" json:"accounts"`
}

func NewUser(email string) *User {
	return &User{
		BaseEntity: *NewBase(),
		Email:      email,
		Roles:      []Role{},
	}
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetRoles(roles []Role) {
	u.Roles = roles
}

func (u *User) GetRoles() []Role {
	return u.Roles
}

func (u *User) AddRole(role Role) {
	for _, r := range u.Roles {
		if r.ID == role.ID {
			return
		}
	}
	u.Roles = append(u.Roles, role)
}

func (u *User) RemoveRole(role Role) {
	newRoles := make([]Role, 0, len(u.Roles))
	for _, r := range u.Roles {
		if r.ID != role.ID {
			newRoles = append(newRoles, r)
		}
	}
	u.Roles = newRoles
}

func (u *User) GetAccounts() []Account {
	return u.Accounts
}
