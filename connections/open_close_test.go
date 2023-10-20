package connections

import "testing"

func TestSelectAllUsers(t *testing.T) {
	users, err := SelectAllUsers()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(users) == 0 {
		t.Errorf("No users found")
	}
}

func BenchmarkSelectAllUsers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := SelectAllUsers()
		if err != nil {
			b.Errorf("Error: %s", err)
		}
	}
}
