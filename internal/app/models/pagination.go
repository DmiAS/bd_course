package models

import "time"

type Pagination struct {
	Limit  int   `query:"limit"`
	Cursor int64 `query:"cursor"`
}

const defaultLimit = 5

func GetPaginationInfo(pag *Pagination) *Pagination {
	var created, limit = pag.Cursor, pag.Limit
	if pag.Cursor == 0 {
		created = time.Now().UnixNano()
	}

	if pag.Limit == 0 {
		limit = defaultLimit
	}

	return &Pagination{
		Limit:  limit,
		Cursor: created,
	}
}
