package utils

import (
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

var err error
var CateDic map[string][]string

func init() {
	f, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatalln("读取文件失败：", err)
	}
	loadCates(f)
}

func loadCates(f *ini.File) {
	CateDic = make(map[string][]string)
	s := f.Section("video").Key("cates").String()
	for _, txt := range strings.Split(s, " ") {
		s := strings.Split(txt, ":")
		cate, patt := s[0], s[1]
		CateDic[cate] = strings.Split(patt, ",")
	}
}
