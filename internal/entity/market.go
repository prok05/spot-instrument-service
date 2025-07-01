package entity

import (
	"github.com/google/uuid"
	"time"
)

type Market struct {
	ID        uuid.UUID
	Name      string
	Enabled   bool
	DeletedAt time.Time
}

type ViewMarketsRequest struct {
	UserRoles []string
}

type ViewMarketsResponse struct {
	Markets []Market
}
