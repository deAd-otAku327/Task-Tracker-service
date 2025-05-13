package serverrors

import (
	"errors"
	"fmt"
	"task-tracker-service/internal/service/_shared/servconsts"
)

var (
	ErrUsernameTooShort = fmt.Errorf("username is too short, required min len = %d", servconsts.UsernameMinLen)
	ErrUsernameTooLong  = fmt.Errorf("username is too long, required max len = %d", servconsts.UsernameMaxLen)
	ErrPasswordTooShort = fmt.Errorf("password is too short, required min len = %d", servconsts.PasswordMinLen)
	ErrPasswordTooLong  = fmt.Errorf("password is too long, required max len = %d", servconsts.PasswordMaxLen)

	ErrEmailInvalid              = errors.New("email address is invalid")
	ErrTaskRelationFilterInvalid = errors.New("task relation filter is invalid")
	ErrTaskStatusInvalid         = errors.New("task status is invalid")

	ErrIDInvalid       = errors.New("provided id is invalid")
	ErrTitleFieldEmpty = errors.New("required non-empty 'title' is empty")
	ErrTextFieldEmpty  = errors.New("required non-empty 'text' is empty")

	ErrSomethingWentWrong = errors.New("sorry, something went wrong")
	ErrInvalidPassword    = errors.New("password is invalid")
)
