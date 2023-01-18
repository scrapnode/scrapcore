package database

type ScanQuery struct {
	Cursor string
	Limit  int
}

type ScanResult[T any] struct {
	Cursor  string
	Records []T
}
