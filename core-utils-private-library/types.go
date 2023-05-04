package utils

type (
	Map map[string]interface{}
)

// ValidationError indicates that occurs a validation error
type ValidationError string

// Error returns the string value of ValidationError
func (v ValidationError) Error() string {
	return string(v)
}

type GormPaginationData struct {
	Limit      int         `json:"page_size,omitempty;query:limit"`
	Page       int         `json:"current_page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *GormPaginationData) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *GormPaginationData) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *GormPaginationData) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *GormPaginationData) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}
