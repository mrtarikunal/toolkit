package toolkit

import "crypto/rand"

const ramdomStringSource = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will have access
// to all the methods with the reciever *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomStringSource
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(ramdomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
