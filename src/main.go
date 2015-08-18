package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}else{
		fmt.Print(fi.IsDir())
		return fi.IsDir()
	}
	return false
}
//获取指定目录下所有文件,不进行下一级目录搜索,可以匹配后缀过滤
func ListDir(dirPth,suffix string)(files []string,err error){
	files=make([]string,0,10)
	dir,err:=ioutil.ReadDir(dirPth)
	if err!=nil{
		return nil,err
	}
	PthSep :=string(os.PathSeparator)
	suffix =strings.ToUpper(suffix)//如果后缀有大小写,进行忽略
	for _,fi:=range dir{
		if fi.IsDir(){
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()),suffix){
			files=append(files,dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

//获取指定目录及以下所有文件,可以匹配后缀过滤
func WalkDir(dirPth string ,suffix string)(files []string,err error){
	files =make([]string,0,30)
	suffix=strings.ToUpper(suffix)

	err= filepath.Walk(dirPth,func(filename string,fi os.FileInfo,err error) error{
		if fi.IsDir(){
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()),suffix){
			files=append(files,filename)
		}
		return nil
	})
	return files,nil
}


func main() {
	files,err := WalkDir("C:\\Users\\vista\\CODES\\workspace\\ws",".vm")
	fmt.Println(err)
	userFile := "files.txt"
	fout,err:=os.Create(userFile)
	defer fout.Close();
	if err!=nil{
		fmt.Println(userFile,err)
		return
	}
	//写入引擎处理列表
	for index,value :=range files{
		if strings.Contains(value,"target"){
			continue
		}else{
			fmt.Printf("files[%d]:%s \n",index,value);
			fout.WriteString(value+"\n")
		}
	}
}
