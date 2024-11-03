package variable

import (
	"fmt"
	"src/common/ctype"
	"src/util/dbutil"
	"testing"
)

var repo Repo

func TestMain(m *testing.M) {
	dbutil.InitDb()
	repo = New(dbutil.Db())

	seedData()
	m.Run()
}

func getData(id uint) ctype.Dict {
	return ctype.Dict{
		"Key":         fmt.Sprintf("key%d", id),
		"Value":       fmt.Sprintf("value%d", id),
		"Description": fmt.Sprintf("description%d", id),
		"DataType":    "STRING",
	}
}

func seedData() {
	for i := 0; i < 10; i++ {
		data := getData(uint(i))
		repo.Create(data)
	}
}

func TestList(t *testing.T) {
	queryOptions := ctype.QueryOptions{}
	result, err := repo.List(queryOptions)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(result) != 10 {
		t.Errorf("Expected 10 items, got %d", len(result))
	}
}

func TestRetrieve(t *testing.T) {
	queryOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"key": "key1"},
	}
	result, err := repo.Retrieve(queryOptions)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
}

func TestCreate(t *testing.T) {
	data := getData(11)
	result, err := repo.Create(data)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
}

func TestUpdate(t *testing.T) {
	data := getData(1)
	item, _ := repo.Retrieve(ctype.QueryOptions{Filters: ctype.Dict{"key": data["Key"]}})
	result, err := repo.Update(item.ID, data)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
}

func TestGetOrCreate(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		data := getData(12)
		queryOptions := ctype.QueryOptions{Filters: ctype.Dict{"key": "key11"}}
		result, err := repo.GetOrCreate(queryOptions, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}
	})
	t.Run("Create", func(t *testing.T) {
		data := getData(14)
		queryOptions := ctype.QueryOptions{Filters: ctype.Dict{"key": data["Key"]}}
		result, err := repo.GetOrCreate(queryOptions, data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result, got nil")
		}

		list, _ := repo.List(ctype.QueryOptions{})
		if len(list) != 12 {
			t.Errorf("Expected 12 items, got %d", len(list))
		}
	})
}

func TestDelete(t *testing.T) {
	item, _ := repo.Retrieve(ctype.QueryOptions{Filters: ctype.Dict{"key": "key1"}})
	_, err := repo.Delete(item.ID)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	list, _ := repo.List(ctype.QueryOptions{})
	if len(list) != 11 {
		t.Errorf("Expected 11 items, got %d", len(list))
	}
}

func TestDeleteList(t *testing.T) {
	list, _ := repo.List(ctype.QueryOptions{})
	ids := make([]uint, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	_, err := repo.DeleteList(ids)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	list, _ = repo.List(ctype.QueryOptions{})
	if len(list) != 0 {
		t.Errorf("Expected 0 items, got %d", len(list))
	}
}
