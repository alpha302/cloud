package main

import (
	"flag"
	"runtime"
	"cloud/funcs"
)

func init()  {
runtime.GOMAXPROCS(runtime.NumCPU())
flag.StringVar(&vars.URL, "u", "", "目标URL")
flag.StringVar(&vars.FILE, "f", "", "导入.txt文件批量扫描")
flag.StringVar(&vars.NAME,"n", "", "写入shell默认一句话马🐎，B冰蝎马，G哥斯拉马")
}

func main() {
	flag.Parse()
	funcs.Menu()

}