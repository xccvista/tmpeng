package utils
import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"io"
	"github.com/xccvista/mahonia"
	"strings"
)
//暂不支持含有中文相关的json

type midfile struct {
	Filename string  `json:"filename"`
	Filepath string `json:"filepath"`
	LastMD5 string `json:"lastMD5"`
	CurrentMD5 string `json:"currentMD5"`
	ScriptFile []string  `json:"scriptFile"`
}

func WriteJson(){
	b:=[]byte(`{
		"filename":"list.vm",
		"filepath":"C:/Users/vista/CODES/workspace/ws/bss-web/src/main/webapp/WEB-INF/pages/banner/",
		"lastMD5":"aaaaaa",
		"currentMD5":"bbbbb",
		"scriptFile":["aaaaaa.js","bbbb.js","ccccc.js"]
	}`)
	var onefile midfile
	err:=json.Unmarshal(b,&onefile)
	if err!=nil{
		fmt.Println("error in translating,",err.Error())
		return
	}

	var a *midfile=&midfile{
		Filename:"all.vm",
		Filepath:"D:/ad/ada/asdfa/",
		LastMD5:"nosql中文序列",
		CurrentMD5:"asdfasdfa",
	}
	fout,err:=os.Create("midfile.json")
	utf8 := mahonia.NewEncoder("utf-8")
	defer fout.Close()
	if bs,err:=json.Marshal(&a);err ==nil{
		fout.WriteString(utf8.ConvertString(string(bs)))
	}
	fmt.Println(onefile.Filename)
	ReadFile("files.txt")
}

func ReadFile(filepath string){
//	PthSep :=string(os.PathSeparator)
	fiac:=strings.Replace(filepath,`\`,"/",-1)
	fmt.Println("-------"+fiac)
	fi,err:=os.Open(fiac)
	defer fi.Close();
	if err!=nil{
		fmt.Println(err.Error())
	}
	buff:=bufio.NewReader(fi)
	for {
		line,err:=buff.ReadString('\n')
		if err!=nil || io.EOF ==err{
			break
		}
		fmt.Println(line)
		ReadFile(line)
	}
}