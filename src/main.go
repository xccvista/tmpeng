package main

import(
	"fmt"
	"os"
	"./utils"
	"strings"
)


func main() {
	files,err := utils.WalkDir("C:/Users/vista/CODES/workspace/ws",".vm")
	fmt.Println(err)
	userFile := "files.txt"
	fout,err:=os.Create(userFile)
	defer fout.Close();
	if err!=nil{
		fmt.Println(userFile,err)
		return
	}
	//写入引擎处理列表,生成处理中间件
	for index,value :=range files{
		if strings.Contains(value,"target"){
			continue
		}else{
			fmt.Printf("files[%d]:%s \n",index,value);
			fout.WriteString(value+"\n")
		}
	}
	utils.WriteJson()
}
