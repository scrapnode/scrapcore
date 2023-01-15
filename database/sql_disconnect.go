package database

import "context"

func (db *SQL) Disconnect(ctx context.Context) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.conn == nil {
		return nil
	}

	database, err := db.conn.DB()
	if err != nil {
		return err
	}
	if err := database.Close(); err != nil {
		return err
	}

	db.logger.Debug("disconnected")
	return nil
}
