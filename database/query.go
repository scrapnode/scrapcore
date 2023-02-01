package database

type ScanQuery struct {
	Cursor string
	Size   int
	Search string
}

type ScanResult struct {
	Cursor string
}
