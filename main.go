package main

import (
	"fmt"
	"time"
)

func main() {
	// consul utils
	// s, _ := GetConsulInfo("test", "10.128.0.180:8500")
	// fmt.Println(string(s))
	// WriteConsulInfo("test", `{'test':tes}`, "10.128.0.180:8500")
	// DelConsulInfo("test", "10.128.0.180:8500")

	// CompareMap
	// testtt()
	fmt.Println(GetFieldName(Student{}))
	fmt.Println(GetFieldName(&Student{}))
	fmt.Println(GetTagName(&Student{}))
	fmt.Println(GetFieldName(Teacher{}))
}
func testtt() {
	time := time.Now()
	var map1 map[string]interface{}
	map1 = make(map[string]interface{})
	map1["japan"] = "tokyo"
	map1["china"] = "beijing"
	map1["america"] = "washington"
	map1["korea"] = "seoul"
	map1["england"] = "london"
	map1["int"] = 19
	map1["time1"] = time
	map1["float"] = 0.9
	map1["int64"] = 9223372036854775807
	var map2 map[string]interface{}
	map2 = make(map[string]interface{})
	map2["japan"] = "beijing"
	map2["china"] = "tokyo"
	map2["america"] = "washington"
	map2["french"] = "paris"
	map2["canada"] = "ottawa"
	map2["int"] = 1
	map2["time1"] = time
	map2["float"] = 0.8
	map2["int64"] = 9223372036854775806
	newMap := CompareMap(map1, map2)
	for key, value := range newMap {
		fmt.Printf("key:%v   value:%v\n", key, value)
	}

}

func CompareMap(mapOne map[string]interface{}, mapTwo map[string]interface{}) map[string]string {
	newMap := make(map[string]string)
	for key1, value1 := range mapOne {
		if value2, ok := mapTwo[key1]; ok {
			if value1 != value2 {
				newMap[key1] = fmt.Sprintf("[%v]->[%v]", value1, value2)
			}
			delete(mapTwo, key1)
		} else {
			newMap[key1] = fmt.Sprintf("[%v]->[]", value1)
		}
	}
	for key, value := range mapTwo {
		newMap[key] = fmt.Sprintf("[]->[%v]", value)
	}
	return newMap
}

/**
key一样 val 一样： key没有
key一样 val 不一样： key "[v1]->[v2]"

key 在m1 存在 ； 但在m2 不存在  ： key“[v1]->[ ]”
反之一样
*/
