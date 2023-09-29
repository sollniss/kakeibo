package domain

type ID uint64

// IDGenerator is used to generate unique IDs for domain objects.
type IDGenerator interface {
	ID() ID
}
