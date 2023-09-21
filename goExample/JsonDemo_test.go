package goExample

import (
	"code.byted.org/aweme-go/ajson"
	. "code.byted.org/gopkg/mockito"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"testing"
)

func JsonDemoTest() {
	// test json
	payload := []byte(`{"properties":{"value":"[\"1.00\"]","currency":"USD"}}`)
	result := gjson.ParseBytes(payload)
	tarValue := result.Get("properties.value").String()
	println(tarValue)

	if j := gjson.Parse(tarValue); j.IsArray() && len(j.Array()) > 0 {
		tarValue = j.Array()[0].String()
	}
	println(tarValue)

}
func getFirstEleFromStringIfArray(payload []byte, key string) string {
	arrString := ajson.Get(payload, key).GetStringDefault("")
	arr, err := ajson.GetFromString(arrString).Array()
	if err != nil {
		return arrString
	} else {
		return arr[0].GetStringDefault("")
	}
}

func Test_Gjson(t *testing.T) {
	PatchConvey("test arr to json", t, func() {
		payload := []byte(`{"arr":[$1.00]}`)

		tarValue := gjson.GetBytes(payload, "arr").String()
		fmt.Println(tarValue)
		if j := gjson.Parse(tarValue); j.IsArray() && len(j.Array()) > 0 {
			tarValue = j.Array()[0].String()
		}
		fmt.Println("tarValue===", tarValue)
	})
}

func Test_Ajson(t *testing.T) {
	// 方案一
	PatchConvey("test ajson", t, func() {
		str := "[1]"
		fmt.Println("ori_value is", str, ajson.GetFromString(str, 0).GetStringDefault(""))
		str = "1"
		fmt.Println("ori_value is", str, ajson.GetFromString(str, 0).GetStringDefault(""))
	})
	PatchConvey("test ajson parse arr", t, func() {
		// ajson
		PatchConvey("0", func() {
			candidates := []string{"", "[1.2]", "1.2", "N/A", "$1.22"}
			dataBytes := []byte(`{"ID":123456789123456789,"Name":{"first":"Alex"},"Colors":""}`)
			key := "Colors"
			for _, candidate := range candidates {
				dataBytes, _ = sjson.SetBytes(dataBytes, key, candidate)
				arr, err := ajson.Get(dataBytes, key).Array()
				if err == nil && len(arr) > 0 {
					fmt.Println(arr[0].GetStringDefault(""))
				} else {
					fmt.Println(ajson.Get(dataBytes, key).GetStringDefault(""), err)
				}

			}
		})

		// unmarshal
		PatchConvey("1", func() {
			dataBytes := []byte(`{"ID":123456789123456789,"Name":{"first":"Alex"},"Colors":"[1,2]"}`)
			var arr []interface{}
			err := ajson.UnmarshalFromString(ajson.Get(dataBytes, "Colors").GetStringDefault(""), &arr)
			fmt.Println(arr[0], err, "1")
		})
		PatchConvey("2", func() {
			dataBytes := []byte(`{"ID":123456789123456789,"Name":{"first":"Alex"},"Colors":"1"}`)
			var arr []interface{}
			err := ajson.UnmarshalFromString(ajson.Get(dataBytes, "Colors").GetStringDefault(""), &arr)
			if err != nil {
				arr = []interface{}{ajson.Get(dataBytes, "Colors").GetStringDefault("")}
			}
			fmt.Println(arr, err, "2")
		})
		PatchConvey("3", func() {
			dataBytes := []byte(`{"ID":123456789123456789,"Name":{"first":"Alex"},"Colors":""}`)
			var arr []interface{}
			err := ajson.UnmarshalFromString(ajson.Get(dataBytes, "Colors").GetStringDefault(""), &arr)
			fmt.Println(arr, err, "3")
		})
	})
}

// test ajson null
func Test_null(t *testing.T) {
	json := `{"value":[1,2]}`
	res := ajson.Get([]byte(json), "value")
	str, err := res.GetString()
	fmt.Println(str, err)
}

type label string

const (
	labelA label = "a"
	labelB label = "b"
)

func TestLabelsSetJson(t *testing.T) {
	labels := []label{labelA, labelB}
	payload := []byte(``)
	payload, _ = sjson.SetBytes(payload, "labels", labels)
	fmt.Println(payload)
}

func TestAjsonPath(t *testing.T) {
	arr := []string{"a", "b"}
	payload := []byte(`{"a":{"b":123}}`)
	res := ajson.Get(payload, arr[0], arr[1]).GetInt64Default(0)
	fmt.Println(res)
}
