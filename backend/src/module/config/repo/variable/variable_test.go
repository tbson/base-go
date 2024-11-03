package variable

import (
	"src/common/ctype"
	"src/util/dbutil"
	"testing"
)

var repo Repo

func TestMain(m *testing.M) {
	dbutil.InitDb()

	repo = New(dbutil.Db())
	m.Run()
}

func TestList(t *testing.T) {
	queryOptions := ctype.QueryOptions{}
	_, err := repo.List(queryOptions)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
