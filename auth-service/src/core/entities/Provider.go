package entities

type Provider struct {
	BaseEntity
	Name string `db:"name" json:"name"`
}

func NewProvider(name string) *Provider {
	return &Provider{
		BaseEntity: *NewBase(),
		Name:       name,
	}
}

func (p *Provider) GetName() string {
	return p.Name
}
