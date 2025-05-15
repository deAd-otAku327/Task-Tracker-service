package helpers

import (
	"task-tracker-service/internal/storage/db/_shared/dbconsts"
	"task-tracker-service/internal/storage/db/_shared/dberrors"

	"github.com/lib/pq"
)

func CatchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case dbconsts.PQInvalidTextRepresentationError:
			return dberrors.ErrsEnumMismatch[pqErr.DataTypeName]

		case dbconsts.PQUniqueViolationError:
			return dberrors.ErrsUniqueCheckViolation[pqErr.Constraint]

		case dbconsts.PQForeignKeyViolation:
			return dberrors.ErrsForeignKeyViolation[pqErr.Constraint]
		}
	}

	return err
}
