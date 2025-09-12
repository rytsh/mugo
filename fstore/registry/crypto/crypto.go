package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("crypto", Crypto{})
}

type Crypto struct{}

var defaultJWTParser = jwt.NewParser()

func (Crypto) JwtParseUnverified(token string) (map[string]any, error) {
	claims := jwt.MapClaims{}

	_, _, err := defaultJWTParser.ParseUnverified(token, claims)

	return claims, err
}

func (Crypto) Base64(v any) (string, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(conv)), nil
}

func (Crypto) Base64B(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}

func (Crypto) Base64Decode(v any) ([]byte, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(conv)
}

func (Crypto) MD5(v any) (string, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return "", err
	}

	hash := md5.Sum([]byte(conv))
	return hex.EncodeToString(hash[:]), nil
}

// SHA1B hashes v and returns its SHA1 checksum in binary.
func (Crypto) MD5B(v []byte) []byte {
	hash := md5.Sum(v)
	return hash[:]
}

// SHA1 hashes v and returns its SHA1 checksum.
func (Crypto) SHA1(v any) (string, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return "", err
	}

	hash := sha1.Sum([]byte(conv))
	return hex.EncodeToString(hash[:]), nil
}

// SHA1B hashes v and returns its SHA1 checksum in binary.
func (Crypto) SHA1B(v []byte) []byte {
	hash := sha1.Sum(v)
	return hash[:]
}

// SHA256 hashes v and returns its SHA256 checksum.
func (Crypto) SHA256(v any) (string, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(conv))
	return hex.EncodeToString(hash[:]), nil
}

// SHA256B hashes v and returns its SHA256 checksum in binary.
func (Crypto) SHA256B(v []byte) []byte {
	hash := sha256.Sum256(v)
	return hash[:]
}

// FNV32a hashes v using fnv32a algorithm.
func (Crypto) FNV32a(v any) (int, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return 0, err
	}
	algorithm := fnv.New32a()
	algorithm.Write([]byte(conv))
	return int(algorithm.Sum32()), nil
}

func (Crypto) HMAC(h any, k any, m any) (string, error) {
	ha, err := cast.ToStringE(h)
	if err != nil {
		return "", err
	}

	var hash func() hash.Hash
	switch ha {
	case "md5":
		hash = md5.New
	case "sha1":
		hash = sha1.New
	case "sha256":
		hash = sha256.New
	case "sha512":
		hash = sha512.New
	default:
		return "", fmt.Errorf("hmac: %s is not a supported hash function", ha)
	}

	msg, err := cast.ToStringE(m)
	if err != nil {
		return "", err
	}

	key, err := cast.ToStringE(k)
	if err != nil {
		return "", err
	}

	mac := hmac.New(hash, []byte(key))
	_, err = mac.Write([]byte(msg))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(mac.Sum(nil)[:]), nil
}
