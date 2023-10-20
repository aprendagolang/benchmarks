package std

import "database/sql"

func DeleteStd(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM categories WHERE ID = $1", id)
	
	return err
}
