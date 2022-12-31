package sql

func (db *SQL) GetConn() any {
	return db.Conn
}
