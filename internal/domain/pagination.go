package domain

type Pagination struct {
	Offset int
	Limit  int
	Sorts  []Sort
}

type Sort struct {
	Field  string
	IsDesc bool
}
