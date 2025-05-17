package dbconsts

const (
	PQDriverName = "postgres"
	// From http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
	PQInvalidTextRepresentationError = "invalid_text_representation"
	PQUniqueViolationError           = "unique_violation"
	PQForeignKeyViolation            = "foreign_key_violation"

	ConstraintUserUniqueUsername       = "users_username_unique"
	ConstraintUserUniqueEmail          = "users_email_unique"
	ConstraintTaskAssignieIDForeignKey = "tasks_assignie_id_fk"
	ConstraintTaskBoardIDForeignKey    = "tasks_board_id_fk"
	ConstraintCommentTaskIDForeignKey  = "comments_task_id_fk"
	ConstraintBoardToAdminPrimaryKey   = "board_to_admin_pk"

	DatatypeEnumTaskStatus = "task_status_enum"

	TableUsers        = "users"
	TableTasks        = "tasks"
	TableComments     = "comments"
	TableDashboards   = "dashboards"
	TableBoardToAdmin = "board_to_admin"

	ColumnUserID       = "id"
	ColumnUserName     = "username"
	ColumnUserEmail    = "email"
	ColumnUserPassword = "password"

	ColumnTaskID          = "id"
	ColumnTaskTitle       = "title"
	ColumnTaskDescription = "description"
	ColumnTaskStatus      = "status"
	ColumnTaskAuthorID    = "author_id"
	ColumnTaskAssignieID  = "assignie_id"
	ColumnTaskBoardID     = "board_id"
	ColumnTaskUpdatedAt   = "updated_at"

	ColumnCommentID       = "id"
	ColumnCommentTaskID   = "task_id"
	ColumnCommentAuthorID = "author_id"
	ColumnCommentText     = "text"
	ColumnCommentDateTime = "date_time"

	ColumnDashboardID          = "id"
	ColumnDashboardTitle       = "title"
	ColumnDashboardDescription = "description"
	ColumnDashboardUpdatedAt   = "updated_at"

	ColumnBoardToAdminBoardID = "board_id"
	ColumnBoardToAdminAdminID = "admin_id"
)
