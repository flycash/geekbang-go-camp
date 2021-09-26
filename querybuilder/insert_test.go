package querybuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []CommonTestCase {
		{
			name: "single",
			builder: Insert().Values(&TestModel{
				Id: 1,
				FistName: "Tom",
				Age: 18,
				LastName: "Jerry",
			}),
			wantSql: "INSERT INTO `test_model`(`id`, `first_name`, `age`, `last_name`) VALUES(?, ?, ?, ?)",
			wantArgs: []interface{}{int64(1), "Tom", 18, "Jerry"},
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			query, err := c.builder.Build()
			assert.Equal(t, c.wantErr, err)
			assert.Equal(t, c.wantSql, query.SQL)
			assert.Equal(t, c.wantArgs, query.Args)
		})
	}

}

type TestModel struct {
	Id int64 `eql:"auto_increment,primary_key"`
	FistName string
	Age int8
	LastName string
}

type CommonTestCase struct {
	name string
	builder QueryBuilder
	wantArgs []interface{}
	wantSql string
	wantErr error
}