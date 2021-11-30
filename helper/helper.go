package helper

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-protos-go/peer"
	"sort"
	"strings"
)

func PrintStruct(yourProject interface{}) {
	fmt.Printf("%#v\n", yourProject)
}

func PrintJson(target interface{}) {
	fooByte, _ := json.MarshalIndent(&target, "", "\t")
	fmt.Println(string(fooByte))
}

func PlayRes(res peer.Response) {
	var target map[string]interface{}
	json.Unmarshal(res.Payload, &target)
	PrintJson(target)

	var message map[string]interface{}
	json.Unmarshal([]byte(res.Message), &message)
	PrintJson(message)
}

func Contains(s []string, match string) bool {
	i := sort.SearchStrings(s, match)
	return i < len(s) && s[i] == match
}

func Filter(src []string) (res []string) {
	for _, s := range src {
		newStr := strings.Join(res, " ")
		if !strings.Contains(newStr, s) {
			res = append(res, s)
		}
	}
	return
}

func Intersections(section1, section2 []string) (intersection []string) {
	str1 := strings.Join(Filter(section1), " ")
	for _, s := range Filter(section2) {
		if strings.Contains(str1, s) {
			intersection = append(intersection, s)
		}
	}
	return
}

func CopyStructToStruct(target interface{}, result interface{}) error {
	fooByte, err := json.Marshal(&target)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fooByte, result)
	if err != nil {
		return err
	}
	return nil
}
