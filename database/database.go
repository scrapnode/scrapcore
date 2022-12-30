package database

type ListQuery struct {
	Cursor int
	Limit  int
}

type ListResult[T any] struct {
	Cursor  int
	Records []T
}
