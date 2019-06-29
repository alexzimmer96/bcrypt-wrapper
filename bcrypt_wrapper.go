package bcrypt_wrapper

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	benchmarkEncryptionAmount = 20
	defaultMinCost            = 10
	defaultMaxCost            = 31
	defaultMaxTime            = int64(time.Millisecond * 250)
)

var (
	NilHash = []byte("")
	NilCost = 0
)

type BCryptWrapper struct {
	cost int
}

// Returns a new BCrypt-Wrapper. Set cost "NilCost" for running a benchmark and using the benchmark-outcome as cost.
// The benchmark will then use 250ms as default value (meaning max amount of 4 passwords per second per thread).
func NewBCryptWrapper(cost int) *BCryptWrapper {
	if cost == NilCost {
		cost = GetSuitableCost(defaultMaxTime)
	}
	if cost < defaultMinCost {
		cost = defaultMinCost
	} else if cost > defaultMaxCost {
		cost = defaultMaxCost
	}
	return &BCryptWrapper{
		cost: cost,
	}
}

// Run a benchmark to get a suitable cost for your application. The benchmark will take a few seconds.
// It will return the last cost-value which is within your maxtime.
func GetSuitableCost(maxtime int64) int {
	pwBytes := make([]byte, 16)
	if _, err := rand.Read(pwBytes); err != nil {
		pwBytes = []byte("E&dWBjxaE*8V")
	}

	var executionTime int64 = 0
	var cost = defaultMinCost

	for (executionTime < maxtime) || cost+1 > defaultMaxCost {
		before := time.Now()
		for i := 0; i < benchmarkEncryptionAmount; i++ {
			_, _ = bcrypt.GenerateFromPassword(pwBytes, cost)
		}
		timeTotal := time.Now().Sub(before).Nanoseconds()
		executionTime = timeTotal / benchmarkEncryptionAmount
		cost++
	}

	return cost - 1
}

// Compares a password to a hashed password. Returns an bcrypt.ErrMismatchedHashAndPassword when password does not match
// the hash. Returns a newly hashed password when the cost from the previous one used was lower then specified in this
// wrapper.
func (wrapper *BCryptWrapper) CompareHashAndPassword(hashedPassword, password []byte) (error, []byte) {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		// Password is incorrect or there's something wrong with the encryption-string
		return err, NilHash
	}
	if passwordCost, err := bcrypt.Cost(hashedPassword); passwordCost < wrapper.cost && err != nil {
		// Password does not use actual cost. Needs re-hashing
		newHash, err := wrapper.GenerateFromPassword(password)
		if err == nil {
			// Just returning the new hash when there was not error creating it
			return nil, newHash
		}
	}
	// Returning no error and no new hash
	return nil, NilHash
}

// Generates a hash value from a password. Uses the wrappers-cost as cost-value for bcrypt.
func (wrapper *BCryptWrapper) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, wrapper.cost)
}
