package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// errors.Is/errors.Asを使えるようにUnwrapメソッドを定義
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}