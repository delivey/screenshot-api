package screenshot

import (
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateHash(url string) string {
	unixTime := time.Now().Unix()
	rawHash := fmt.Sprintf("%s_%d", url, unixTime)
	encodedHash := base64.StdEncoding.EncodeToString([]byte(rawHash))
	return encodedHash
}
