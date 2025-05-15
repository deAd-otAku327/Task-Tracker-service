package dbconsts

const (
	PGDriverName = "postgres"
	// From http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
	PGInvalidTextRepresentationError = "invalid_text_representation"
	PGUniqueViolationError           = "unique_violation"

	ConstraintUniqueUsername = "users_username_unique"
	ConstraintUniqueEmail    = "users_email_unique"

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

	ColumnTaskStatus = "status"
)
