package database

type ScanQuery struct {
	Cursor string
	Limit  int
}

type ScanResult struct {
	Cursor string
}
