package store

import (
	"fmt"
	"strings"
	"testing"
)

/* TestStore is a fuction that helps us test db functions
	It opens database, and returns function that will clear and close db after testing*/
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // says our test that it is a help method, no need to test it

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE",
			strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}