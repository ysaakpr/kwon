package orders

import "time"

// BaseModel for all other nodes in dgraph
type BaseModel struct {
	UID       string     `json:"uid,omitempty"`
	DeletedBy *int       `json:"deleted_by,omitempty"`
	CreatedBy int        `json:"created_by,omitempty"`
	UpdatedBy *int       `json:"updated_by,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
}
