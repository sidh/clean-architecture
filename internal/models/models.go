package models

// AuthAction types of actions requiring authentication
type AuthAction int

const (
	// AuthActionStore action for storing key/value pair
	AuthActionStore AuthAction = iota
	// AuthActionLoad action for loading value for a given key
	AuthActionLoad
)

// Value represents value of key/value pair
type Value struct {
	// Data is key's data
	Data string
	// Meta are additional labels with data
	Meta map[string]string
	// Count is a count of something, can be absent
	Count *int
}
