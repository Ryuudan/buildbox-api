package render

import (
	"math"
	"net/url"
	"strconv"
)

type Meta struct {
	Page          int    `json:"page"`
	Size          int    `json:"size"`
	TotalFiltered int    `json:"total_filtered"`
	Count         int    `json:"count"`
	LastPage      int    `json:"last_page"`
	From          int    `json:"from"`
	To            int    `json:"to"`
	Query         string `json:"query"`
}

// Basic query params for pagination and searching
type QueryParams struct {
	Page  int    `json:"page,omitempty"`  // Offset for pagination (default: 0)
	Limit int    `json:"limit,omitempty"` // Number of items per page (default: 10)
	Query string `json:"query,omitempty"` // Query string for filtering
}

type PaginatedResults struct {
	Meta    *Meta       `json:"meta"`
	Results interface{} `json:"results"`
}

func ParseQueryFilterParams(rawQuery string) (*QueryParams, error) {
	values, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	params := &QueryParams{}
	params.Page = getIntValue(values, "page", 1)       // set default values when not specified
	params.Limit = getIntValue(values, "limit", 15)    // set default values when not specified
	params.Query = getStringValue(values, "query", "") // set default values when not specified

	return params, nil
}

func getIntValue(values url.Values, key string, defaultValue int) int {
	value := values.Get(key)

	if result, err := strconv.Atoi(value); err == nil && result != 0 {
		return result
	}

	return defaultValue
}

// getStringValue retrieves a string value from url.Values
func getStringValue(values url.Values, key string, defaultValue string) string {
	value := values.Get(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetSortDirectionString(sortDirection int) string {
	if sortDirection == 1 {
		return "asc"
	} else {
		return "desc"
	}
}

func GenerateMeta(total int, queryParams *QueryParams, count int) *Meta {
	return &Meta{
		Page:          queryParams.Page,
		Size:          queryParams.Limit,
		TotalFiltered: total,
		Count:         count,
		LastPage:      int(math.Ceil(float64(total) / float64(queryParams.Limit))),
		From:          (queryParams.Page-1)*queryParams.Limit + 1,
		To:            (queryParams.Page-1)*queryParams.Limit + count,
		Query:         queryParams.Query,
	}
}
