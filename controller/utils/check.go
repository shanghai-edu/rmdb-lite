package utils

import (
	"github.com/shanghai-edu/rmdb-lite/g"
)

//GetUserFromKey 根据 X-API-KEY 获取用户角色
func GetUserFromKey(xAPIKey string) (user g.ACLUser) {
	for _, ACL := range *g.Config().AccessControl {
		if ACL.XAPIKEY == xAPIKey {
			return ACL
		}
	}
	return
}
