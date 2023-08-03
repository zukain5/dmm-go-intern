package object

import (
	"time"
)

type Status struct {
	ID       int64     `json:"id,omitempty"`
	Account  *Account  `json:"account,omitempty"`
	Content  string    `json:"content,omitempty" db:"content"`
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
