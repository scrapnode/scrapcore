package database

func (db *SQL) GetConn() any {
	return db.conn
}
