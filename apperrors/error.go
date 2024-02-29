package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err 	 	error
}

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDetaFailed 	ErrCode = "S001"
	GetDataFailed 		ErrCode = "S002"
	NAData 						ErrCode = "S003"
	NoTargetData 			ErrCode = "S004"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam		 				ErrCode = "R002"
)

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
