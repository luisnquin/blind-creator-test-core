package utils

import (
	"encoding/json"

	"net/http"
)

import (
	"fmt"
	"mime"
)

// NotFoundHandler does...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ParseToJson(
		w,
		http.StatusNotFound,
		Map{"error": fmt.Sprintf("path '%s' does not exists", r.URL.Path)},
	)
}

// SuccessPaginationResponse does...
func SuccessPaginationResponse(w http.ResponseWriter, status int, pagination GormPaginationData) {
	response := Map{
		"status":       "SUCCESS",
		"total_items":  pagination.TotalRows,
		"total_pages":  pagination.TotalPages,
		"page_size":    pagination.Limit,
		"current_page": pagination.Page,
		"data":         pagination.Rows,
	}

	ParseToJson(w, status, response)
}

// ParseToJson does...
func ParseToJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}

// DecodeData does...
func DecodeData(r *http.Request, i interface{}) error {
	t := r.Header.Get("Content-Type")

	t, _, err := mime.ParseMediaType(t)
	if err != nil {
		return err
	}

	switch t {
	case "application/json":
		return json.NewDecoder(r.Body).Decode(i)
	}

	return fmt.Errorf(`unsupported mime type '%s'`, t)
}
