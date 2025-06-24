package svcaccess

// AssertionColumn represents a column assertion
type AssertionColumn struct {
	column string
	table  string
	s      string
}

// NewAssertionColumn creates a new column assertion
func NewAssertionColumn(table, column string) AssertionColumn {
	return AssertionColumn{
		column: column,
		table:  table,
		s:      table + "." + column,
	}
}

// String returns the string representation of the column assertion
func (ac AssertionColumn) String() string {
	return ac.s
}
