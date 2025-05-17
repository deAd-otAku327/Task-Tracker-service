package dberrors

import (
	"errors"
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
)

var (
	ErrNoRowsReturned = errors.New("required rows in query result missing")
	ErrNoRowsAffected = errors.New("required affected rows in query result missing")

	// Error map with datatype name as a key.
	ErrsEnumMismatch = map[string]error{
		dbconsts.DatatypeEnumTaskStatus: errors.New("server and db task status enum mismatched"),
	}

	// Error map with constraint as a key.
	ErrsUniqueCheckViolation = map[string]error{
		dbconsts.ConstraintUserUniqueUsername:     errors.New("unique db check on username violated"),
		dbconsts.ConstraintUserUniqueEmail:        errors.New("unique db check on email violated"),
		dbconsts.ConstraintBoardToAdminPrimaryKey: errors.New("unique db check on board_id & admin_id violated"),
	}

	// Error map with constraint as a key.
	ErrsForeignKeyViolation = map[string]error{
		dbconsts.ConstraintTaskAssignieIDForeignKey: errors.New("foreign key task assignie_id to users violated"),
		dbconsts.ConstraintTaskBoardIDForeignKey:    errors.New("foreign key task board_id to dashboards violated"),
		dbconsts.ConstraintCommentTaskIDForeignKey:  errors.New("foreign key comment task_id to tasks violated"),
	}
)
