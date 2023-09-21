package goExample

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

var SHA256RGX = regexp.MustCompile(`[0-9a-fA-F]{64}`)

func GetIntHashValue(ctx context.Context, input []byte) uint64 {
	var h = sha256.New()
	h.Write(input)
	byteHashValue := h.Sum(nil)
	return binary.BigEndian.Uint64(byteHashValue)
}

func IsSha256Str(str string) bool {
	if len(str) != 64 {
		return false
	}
	matchString := SHA256RGX.MatchString(str)
	return matchString
}

func Sha256Hash(id string) [32]byte {
	return sha256.Sum256([]byte(id))
}

func Sha256HashString(id string) string {
	value32B := Sha256Hash(id)
	value := hex.EncodeToString(value32B[:])
	return value
}

func TestSha256(t *testing.T) {
	var s string = "foo"
	fmt.Printf("sha256 of %v is %v\n", s, Sha256HashString(s))
}

func TestIsSha256(t *testing.T) {
	arr := []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"eb045d78d273107348b0300c01d29b7552d622abbc6faf81b3ec55359aa9950c",
		"2a539d6520266b56c3b0c525b9e6128858baeccb5ee9b694a2906e123c8d6dd3",
		"c6e52c372287175a895926604fa738a0ad279538a67371cd56909c7917e69ea1",
		"a683c5c5349f6f7fb903ba8a9e7e55d0ba1b8f03579f95be83f4954c33e81098",
		"f18a2548c063c5a2b1560c6f2b9ec44bf9ed9017884404016d74f330119aaefe",
		"74234e98afe7498fb5daf1f36ac2d78acc339464f950703b8c019892f982b90b",
		"1aef939917215ce780284dd9047f170482753bbe40a0998b3decb2f02c2606ac",
		"6f49cdbd80e1b95d5e6427e1501fc217790daee87055fa5b4e71064288bddede",
		"fcbcf165908dd18a9e49f7ff27810176db8e9f63b4352213741664245224f8aa",
		"f215faf9d88b7f0a881632ee22459ee452a296c808d261b6cc993d3a1fd0600e",
		"cac335bb965eee43de1f92cdc0f9ecf2ecc4cd2efd4694675329525bb3eade6a",
		"2f183a4e64493af3f377f745eda502363cd3e7ef6e4d266d444758de0a85fcc8",
		"140bedbf9c3f6d56a9846d2ba7088798683f4da0c248231336e6a05679e4fdfe",
		"505aed223b9003407e490fd0bc92cbcc394641cb20712588e79a6ccc01211647",
		"b5e5783eba2b482c4153b1d436098a0679d855c5f222c8494f2eb94690868a0f",
		"7e071fd9b023ed8f18458a73613a0834f6220bd5cc50357ba3493c6040a9ea8c",
		"3d9fc4bde7ceef058d65b00186e79c1f14b42687b491644c303065135b644e18",
		"c9a560c74de368dcd2e6ac5544ad12a9afd32584b71a52147303479f726e0408",
	}
	for _, a := range arr {
		if !IsSha256Str(a) {
			fmt.Println("err", a)
		}
	}
	fmt.Printf("no err")
}

func TestGenSha256(t *testing.T) {
	emailBlackList := []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"eb045d78d273107348b0300c01d29b7552d622abbc6faf81b3ec55359aa9950c",
		"2a539d6520266b56c3b0c525b9e6128858baeccb5ee9b694a2906e123c8d6dd3",
		"c6e52c372287175a895926604fa738a0ad279538a67371cd56909c7917e69ea1",
		"a683c5c5349f6f7fb903ba8a9e7e55d0ba1b8f03579f95be83f4954c33e81098",
		"f18a2548c063c5a2b1560c6f2b9ec44bf9ed9017884404016d74f330119aaefe",
		"74234e98afe7498fb5daf1f36ac2d78acc339464f950703b8c019892f982b90b",
		"1aef939917215ce780284dd9047f170482753bbe40a0998b3decb2f02c2606ac",
		"6f49cdbd80e1b95d5e6427e1501fc217790daee87055fa5b4e71064288bddede",
		"fcbcf165908dd18a9e49f7ff27810176db8e9f63b4352213741664245224f8aa",
		"f215faf9d88b7f0a881632ee22459ee452a296c808d261b6cc993d3a1fd0600e",
		"cac335bb965eee43de1f92cdc0f9ecf2ecc4cd2efd4694675329525bb3eade6a",
		"2f183a4e64493af3f377f745eda502363cd3e7ef6e4d266d444758de0a85fcc8",
		"140bedbf9c3f6d56a9846d2ba7088798683f4da0c248231336e6a05679e4fdfe",
		"505aed223b9003407e490fd0bc92cbcc394641cb20712588e79a6ccc01211647",
		"b5e5783eba2b482c4153b1d436098a0679d855c5f222c8494f2eb94690868a0f",
		"7e071fd9b023ed8f18458a73613a0834f6220bd5cc50357ba3493c6040a9ea8c",
		"3d9fc4bde7ceef058d65b00186e79c1f14b42687b491644c303065135b644e18",
		"c9a560c74de368dcd2e6ac5544ad12a9afd32584b71a52147303479f726e0408",
	}
	rawPhoneBlackList := []string{
		"NULL",
		"null",
		"undefined",
		"N/A",
		"n/a",
		"not set",
		"NaN",
		"nan",
		"None",
		"none",
		"Nil",
		"nil",
		"",
		"0",
		" ",
		"Undefined",
		"UNDEFINED",
	}
	var hashedPhoneBlackList []string
	for _, phone := range rawPhoneBlackList {
		hashedPhoneBlackList = append(hashedPhoneBlackList, Sha256HashString(phone))
	}
	for i, phone := range hashedPhoneBlackList {
		fmt.Println(strconv.Quote(phone) + "," + "// " + rawPhoneBlackList[i])
	}
	fmt.Println("====================")
	// diff
	var diffList []mapping
	for i, phone := range hashedPhoneBlackList {
		flag := false
		for _, email := range emailBlackList {
			if phone == email {
				flag = true
				break
			}
		}
		if !flag {
			diffList = append(diffList, mapping{raw: rawPhoneBlackList[i], hashed: phone})
		}
	}
	for _, diff := range diffList {
		fmt.Println(strconv.Quote(diff.hashed) + "," + "// " + diff.raw)
	}
}

type mapping struct {
	hashed string
	raw    string
}

func TestGenSha256Randomly(t *testing.T) {
	email := "xiaozhuoops@outlook.com"
	phone := "13262764732"
	//sha256_email: b22e7ff6fbbd5273c920413b4373c3053dba16c9b21b33545977a7b39f5feacf
	//sha256_phone: d9dc9a5cd595f9a6760442ca332ac6298d3a759bdafdf0e7286562814f43819a
	fmt.Println("sha256_email:", Sha256HashString(email))
	fmt.Println("sha256_phone:", Sha256HashString(phone))
}
