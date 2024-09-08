package config

import "testing"

func TestGetClaimsTags(t *testing.T) {
	type Claims struct {
		TestField string `claim:"test_field"`
	}

	want := "Test Field Value"
	claims := GetClaimMap(Claims{
		TestField: want,
	}, nil)

	if len(claims) == 0 {
		t.Fatalf("Claims map is empty")
	}
	result := claims["test_field"]
	if result != want {
		t.Fatalf("Received incorrect value for key.  Received: %v, Expected: %v", result, want)
	}
}
