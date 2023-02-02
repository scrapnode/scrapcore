package database

func (db *SQL) Conn() any {
	return db.conn
}
