package connections

// SelectAllUsers select all records from users table
func SelectAllUsers() ([]User, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// select all users
	rows, err := db.Query("SELECT * FROM users")

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
