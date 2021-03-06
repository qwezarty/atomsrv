package engine

import (
	"testing"

	"github.com/qwezarty/atomsrv/models"
	"github.com/stretchr/testify/assert"
)

func TestSqlite(t *testing.T) {
	Sqlite = "./engine.db"

	bean := &models.Base{}
	db := Startup("sqlite3", bean)
	db.DropTable(bean)

	Sqlite = "./engine/engine.db"
}

func TestGetConn(t *testing.T) {
	testCases := []struct {
		dialect string
		want    string
	}{
		{dialect: "mssql", want: "conn string of mssql"},
		{dialect: "sqlite3", want: "./engine/engine.db"},
		{dialect: "mysql", want: ""},
	}

	for _, tc := range testCases {
		got := getConn(tc.dialect)
		assert.Equal(t, tc.want, got)
	}
}
