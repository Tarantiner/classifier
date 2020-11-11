package main

import (
	"classifier/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var err error
var root string
var srcPath string

func checkPath(dr string) bool {
	_, err = os.Stat(dr)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
	return true
}

func ensureDir(dr string) {
	_, err = os.Stat(dr)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dr, 0755)
			if err != nil {
				log.Fatalln("创建文件夹失败：", dr)
			}
			return
		}
		log.Fatalln("获取文件信息失败：", dr)
	}
}

func initCateDir() {
	// 根据配置文件创建分类文件夹
	for cate := range utils.CateDic {
		ph := path.Join(root, cate)
		ensureDir(ph)
	}
}

func doRealCopy(src, dst string) {
	cmd := exec.Command("cmd", "/C", "xcopy", src, dst)
	_, err = cmd.Output()
	if err != nil {
		log.Printf("复制文件【%v】到【%v】失败：%v\n", src, dst, err)
	}
}

func copyFile(ph string, info os.FileInfo, err error) error {
	if err != nil || info.IsDir() {
		return nil
	}
	fileName := info.Name()
	for cate, pattS := range utils.CateDic {
		for _, patt := range pattS {
			if strings.Contains(fileName, patt) {
				dPath := strings.Replace(path.Join(root, cate), "/", "\\", -1)
				dName := dPath + "\\" + fileName
				if checkPath(dName) {
					continue
				}
				doRealCopy(ph, dPath)
			}
		}
	}
	return nil
}

func classifyAll() {
	// 遍历目录，若为文件夹，则排除，若为文件，则进行拷贝
	err = filepath.Walk(srcPath, copyFile)
	if err != nil {
		log.Fatalln("遍历源文件过程中失败：", err)
	}
	fmt.Println("所有文件完成分类！")
}

func init() {
	// 根据输入确定目标根目录并创建文件夹
	flag.StringVar(&root, "d", "e:\\cate", "目标文件夹绝对路径")
	flag.StringVar(&srcPath, "s", "", "目标文件夹绝对路径")
	flag.Parse()
	if srcPath == "" {
		log.Fatalln("必须指定源目录")
	}
	if !checkPath(srcPath) {
		log.Fatalln("源文件必须存在")
	}
	log.Printf("将在【%v】目录创建分类\n", root)
	ensureDir(root)
}

func main() {
	initCateDir()
	classifyAll()
}
