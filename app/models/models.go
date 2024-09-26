package models

import "time"

type BaseModel struct {
	ID uint16 `json:"id,omitempty`
}

type CommonTimestampsField struct {
	CreateAt  time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}
