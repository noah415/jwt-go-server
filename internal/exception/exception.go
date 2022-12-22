package exception

type ErrorType int16

const (
	ValidationError ErrorType = iota
	BadRequestError
	AuthorizationError
)

type Exception struct {
	ErrType ErrorType
	Err     error
}

func (e *Exception) Error() string {
	return e.ErrString() + ":" + e.Err.Error()
}

func (e *Exception) ErrString() string {
	switch e.ErrType {
	case ValidationError:
		return "Validation"
	case BadRequestError:
		return "Bad Request"
	case AuthorizationError:
		return "Authorization"
	default:
		return "Unknown"
	}
}
