package exception

type UnauthorizedError struct {
	Err string
}

func NewUnauthorizedError(err string) UnauthorizedError {
	return UnauthorizedError{Err: err}
}

func (e UnauthorizedError) Error() string {
	return e.Err
}
