package utils

import (
	"github.com/shanghai-edu/rmdb-lite/g"
)

func GetUserFromKey(x_api_key string) (user g.ACLUser) {
	for _, ACL := range *g.Config().AccessControl {
		if ACL.X_API_KEY == x_api_key {
			return ACL
		}
	}
	return
}
