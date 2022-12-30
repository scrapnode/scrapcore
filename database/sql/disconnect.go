package sql

import "context"

func (db *SQL) Disconnect(ctx context.Context) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.Conn == nil {
		return nil
	}

	database, err := db.Conn.DB()
	if err != nil {
		return err
	}
	if err := database.Close(); err != nil {
		return err
	}

	db.Logger.Debug("disconnected")
	return nil
}
