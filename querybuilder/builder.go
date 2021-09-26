package querybuilder

// QueryBuilder is used to build a query
type QueryBuilder interface {
	Build() (*Query, error)
}

// Query represents a query
type Query struct {
	SQL string
	Args []interface{}
}

