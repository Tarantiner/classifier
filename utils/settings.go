package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

var CateDic map[string][]string

func init() {
	CateDic = make(map[string][]string)
	b, err := ioutil.ReadFile("conf.ini")
	if err != nil {
		log.Fatalln("打开配置文件失败")
	}
	lineLis := strings.Split(string(b), "\r\n")
	var section string
	for _, line := range lineLis {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			s := strings.TrimSpace(line[1 : len(line)-1])
			if len(s) > 0 {
				section = s
			}
			continue
		}
		s := strings.TrimSpace(line)
		if len(s) > 0 {
			if lis, ok := CateDic[section]; !ok {
				CateDic[section] = []string{s}
			} else {
				CateDic[section] = append(lis, s)
			}
		}
	}
}
