package bcrypt_wrapper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const (
	password      = "E&dWBjxaE*8V"
	wrongPassword = "e&dWBjxaE*8V"
	hashed        = "$2y$12$Su9Bpprp2BYnc3f12T.DW.gpQSUBoPu41RaXguWHXQmVE4c.Wvwoe"
	wrongHash     = "$2y$12$Saf423uhp2398fh98epfh3q47q9ph24734q."
	outdatedHash  = "$2y$11$KPR/tNhxP0RQcO7gSNjgJuXwu1jrwkfuEt2resN98faTtNnzq0DMa"
)

func TestNewBCryptWrapper(t *testing.T) {
	t.Run("NilCost should lead to automatic cost", func(t *testing.T) {
		wrapper := NewBCryptWrapper(NilCost)
		if wrapper.Cost < defaultMinCost || wrapper.Cost > defaultMaxCost {
			fmt.Printf("excepted a value between %d and %d as cost, got %d.\n",
				defaultMinCost, defaultMaxCost, wrapper.Cost)
			t.Fail()
		}
	})
	t.Run("Cannot undercut minCost", func(t *testing.T) {
		wrapper := NewBCryptWrapper(defaultMinCost - 1)
		if wrapper.Cost != defaultMinCost {
			fmt.Printf("excepted %d as cost, got %d.\n", defaultMinCost, wrapper.Cost)
			t.Fail()
		}
	})
	t.Run("Cannot overcome maxCost", func(t *testing.T) {
		wrapper := NewBCryptWrapper(defaultMaxCost + 1)
		if wrapper.Cost != defaultMaxCost {
			fmt.Printf("excepted %d as cost, got %d.\n", defaultMaxCost, wrapper.Cost)
			t.Fail()
		}
	})
	t.Run("set value directly", func(t *testing.T) {
		wrapper := NewBCryptWrapper(14)
		if wrapper.Cost != 14 {
			fmt.Printf("excepted %d as cost, got %d.\n", 14, wrapper.Cost)
			t.Fail()
		}
	})
}

func TestBCryptWrapper_CompareHashAndPassword(t *testing.T) {
	wrapper := NewBCryptWrapper(12)
	t.Run("wrong hash should result in error", func(t *testing.T) {
		newpass, err := wrapper.CompareHashAndPassword([]byte(wrongHash), []byte(password))
		if len(newpass) != 0 || err == nil {
			t.Fail()
		}
	})
	t.Run("right password with outdated password should newly hash password", func(t *testing.T) {
		newpass, err := wrapper.CompareHashAndPassword([]byte(outdatedHash), []byte(password))
		if err != nil {
			fmt.Printf("failing test due to error: %s.\n", err.Error())
			t.Fail()
		}
		if len(newpass) == 0 {
			fmt.Print("failing because no new password was generated.\n")
			t.Fail()
		}
	})
	t.Run("right password and matching bcrypt cost", func(t *testing.T) {
		newpass, err := wrapper.CompareHashAndPassword([]byte(hashed), []byte(password))
		if len(newpass) != 0 || err != nil {
			t.Fail()
		}
	})
	t.Run("wrong password", func(t *testing.T) {
		newpass, err := wrapper.CompareHashAndPassword([]byte(hashed), []byte(wrongPassword))
		if err != bcrypt.ErrMismatchedHashAndPassword {
			fmt.Print("failing because \"bcrypt.ErrMismatchedHashAndPassword\" was excepted as error-object.\n")
			t.Fail()
		}
		if len(newpass) > 0 {
			fmt.Print("failing because no new password should have been generated.")
			t.Fail()
		}
	})
}

func TestBCryptWrapper_GenerateFromPassword(t *testing.T) {
	wrapper := NewBCryptWrapper(12)
	result, err := wrapper.GenerateFromPassword([]byte(password))
	if costFound, err2 := bcrypt.Cost(result); err != nil || err2 != nil || costFound != wrapper.Cost {
		t.Fail()
	}
}
