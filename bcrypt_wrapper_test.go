package bcrypt_wrapper

import (
	"fmt"
	"testing"
	"time"
)

func TestGetSuitableComplexity(t *testing.T) {
	cost := GetSuitableCost(int64(time.Millisecond * 250))
	fmt.Printf("found suitable cost: %d\n", cost)
}
