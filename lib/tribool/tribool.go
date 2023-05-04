package tribool

type Tribool int

const (
	Unset Tribool = iota
	With
	Without
)
