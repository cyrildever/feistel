package exception

// WrongCipherParametersError ...
type WrongCipherParametersError struct {
	message string
}

func (e *WrongCipherParametersError) Error() string {
	return e.message
}

// NewWrongCipherParametersError ...
func NewWrongCipherParametersError() *WrongCipherParametersError {
	return &WrongCipherParametersError{
		message: "wrong cipher parameters: keys and rounds can't be null",
	}
}
