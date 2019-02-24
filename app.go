package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

//Golang遍历文件夹下的所有文件
func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}

func listAll(path string, curHier int){
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil{fmt.Println(err); return}

	for _, info := range fileInfos{
		if info.IsDir(){
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(),"\\")
			listAll(path + "/" + info.Name(),curHier + 1)
		}else{
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
func main()  {
	str, _ := os.Getwd()
	//err := GetAllFile(str)
	//fmt.Println(err)

	dir := os.Args[0]
	fmt.Println(dir)  // C:\Users\Administrator\AppData\Local\Temp\___go_build_app_go.exe

	listAll(str,0)
}