package objects

type Object interface {
	// Returns the name of the object
	Name() string
	// Returns the hash representation of the object.
	Hash() int
	// String representation of the object.
	String() string
}

