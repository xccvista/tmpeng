package main
import (
	"fmt"
	"os"
//	"flag"
	"io"
	"io/ioutil"
	"bufio"
//	"time"
)

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {panic(err)}
		if 0 ==n {break}
		chunks=append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {panic(err)}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {panic(err)}
		if 0 ==n {break}
		chunks=append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {panic(err)}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

//TODO 从文件中读取到内存
//TODO 读取每一行代码移除文档注释,可配置性模板注释类型
//TODO 通用模板引擎预解析处理,有开合,无开合, for\each\if else\var|constant|转义\micro\function

func main() {

//read file into cache
	f, err := ioutil.ReadFile("C:/Users/vista/CODES/workspace/ws/bss-web/src/main/webapp/WEB-INF/pages/banner/list.vm")
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	fmt.Println(string(f))

//	start := time.Now()
//	read1(file)
//	t1 := time.Now()
//	fmt.Printf("Cost time %v\n", t1.Sub(start))
//	read2(file)
//	t2 := time.Now()
//	fmt.Printf("Cost time %v\n", t2.Sub(t1))
//	read3(file)
//	t3 := time.Now()
//	fmt.Printf("Cost time %v\n", t3.Sub(t2))
}