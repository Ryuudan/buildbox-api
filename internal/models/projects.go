package models

import "github.com/Pyakz/buildbox-api/utils/render"

type ProjectsQueryParams struct {
	render.QueryParams
	OrderBy        string `json:"order_by,omitempty"`
	OrderDirection string `json:"order_direction,omitempty"`
}
