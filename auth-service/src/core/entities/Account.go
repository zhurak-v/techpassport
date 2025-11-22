package entities

import "github.com/google/uuid"

type Account struct {
	BaseEntity
	UserID     uuid.UUID `db:"user_id" json:"userId"`
	ProviderID uuid.UUID `db:"provider_id" json:"providerId"`
	Verified   bool      `db:"verified" json:"verified"`
	User       *User     `db:"-" json:"user"`
	Provider   *Provider `db:"-" json:"provider"`
}

func NewAccount(
	user *User,
	provider *Provider,
	verified bool,
) *Account {
	return &Account{
		BaseEntity: *NewBase(),
		UserID:     user.ID,
		ProviderID: provider.ID,
		Verified:   verified,
		User:       user,
		Provider:   provider,
	}
}

func (a *Account) GetUserID() uuid.UUID {
	return a.UserID
}

func (a *Account) GetProviderID() uuid.UUID {
	return a.ProviderID
}

func (a *Account) GetVerified() bool {
	return a.Verified
}

func (a *Account) GetUser() *User {
	return a.User
}

func (a *Account) GetProvider() *Provider {
	return a.Provider
}
