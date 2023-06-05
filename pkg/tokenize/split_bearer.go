package tokenize

import "strings"

func SplitBearer(token string) string {
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token
}
