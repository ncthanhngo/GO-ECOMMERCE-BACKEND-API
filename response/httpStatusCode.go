package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrUserAlreadyExist = 20004
)

// mesage
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
	ErrUserAlreadyExist: "user already exist",
}
