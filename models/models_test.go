package models

import (
	"testing"

	"github.com/shanghai-edu/rmdb-lite/g"
)

func init() {
	g.ParseConfig("cfg.json")
	err := InitData(g.Config().Sqlite, "routers.csv")
	if err != nil {
		panic(err)
	}
	g.InitDB()

}

func Test_ReadAllRouters(t *testing.T) {
	routers := ReadAllRouters()
	t.Log(routers)
}

func Test_ReadRouter(t *testing.T) {
	router := ReadRouter("1.1.1.0")

	t.Log(router)

	router = ReadRouter("1.1.1.1")
	t.Log(router)

}

func Test_ReadMultiRouters(t *testing.T) {
	routers, failedList := ReadMultiRouters([]string{"1.1.1.1", "1.1.1.2", "1.1.1.0"})
	t.Log(routers)
	t.Log(failedList)

}

func Test_DeleteRouter(t *testing.T) {
	err := DeleteRouter("1.1.1.1")
	if err != nil {
		t.Error(err)
	}
}

func Test_UpdateRouter(t *testing.T) {
	router := Router{
		IP:         "1.1.1.2",
		Node:       "复旦大学",
		NodeDetail: "复旦大学二号节点",
	}
	err := UpdateRouter(router)
	if err != nil {
		t.Error(err)
	}
}

func Test_CreateRouters(t *testing.T) {
	router := Router{
		IP:         "1.1.1.1",
		Node:       "复旦大学",
		NodeDetail: "复旦大学二号节点",
	}
	err := CreateRouter(router)
	if err != nil {
		t.Error(err)
	}
}

func Test_ReadAllRouters2(t *testing.T) {
	routers := ReadAllRouters()

	t.Log(routers)
}
