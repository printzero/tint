package mod

// Tint struct holds the whole library
type Tint struct {
	terminalLevel int
}

// Init initializes variables that tint uses and then returns the
// pointer to a Tint struct
func Init() *Tint {
	return &Tint{}
}
