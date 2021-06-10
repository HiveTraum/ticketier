package domain

import (
	"github.com/google/uuid"
	"net/url"
)

type Executor struct {
	ID    uuid.UUID
	Title string
	URL   url.URL
}

type ExecutorRepository interface {
	Get(id uuid.UUID) *Executor
	Fetch() []*Executor
}
