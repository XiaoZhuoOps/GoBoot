package goExample

import (
	"code.byted.org/aweme-go/ajson"
	"fmt"
	"testing"
)

type PeopleV1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// json zero-value null
func Test_MarshalJson_V1(t *testing.T) {
	// V1 server -> client
	jsonV1, _ := ajson.MarshalToString(PeopleV1{Name: "zhangsan"})
	// jsonV1 {"name":"zhangsan","age":0}
	println("jsonV1", jsonV1)

	// V1 client -> server
	jsonV1 = `{"name":"lisi"}`
	peopleV1 := PeopleV1{}
	// 这里需要修改外部内容，所以传引用类型
	// ajson.UnmarshalFromString(jsonV1, peopleV1)
	ajson.UnmarshalFromString(jsonV1, &peopleV1)
	// "lisi" 0
	println(peopleV1.Name, peopleV1.Age)
}

type PeopleV2 struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func Test_PeopleV2(t *testing.T) {
	// V2 server -> client
	jsonV2, _ := ajson.MarshalToString(PeopleV2{Name: "zhangsan"})
	// jsonV2 {"name":"zhangsan"}
	println("jsonV2", jsonV2)

	// V2 client -> server
	jsonV2 = `{"name":"lisi", "age":0}`
	peopleV2 := PeopleV2{}
	// 这里需要修改外部内容，所以传引用类型
	// ajson.UnmarshalFromString(jsonV2, peopleV2)
	ajson.UnmarshalFromString(jsonV2, &peopleV2)
	// "lisi" 0
	println(peopleV2.Name, peopleV2.Age)
}

type PeopleV3 struct {
	Name *string `json:"name,omitempty"`
	Age  *int    `json:"age,omitempty"`
}

func Test_PeopleV3(t *testing.T) {
	name := "zhangsan"
	jsonV3, _ := ajson.MarshalToString(PeopleV3{Name: &name})
	// jsonV3 {"name":"zhangsan"}
	println("jsonV3", jsonV3)

	jsonV3 = `{"name":"wanger"}`
	peopleV3 := PeopleV3{}
	ajson.UnmarshalFromString(jsonV3, &peopleV3)
	// "lisi" 0
	// error
	// nil pointer dereference
	// 对一个空指针，取它的内容，相当于 解引用 dereference
	//println(*(peopleV3.Name), *(peopleV3.Age))
	fmt.Println(*peopleV3.Name)

	jsonV3 = `{"name":"wanger", "age":0}`
	peopleV3 = PeopleV3{}
	ajson.UnmarshalFromString(jsonV3, &peopleV3)
	// "lisi" 0
	println(*(peopleV3.Name), *(peopleV3.Age))
}

func TestAJsonUnmarshal(t *testing.T) {
	json := `"a":123`
	b, err := ajson.Get([]byte(json), "anything").MarshalJSON()
	t.Log(b)
	t.Log(err)
	peopleV2 := PeopleV2{}
	err = ajson.Unmarshal(b, &peopleV2)
	t.Log(err)
}
