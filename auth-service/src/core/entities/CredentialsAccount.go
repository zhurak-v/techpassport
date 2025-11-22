package entities

import "github.com/google/uuid"

type CredentialsAccount struct {
	Account
	Password  string    `db:"password" json:"password"`
	AccountID uuid.UUID `db:"account_id" json:"accountId"`
}

func NewCredentialsAccount(
	password string,
	user *User,
	provider *Provider,
	verified bool,
) *CredentialsAccount {
	account := NewAccount(user, provider, verified)

	return &CredentialsAccount{
		Account:   *account,
		AccountID: account.ID,
		Password:  password,
	}
}

func (c *CredentialsAccount) GetPassword() string {
	return c.Password
}

func (c *CredentialsAccount) GetAccount() *Account {
	return &c.Account
}
