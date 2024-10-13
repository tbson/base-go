package infra

import (
	"src/module/config/schema"
	"src/util/dbutil"
)

type VariableListRepo struct{}

func (vr VariableListRepo) ListRestful(options dbutil.ListOptions, searchableFields []string) (dbutil.ListRestfulResult[schema.Variable], error) {
	pageSize := dbutil.DEFAULT_PAGE_SIZE
	var items []schema.Variable
	emptyResult := dbutil.ListRestfulResult[schema.Variable]{Items: items, Total: 0, PageSize: pageSize, TotalPages: 0, Pages: dbutil.Pages{Next: 0, Prev: 0}}
	db := dbutil.Db()
	query := db.Model(new(*schema.Variable))

	// Apply search logic
	query = dbutil.ApplySearch(query, options.Search, searchableFields)

	// Apply filters
	query = dbutil.ApplyFilters(query, options.Filters)

	// Apply order
	query = dbutil.ApplyOrder(query, options.Order)

	// Count total records before pagination
	total, err := dbutil.GetTotalRecords(query)
	if err != nil {
		return emptyResult, err
	}

	// Apply paging
	pagingREsult := dbutil.ApplyPaging(query, options.Page, total)
	query = pagingREsult.Query
	pages := pagingREsult.Pages
	totalPages := pagingREsult.TotalPages

	// Fetch the results
	result := query.Find(&items)
	if result.Error != nil {
		return emptyResult, result.Error
	}
	return dbutil.ListRestfulResult[schema.Variable]{Items: items, Total: total, Pages: pages, PageSize: pageSize, TotalPages: totalPages}, nil
}
