package render

type Meta struct {
	Page          int    `json:"page"`
	Size          int    `json:"size"`
	Total         int    `json:"total"`
	Count         int    `json:"count"`
	LastPage      int    `json:"last_page"`
	From          int    `json:"from"`
	To            int    `json:"to"`
	Query         string `json:"query"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
}

type QueryParams struct {
	Offset         int    `json:"offset,omitempty"`          // Offset for pagination (default: 0)
	Limit          int    `json:"limit,omitempty"`           // Number of items per page (default: 10)
	Query          string `json:"query,omitempty"`           // Query string for filtering
	OrderBy        string `json:"order_by,omitempty"`        // Field to order by (default: "created_at")
	OrderDirection string `json:"order_direction,omitempty"` // Order direction: "asc" or "desc" (default: "desc")
}

type PaginatedResults struct {
	Code    int         `json:"code"`
	Meta    *Meta       `json:"meta"`
	Results interface{} `json:"results"`
}
