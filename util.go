package ysx

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getUserSign(params map[string]interface{}, key string) string {
	keys := getSortKey(params)
	sortStr := ""

	keyLen := len(keys)
	for idx, k := range keys {
		sortStr += fmt.Sprintf("%s=%v", k, params[k])
		if idx < keyLen-1 {
			sortStr = fmt.Sprintf("%s&", sortStr)
		}
	}
	sortStr += key

	return encodeByMd5(sortStr)
}

func getSortKey(params map[string]interface{}) []string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func getAPIToken(apiKey string, apiSecret string, ecid string) string {
	exp := time.Now().Add(10 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"iss":  apiKey,
		"ecid": ecid,
		"exp":  exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signStr, _ := token.SignedString([]byte(apiSecret))

	return signStr
}

func encodeByMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}
