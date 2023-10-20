package connections

import "database/sql"

// UserRepository defines the structure of a user repository
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// SelectAllUsers select all records from users table
func (r *UserRepository) SelectAllUsers() ([]User, error) {
	// select all users
	rows, err := r.DB.Query("SELECT * FROM users")

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
