package models

import (
	"fmt"
	"net/mail"
	"task-tracker-service/internal/service/_shared/servconsts"
	"task-tracker-service/internal/service/_shared/serverrors"
	"task-tracker-service/internal/types/enum"
)

func (m UserLoginModel) Validate() error {
	if len(m.Username) < servconsts.UsernameMinLen {
		return serverrors.ErrUsernameTooShort
	}

	if len(m.Username) > servconsts.UsernameMaxLen {
		return serverrors.ErrUsernameTooLong
	}

	if len(m.Password) < servconsts.PasswordMinLen {
		return serverrors.ErrPasswordTooShort
	}

	if len(m.Password) > servconsts.PasswordMaxLen {
		return serverrors.ErrPasswordTooLong
	}

	return nil
}

func (m UserRegisterModel) Validate() error {
	if _, err := mail.ParseAddress(m.Email); err != nil {
		return serverrors.ErrEmailInvalid
	}

	if len(m.Username) < servconsts.UsernameMinLen {
		return serverrors.ErrUsernameTooShort
	}

	if len(m.Username) > servconsts.UsernameMaxLen {
		return serverrors.ErrUsernameTooLong
	}

	if len(m.Password) < servconsts.PasswordMinLen {
		return serverrors.ErrPasswordTooShort
	}

	if len(m.Password) > servconsts.PasswordMaxLen {
		return serverrors.ErrPasswordTooLong
	}

	return nil
}

func (m TaskFilterModel) Validate() error {
	if !enum.CheckTaskRelation(m.Relation) {
		return fmt.Errorf("%s:%s", serverrors.ErrTaskRelationFilterInvalid, m.Relation)
	}

	for _, status := range m.Status {
		if !enum.CheckTaskStatus(status) {
			return fmt.Errorf("%s:%s", serverrors.ErrTaskStatusInvalid, status)
		}
	}

	return nil
}

func (m TaskSummaryParamModel) Validate() error {
	if m.TaskID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.TaskID)
	}
	return nil
}

func (m TaskCreateModel) Validate() error {
	if len(m.Title) == 0 {
		return serverrors.ErrTitleFieldEmpty
	}

	if m.AssignieID != nil {
		if *m.AssignieID < 1 {
			return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, *m.AssignieID)
		}
	}

	if m.LinkedBoardID != nil {
		if *m.LinkedBoardID < 1 {
			return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, *m.LinkedBoardID)
		}
	}

	return nil
}

func (m TaskUpdateModel) Validate() error {
	if m.ID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.ID)
	}

	if m.Title == nil && m.Status == nil && m.Description == nil && m.AssignieID == nil && m.LinkedBoardID == nil {
		return serverrors.ErrNoFieldsToUpdate
	}

	if m.Title != nil {
		if len(*m.Title) == 0 {
			return serverrors.ErrTitleFieldEmpty
		}
	}

	if m.Status != nil {
		if !enum.CheckTaskStatus(*m.Status) {
			return fmt.Errorf("%s:%s", serverrors.ErrTaskStatusInvalid, *m.Status)
		}
	}

	if m.AssignieID != nil {
		if *m.AssignieID < 1 {
			return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, *m.AssignieID)
		}
	}

	if m.LinkedBoardID != nil {
		if *m.LinkedBoardID < 1 {
			return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, *m.LinkedBoardID)
		}
	}

	return nil
}

func (m CommentCreateModel) Validate() error {
	if m.TaskID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.TaskID)
	}

	if len(m.Text) == 0 {
		return serverrors.ErrTextFieldEmpty
	}

	return nil
}

func (m DashboardSummaryParamModel) Validate() error {
	if m.BoardID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.BoardID)
	}
	return nil
}

func (m DashboardCreateModel) Validate() error {
	if len(m.Title) == 0 {
		return serverrors.ErrTitleFieldEmpty
	}
	return nil
}

func (m DashboardUpdateModel) Validate() error {
	if m.ID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.ID)
	}

	if m.Title == nil && m.Description == nil {
		return serverrors.ErrNoFieldsToUpdate
	}

	if m.Title != nil {
		if len(*m.Title) == 0 {
			return serverrors.ErrTitleFieldEmpty
		}
	}

	return nil
}

func (m DashboardDeleteModel) Validate() error {
	if m.BoardID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.BoardID)
	}
	return nil
}

func (m DashboardAdminActionModel) Validate() error {
	if m.BoardID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.BoardID)
	}
	if m.UserID < 1 {
		return fmt.Errorf("%s:%d", serverrors.ErrIDInvalid, m.UserID)
	}
	return nil
}
