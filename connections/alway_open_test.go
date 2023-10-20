package connections

import "testing"

func TestUserRepository_SelectAllUsers(t *testing.T) {
	db, err := Open()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	repo := NewUserRepository(db)

	users, err := repo.SelectAllUsers()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(users) == 0 {
		t.Errorf("No users found")
	}
}

func BenchmarkUserRepository_SelectAllUsers(b *testing.B) {
	db, err := Open()
	if err != nil {
		b.Errorf("Error: %s", err)
	}

	repo := NewUserRepository(db)
	defer repo.DB.Close()

	for n := 0; n < b.N; n++ {
		_, err := repo.SelectAllUsers()
		if err != nil {
			b.Errorf("Error: %s", err)
		}
	}
}
