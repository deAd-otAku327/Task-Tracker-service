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

	ErrSomethingWentWrong      = errors.New("sorry, something went wrong")
	ErrUsernameIsNotRegistered = errors.New("username is not registered")
	ErrInvalidPassword         = errors.New("password is invalid")
	ErrUsernameOccupied        = errors.New("username is already occupied")
	ErrEmailOccupied           = errors.New("email is already linked to another account")
	ErrNoUserToAssign          = errors.New("no user with provided id to be assigned")
	ErrNoBoardToLink           = errors.New("no board with provided id to be linked with")
)
