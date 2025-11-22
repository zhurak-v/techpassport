package entities

import "github.com/google/uuid"

type GoogleAccount struct {
	Account
	Sub       string    `db:"sub" json:"sub"`
	AccountID uuid.UUID `db:"account_id" json:"accountId"`
}

func NewGoogleAccount(
	sub string,
	user *User,
	provider *Provider,
	verified bool,
) *GoogleAccount {

	account := NewAccount(user, provider, verified)

	return &GoogleAccount{
		Account:   *account,
		AccountID: account.ID,
		Sub:       sub,
	}
}

func (g *GoogleAccount) GetSub() string {
	return g.Sub
}

func (g *GoogleAccount) GetAccount() *Account {
	return &g.Account
}
