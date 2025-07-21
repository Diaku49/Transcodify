package utilities

import (
	"net/http"
	"strconv"
)

func GetPageLimit(r *http.Request) (int, int) {
	page := 1
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if n, err := strconv.Atoi(pageStr); err == nil {
			page = n
		}
	}

	limit := 10 // default
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}

	return page, limit
}
