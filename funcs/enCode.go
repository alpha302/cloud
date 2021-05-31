package funcs

import (
	"cloud/vars"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func enCode(pass, name string) string {
	//自定义tomcat口令密码
	switch {
	case pass == "cmd" && name == "cmd":
		return vars.Shellcode

	case pass != "cmd" && name == "cmd":
		h := sha256.New()
		h.Write([]byte(pass))
		replace := hex.EncodeToString(h.Sum(nil))
		shell, _ := base64.StdEncoding.DecodeString(vars.Shellcode)
		data := strings.Replace(string(shell), "04dc5b2136328a0dcb189df97734c7c72e5e1227fa0c03469a6ce608f32f1b66", replace, -1)
		data = base64.StdEncoding.EncodeToString([]byte(data))
		return data

	case pass != "cmd" && name != "cmd":
		h := sha256.New()
		h.Write([]byte(pass))
		replace := hex.EncodeToString(h.Sum(nil))
		shell, _ := base64.StdEncoding.DecodeString(vars.Shellcode)
		data := strings.Replace(string(shell), "04dc5b2136328a0dcb189df97734c7c72e5e1227fa0c03469a6ce608f32f1b66", replace, -1)
		base := base64.StdEncoding.EncodeToString([]byte(data))

		shell, _ = base64.StdEncoding.DecodeString(base)
		data = strings.Replace(string(shell), "cmd", name, -1)
		base = base64.StdEncoding.EncodeToString([]byte(data))
		return base

	case pass == "cmd" && name != "cmd":
		shell, _ := base64.StdEncoding.DecodeString(vars.Shellcode)
		data := strings.Replace(string(shell), "cmd", name, -1)
		data = base64.StdEncoding.EncodeToString([]byte(data))
		return data
	}
	return ""
}

