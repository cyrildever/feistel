package exception

// NotUint64Error ...
type NotUint64Error struct {
	message string
}

func (e *NotUint64Error) Error() string {
	return e.message
}

// NewNotUint64Error ...
func NewNotUint64Error() *NotUint64Error {
	return &NotUint64Error{
		message: "probably not a uint64 byte array",
	}
}

// TooSmallToPreserveLengthError ...
type TooSmallToPreserveLengthError struct {
	message string
}

func (e *TooSmallToPreserveLengthError) Error() string {
	return e.message
}

// NewTooSmallToPreserveLengthError ...
func NewTooSmallToPreserveLengthError() *TooSmallToPreserveLengthError {
	return &TooSmallToPreserveLengthError{
		message: "too small to preserve length",
	}
}

// UnkownEngineError ...
type UnkownEngineError struct {
	message string
}

func (e *UnkownEngineError) Error() string {
	return e.message
}

// NewUnkownEngineError ...
func NewUnkownEngineError() *UnkownEngineError {
	return &UnkownEngineError{
		message: "unkown hash algorithm",
	}
}

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
