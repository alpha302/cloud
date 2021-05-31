package funcs

import (
	"cloud/vars"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func addHistory(word string)  {
	//检测到漏洞时将URL信息写入history.txt
	file, err := os.OpenFile(
		"history.txt",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Printf("\033[1;32m%s%v\033[0m\n","[-]无法写入history.txt", err)
	}
	var mu sync.Mutex
	defer mu.Unlock()
	mu.Lock()
	{
		defer file.Close()
		// 写字节到文件中
		word = word + "\n"
		byteSlice := []byte(word)
		_ , err = file.Write(byteSlice)
		if err != nil {
			fmt.Printf("\033[1;32m%s%v\033[0m\n","[-]无法写入history.txt", err)
		}
	}
}

func Url(s string) string {
	urll, _ := url.Parse(s)
	target := urll.Scheme + "://" + urll.Host + "/manager/html"
	return target
}

func Config() (string, string) {
	var pass string
	switch {
	case vars.NAME == "":
		//默认是admin:admin爆破
		if vars.PASS == "" {
			pass = vars.Shellcode
		} else {
			pass = vars.PASS
		}
	case vars.NAME == "A":
		//admin:admin
		pass = vars.Shellcode
		return "A", pass

	case vars.NAME == "T":
		//tomcat:tomcat
		pass = vars.Shellcode2
		return "T", pass

	default:
		return "", ""
	}
}

func Menu() {
	now := time.Now()
	fmt.Printf("\033[1;35m%s\033[0m\n", `
	| |__   _____      _| |_(_)_ __ ___   ___
	| '_ \ / _ \ \ /\ / / __| | '_ ' _ \ / _ \
	| | | | (_) \ V  V /| |_| | | | | | |  __/
	|_| |_|\___/ \_/\_/  \__|_|_| |_| |_|\___|`)
	fmt.Printf("\033[1;35m%d-%02d-%02d %02d:%02d:%02d\033[0m\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}



