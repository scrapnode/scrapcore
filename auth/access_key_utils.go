package auth

import "strings"

var (
	ACCESS_KEY_ID_PREF     = "aki"
	ACCESS_KEY_SECRET_PREF = "aks"
)

func IsAccessKeyPair(id, secret string) bool {
	return strings.HasPrefix(id, ACCESS_KEY_ID_PREF) && strings.HasPrefix(secret, ACCESS_KEY_SECRET_PREF)
}
