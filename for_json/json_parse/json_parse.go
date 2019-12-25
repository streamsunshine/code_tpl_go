package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"12":"","19":"","21":"","23":"","25":"【内容中心】【下架】与平台已有内容重复","26":"【内容中心】【下架】与平台已有内容重复","29":"","3":"","32":"","34":"","37":"","39":"【内容中心】【下架】文","44":"","45":"","46":"【内容中心】【下架】内容不符合规范","47":"","48":"","49":"【内容中心】【下架】内容不符合规范","5":"","52":"【内容中心】【下架】账号判定为低质账号","53":"【内容中心】【下与平台已有内容重复","9":"","99":""}`)
	data := make(map[string]string)
	err := json.Unmarshal(jsonData, &data)
	fmt.Println(err)
}