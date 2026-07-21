package helper

import "testing"

func TestHashAndComparePassword(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword error: %v", err)
	}
	if hash == password {
		t.Fatalf("Hash same plain password")
	}

	if err := ComparePassword(hash, password); err != nil {
		t.Errorf("Compare password fail: %v", err)
	}

	if err := ComparePassword(hash, "wrongPassword"); err == nil {
		t.Error("Invalid password, but compare pass")
	}
}
