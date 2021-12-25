package querybuilder

// Inserter is used to construct an insert query
type Inserter struct {
}

func (i *Inserter) Build() (*Query, error) {
	panic("implement me")
}

// Insert generate Inserter to factory insert query
func Insert() *Inserter {
	return &Inserter{}
}

// Columns specifies the columns that need to be inserted
// if cs is empty, all columns will be inserted except auto increment columns
func (i *Inserter) Columns(cs ...string) *Inserter {
	panic("implements me")
}

// Values specify the rows
// all the elements must be the same structure
func (i *Inserter) Values(values ...interface{}) *Inserter {
	panic("implement me")
}
