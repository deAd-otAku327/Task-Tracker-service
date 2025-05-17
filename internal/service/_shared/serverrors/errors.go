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

	ErrNoFieldsToUpdate = errors.New("provided no fields for update")

	ErrSomethingWentWrong     = errors.New("sorry, something went wrong")
	ErrAccountIsNotRegistered = errors.New("account is not registered")
	ErrInvalidPassword        = errors.New("password is invalid")
	ErrUsernameOccupied       = errors.New("username is already occupied")
	ErrEmailOccupied          = errors.New("email is already linked to another account")

	ErrManipulationImpossible = errors.New("you have no resourse with provided id in your jurisdiction")

	ErrNoUserToAssign = errors.New("no user with provided id to be assigned")
	ErrNoBoardToLink  = errors.New("no board with provided id to be linked with")

	ErrNoTask             = errors.New("no task with provided id")
	ErrNoDashboard        = errors.New("no dashboard with provided id")
	ErrUserIsAlreadyAdmin = errors.New("user with provided id is already admin for provided dashboard")
)
