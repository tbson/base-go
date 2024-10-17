package infra

import (
	"src/module/config/repo/variable"
	"src/module/config/schema"
	"src/util/restlistutil"

	"gorm.io/gorm"
)

type Repo struct {
	*variable.Repo
	db *gorm.DB
}

func (r Repo) New(db *gorm.DB) Repo {
	parent := variable.Repo{}.New(db)
	return Repo{
		Repo: &parent,
		db:   db,
	}
}

func (r Repo) List(
	options restlistutil.ListOptions,
	searchableFields []string,
) (restlistutil.ListRestfulResult[schema.Variable], error) {
	db := r.db
	pageSize := restlistutil.DEFAULT_PAGE_SIZE
	var items []schema.Variable
	emptyResult := restlistutil.ListRestfulResult[schema.Variable]{
		Items:      items,
		Total:      0,
		PageSize:   pageSize,
		TotalPages: 0,
		Pages: restlistutil.Pages{
			Next: 0,
			Prev: 0,
		},
	}
	query := db.Model(new(*schema.Variable))

	// Apply search logic
	query = restlistutil.ApplySearch(query, options.Search, searchableFields)

	// Apply filters
	query = restlistutil.ApplyFilters(query, options.Filters)

	// Apply order
	query = restlistutil.ApplyOrder(query, options.Order)

	// Count total records before pagination
	total, err := restlistutil.GetTotalRecords(query)
	if err != nil {
		return emptyResult, err
	}

	// Apply paging
	pagingREsult := restlistutil.ApplyPaging(query, options.Page, total)
	query = pagingREsult.Query
	pages := pagingREsult.Pages
	totalPages := pagingREsult.TotalPages

	// Fetch the results
	result := query.Find(&items)
	if result.Error != nil {
		return emptyResult, result.Error
	}
	return restlistutil.ListRestfulResult[schema.Variable]{Items: items, Total: total, Pages: pages, PageSize: pageSize, TotalPages: totalPages}, nil
}
