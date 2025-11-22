package entities

import (
	"time"

	"github.com/google/uuid"
)

type IBaseEntityReader interface {
	GetID() uuid.UUID
	GetCreateAt() time.Time
	GetUpdateAt() time.Time
}

type IBaseEntityWriter interface {
	SetUpdateAt()
}

type IBaseEntityContract interface {
	IBaseEntityReader
	IBaseEntityWriter
}

type BaseEntity struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func NewBase() *BaseEntity {
	return &BaseEntity{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	}
}

func (b *BaseEntity) GetID() uuid.UUID       { return b.ID }
func (b *BaseEntity) GetCreateAt() time.Time { return b.CreatedAt }
func (b *BaseEntity) GetUpdateAt() time.Time {
	if b.UpdatedAt == nil {
		return time.Time{}
	}
	return *b.UpdatedAt
}

func (b *BaseEntity) SetUpdateAt() {
	now := time.Now().UTC()
	b.UpdatedAt = &now
}
