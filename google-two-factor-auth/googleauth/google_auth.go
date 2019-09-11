package googleauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func GenGoogleAuthTotpUri(secret, accountName, issuer, algo, digits, period string) string {
	inputNoSpaces := strings.Replace(secret, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	issuerStr := url.QueryEscape(issuer)
	accountStr := url.QueryEscape(accountName)
	return "otpauth://totp/" + issuerStr + ":" + accountStr + "?secret=" + inputNoSpacesUpper + "&issuer=" + issuerStr + "&algorithm=" + algo + "&digits=" + digits + "&period=" + period
}

func GetGoogleAuth(input string, ts int64) (string, error) {
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	key, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	cData := make([]byte, 8)
	binary.BigEndian.PutUint64(cData, uint64(ts/30))
	pwd := oneTimePassword(key, cData)
	return fmt.Sprintf("%06d", pwd), nil
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	//fmt.Println("hash to bytes:", hash)
	offset := hash[len(hash)-1] & 0x0F
	//fmt.Println("offset is:", offset)
	hashParts := hash[offset : offset+4]
	//fmt.Println("hashParts is:", hashParts)
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)

	pwd := number % 1000000
	return pwd
}
