package errs

type Error struct {
	Cod Code
	Msg string
}

type Code string
type ID string
