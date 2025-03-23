package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"math/big"
	"strings"
	"warhoop/app/log"
)

const byteSize = 32

// h1Hash calculates H1 hash based on username and password.
func h1Hash(username, password string) []byte {
	combined := strings.ToUpper(username) + ":" + strings.ToUpper(password)
	hash := sha1.Sum([]byte(combined))
	return hash[:]
}

// h2Hash calculates H2 hash based on salt and H1 hash.
func h2Hash(salt string, h1 string) []byte {
	combined := salt + h1
	hash := sha1.Sum([]byte(combined))
	return hash[:]
}

// bytesToLittleEndian converts a byte slice to its little-endian representation.
func bytesToLittleEndian(b []byte) []byte {
	// Create a copy to avoid modifying the original slice
	reversed := make([]byte, len(b))
	copy(reversed, b)
	for i := 0; i < len(reversed)/2; i++ {
		reversed[i], reversed[len(reversed)-1-i] = reversed[len(reversed)-1-i], reversed[i]
	}
	return reversed
}

// bytesToBigInt converts a byte slice to a *big.Int in little-endian format.
func bytesToBigInt(b []byte) *big.Int {
	leBytes := bytesToLittleEndian(b)
	bigInt := new(big.Int)
	bigInt.SetBytes(leBytes)
	return bigInt
}

// strPadByType pads the input string based on the specified type.
func strPadByType(input string, padLength int, padString string, padType int) string {
	if len(input) >= padLength {
		return input
	}

	padCount := padLength - len(input)
	switch padType {
	case 0: // Left padding
		return strings.Repeat(padString, padCount) + input
	case 1: // Right padding
		return input + strings.Repeat(padString, padCount)
	case 2: // Both sides
		leftPad := padCount / 2
		rightPad := padCount - leftPad
		return strings.Repeat(padString, leftPad) + input + strings.Repeat(padString, rightPad)
	default:
		return input // No padding for invalid type
	}
}

// SRP struct represents Secure Remote Password protocol parameters.
type SRP struct {
	N        *big.Int
	g        *big.Int
	saltByte []byte
}

// NewSRP initializes and returns an SRP instance with default values.
func NewSRP() *SRP {
	srp := &SRP{
		N: new(big.Int),
		g: new(big.Int),
	}
	srp.N.SetString("894B645E89E1535BBDAD5B8B290650530801B18EBFBF5E8FAB3C82872A3E9BB7", 16)
	srp.g.SetInt64(7)
	return srp
}

// GenerateSalt generates a random salt and stores it in the SRP instance.
func (srp *SRP) GenerateSalt() error {
	salt := make([]byte, byteSize)
	_, err := rand.Read(salt)
	if err != nil {
		log.Get().Error("utils.GenerateSalt",
			log.String("err", err.Error()),
		)
		return err
	}
	srp.saltByte = salt
	return nil
}

// GetSalt returns the salt as a string.
func (srp *SRP) GetSalt() string {
	return string(srp.saltByte)
}

// PowAndMod calculates (g^h2) mod N using SRP parameters.
func (srp *SRP) PowAndMod(h2 *big.Int) *big.Int {
	result := new(big.Int)
	result.Exp(srp.g, h2, srp.N)
	return result
}

// CreateVerifier generates a salt and verifier for a given username and password.
func CreateVerifier(username, password string) (salt []byte, verifier []byte, err error) {
	srp := NewSRP()

	// Generate salt
	err = srp.GenerateSalt()
	if err != nil {
		log.Get().Error("utils.CreateVerifier",
			log.String("err", err.Error()),
		)
		return nil, nil, err
	}
	salt = srp.saltByte

	// Compute H1 and H2 hashes
	h1 := h1Hash(username, password)
	h2 := h2Hash(string(salt), string(h1))
	h2BigInt := bytesToBigInt(h2)

	// Calculate verifier
	verifierBigInt := srp.PowAndMod(h2BigInt)
	verifierBytes := bytesToLittleEndian(verifierBigInt.Bytes())

	// Pad the verifier to the desired length
	verifierStr := strPadByType(string(verifierBytes), 32, "0", 1)
	verifier = []byte(verifierStr)
	return salt, verifier, nil
}

// ConfirmVerifier verifies the username and password against a given salt and verifier.
func ConfirmVerifier(username, password string, salt []byte) ([]byte, []byte, error) {
	srp := NewSRP()

	// Compute H1 and H2 hashes
	h1 := h1Hash(username, password)
	h2 := h2Hash(string(salt), string(h1))
	h2BigInt := bytesToBigInt(h2)

	// Calculate verifier
	verifierBigInt := srp.PowAndMod(h2BigInt)
	verifierBytes := bytesToLittleEndian(verifierBigInt.Bytes())

	// Pad the verifier to the desired length
	verifierStr := strPadByType(string(verifierBytes), 32, "0", 1)
	verifier := []byte(verifierStr)
	return salt, verifier, nil
}
