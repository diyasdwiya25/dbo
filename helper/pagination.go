package helper

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func PaginationFormat(recordcount int, limit, page int) *Pagination {
	var (
		tmpl = Pagination{}
	)

	total := (recordcount / limit)

	// Calculator Total Page
	remainder := (recordcount % limit)
	if remainder == 0 {
		tmpl.TotalPage = total
	} else {
		tmpl.TotalPage = total + 1
	}

	// Set current/record per page meta data
	tmpl.CurrentPage = page
	tmpl.RecordPerPage = limit

	// Calculator the Next/Previous Page
	if page <= 0 {
		tmpl.Next = page + 1
	} else if page < tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = page + 1
	} else if page == tmpl.TotalPage {
		tmpl.Previous = page - 1
		tmpl.Next = 0
	}

	return &tmpl
}
