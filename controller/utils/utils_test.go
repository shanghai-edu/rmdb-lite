package utils

import (
	"testing"

	"github.com/shanghai-edu/rmdb-lite/g"
)

func init() {
	g.ParseConfig("cfg.json")
}

func Test_GetUserFromKey(t *testing.T) {
	t.Log(GetUserFromKey("123"))
	t.Log(GetUserFromKey("3a1179104ae0188ed4a3c100b38343ff"))
	t.Log(GetUserFromKey("bc055c13064927a670744e4459b86c80"))
	t.Log(GetUserFromKey("771261d709af69da179a330ed8f16fa1"))
}
