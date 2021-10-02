package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Info struct {

}
// infoMap map[string]string
func InitInfoByRedisInfo(infoMap map[string]string) *Info {
	obj := Info{}  //传入 interface  无法进行 set
	r := reflect.ValueOf(&obj)
	n := reflect.TypeOf(obj)

	fieldNum := n.NumField()
	for i := 0; i < fieldNum; i++ {
		tag := n.Field(i).Tag.Get("json")
		if len(tag) > 0 {
			//取标签名。标签可能包含多个属性
			if index := strings.Index(tag, ","); index != -1 {
				tag = tag[:index]
			}
			value := r.Elem().Field(i).Interface()
			switch value.(type) {
			case string:
				r.Elem().Field(i).SetString(infoMap[tag])
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				tmp, _ := strconv.ParseFloat(infoMap[tag], 64)
				r.Elem().Field(i).SetInt(int64(tmp))
			case float32:
			case float64:
				tmp, _ := strconv.ParseFloat(infoMap[tag], 64)
				r.Elem().Field(i).SetFloat(tmp)
			default:
				//interface 类型，直接赋值 string
				tmp := reflect.ValueOf(infoMap[tag])
				r.Elem().Field(i).Set(tmp)
			}
		}
	}
	return &obj
}
func main() {
	a := "a"
	d := &Data{
		a: "a",
		b: &a,
	}
	c := InterfaceCopy(d)
	d.a = "c"
	e := c.(**Data)
	fmt.Println(c, *e)
}