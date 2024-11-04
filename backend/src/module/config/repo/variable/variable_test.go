package variable

import (
	"fmt"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/localeutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo Repo

var queryKey = "Key"

func TestMain(m *testing.M) {
	localeutil.Init("en")
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
	t.Run("Success", func(t *testing.T) {
		data := getData(1)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{queryKey: data[queryKey]},
		}
		_, err := repo.Retrieve(queryOptions)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("Not found", func(t *testing.T) {
		data := getData(99)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{queryKey: data[queryKey]},
		}
		_, err := repo.Retrieve(queryOptions)
		assert.EqualError(t, err, "record not found")
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data := getData(11)
		_, err := repo.Create(data)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})

	t.Run("Duplicate key", func(t *testing.T) {
		data := getData(11)
		_, err := repo.Create(data)
		assert.EqualError(t, err, "value already exists")
	})
}

func TestUpdate(t *testing.T) {
	data := getData(1)
	item, _ := repo.Retrieve(
		ctype.QueryOptions{Filters: ctype.Dict{queryKey: data[queryKey]}},
	)
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
		data11 := getData(11)
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{queryKey: data11[queryKey]},
		}
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
		queryOptions := ctype.QueryOptions{
			Filters: ctype.Dict{queryKey: data[queryKey]},
		}
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
	t.Run("Success", func(t *testing.T) {
		data := getData(1)
		item, _ := repo.Retrieve(
			ctype.QueryOptions{Filters: ctype.Dict{queryKey: data[queryKey]}},
		)
		_, err := repo.Delete(item.ID)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		list, _ := repo.List(ctype.QueryOptions{})
		if len(list) != 11 {
			t.Errorf("Expected 11 items, got %d", len(list))
		}
	})
	t.Run("Fail", func(t *testing.T) {
		_, err := repo.Delete(9999)
		assert.EqualError(t, err, "record not found")
		list, _ := repo.List(ctype.QueryOptions{})
		if len(list) != 11 {
			t.Errorf("Expected 11 items, got %d", len(list))
		}
	})
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
