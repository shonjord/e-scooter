package mysql

const (
	noRows = "sql: no rows in result set"
)

// errorIsNoRows compares if the given errors is a no rows result set errors.
func errorIsNoRows(err error) bool {
	return err.Error() == noRows
}
