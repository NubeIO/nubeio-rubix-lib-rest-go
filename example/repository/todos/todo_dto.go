package todos

type TODODto struct {
	Title     string `json:"title"`
	IsChecked bool   `json:"is_checked"`
}
