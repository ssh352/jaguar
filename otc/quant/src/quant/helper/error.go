package helper

// Error struct is uesd for return error message
type Error struct {
	ErrorMsg string
}

func (e *Error) Error() string {
	return e.ErrorMsg
}
