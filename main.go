package main

import (
	"errors"
	"fmt"
	"strings"
	"test/azureType"
	"test/zhlog"
)

func main() {
	// gotoloop()
	// err := f()
	// fmt.Println("Returned normally from f.", err)
	// split()
	// mapp()
	orm := CloudprojectEngine()
	orm.ShowSQL(true)
	dataTo := make([]azureType.BillAzureDetailReportMetadata, 0)

	err := orm.Where("bill_account_uuid=?", 1101).And("deleted=0").OrderBy("uuid").Limit(10, 10).Find(&dataTo)
	fmt.Println("erer:", err)

	zhlog.Error("Begin", "%s", err.Error())
}

func mapp() {
	test := make(map[string]interface{})

	// ele1 := make(map[string]string)
	test["key1"] = "val1"

	// ele2 := make(map[string][]string)
	test["key2"] = []string{"val2", "val22"}

	// test = append(test, ele1)

	// test = append(test, ele2)
	fmt.Println("test:", test)
	// map[string]
}

func split() {
	tag := "key:cluster value:master test:qqq"
	ele := strings.Split(tag, " ")
	fmt.Println("ele:", ele)
	if len(ele) > -2 {
		fmt.Println("len(ele):", len(ele))
	}
	for k, v := range ele {
		if k > 1 {
			break
		}
		s := strings.Split(v, ":")
		fmt.Println("s:", s)
		if s[0] == "key" {
			fmt.Println("key:", s[1])
		}
		if s[0] == "value" {
			fmt.Println("value:", s[1])
		}
	}
}

func f() (err error) {
	defer func() {
		fmt.Println("Calling g.")
		if e := recover(); e != nil {
			fmt.Println("Recovered in f", e)
			err = errors.New("new")
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
	return nil
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func gotoloop() {
	for i := 0; i < 10; i++ {
		if i > 3 {
			goto LAbEL2
		}
		fmt.Println("i:", i)
	LAbEL2:
		fmt.Println("LastLAbEL:", i)
	}
}
