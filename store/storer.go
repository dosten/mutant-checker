package store

// Storer implements a basic key-value store interface
type Storer interface {
	Get(string) (string, error)
	Set(string, string) error
	Increment(string) error
}
