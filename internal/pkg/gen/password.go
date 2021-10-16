package gen

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	random "math/rand"
	"time"

	"github.com/alexsergivan/transliterator"
)

const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func init() {
	random.Seed(time.Now().Unix())
}
func GenReadableString(size int) string {
	length := len(chars)
	res := make([]byte, size)

	for i := range res {
		res[i] = chars[random.Intn(length)]
	}

	return string(res)
}

func Login(firstName, lastName string) string {
	trans := transliterator.NewTransliterator(nil)
	firstName = trans.Transliterate(firstName, "en")
	lastName = trans.Transliterate(lastName, "en")

	randomString := GenReadableString(5)
	today := time.Now()
	dayNumber := today.Weekday()
	monthNumber := today.Month()
	yearNumber := today.Year() % 100 // get last two digits
	login := fmt.Sprintf(
		"%c%c%d%s%d%d",
		firstName[0],
		lastName[0],
		dayNumber,
		randomString,
		monthNumber,
		yearNumber,
	)
	return login
}

func GenerateRandomString(size int) ([]byte, error) {
	str := make([]byte, size)
	if _, err := rand.Read(str); err != nil {
		return nil, err
	}

	return str, nil
}

func PasswordWithSalt(password, salt []byte) ([]byte, error) {
	var sha512Hasher = sha256.New()

	// Append salt to gen
	password = append(password, salt...)

	// Write gen bytes to the hasher
	if _, err := sha512Hasher.Write(password); err != nil {
		return nil, err
	}

	// Get the SHA-256 hashed gen
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return hashedPasswordBytes, nil
}
