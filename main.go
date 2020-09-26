package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PostData 要打包給cf的資料
type PostData struct {
	Mode          string `json:"mode"`
	Configuration Cfdata `json:"configuration"`
	Notes         string `json:"notes"`
}

//Cfdata 放進Configuration資料
type Cfdata struct {
	Target string `json:"target"`
	Value  string `json:"value"`
}

func main() {
	addlist()
	fmt.Println("=================================")
	getlist()
}

func addlist() {
	client := &http.Client{}
	//生成要访问的url
	emp := &PostData{Mode: "whitelist", Configuration: Cfdata{Target: "ip", Value: "220.132.244.43"}, Notes: "whitelist ip"}
	e, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := "https://api.cloudflare.com/client/v4/zones/78504f4c7e9650f9cf2752689597cdee/firewall/access_rules/rules"
	//提交请求
	//var postjson = []byte(e)
	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer(e))

	//增加header选项
	reqest.Header.Add("X-Auth-Email", "lawolgood@gmail.com")
	reqest.Header.Add("X-Auth-Key", "22687ea53329db8a338f7effbee6f2405e50e")
	reqest.Header.Add("Authorization", "Bearer pkpzbvtHrHcAxPH9x3UjkXjAwpQsrrTEcF1qjulO")
	reqest.Header.Add("X-Auth-User-Service-Key", "v1.0-268fb96431256228391ab1bc-a2664f8a21d34fb0cceff16d76dc711d4a720138e838b19519b1e929e6a24a9137d296866598fed16b67ff53366d74c3420f5161764052fee8a266fff2f2e32d043ddabd7ac967")
	reqest.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	sitemap, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", sitemap)
}

func getlist() {
	client := &http.Client{}
	//生成要访问的url
	url := "https://api.cloudflare.com/client/v4/zones/37867986d29a2c886e39cdc555637edf/firewall/access_rules/rules?configuration.value=220.132.244.43"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("X-Auth-Email", "lawolgood@gmail.com")
	reqest.Header.Add("X-Auth-Key", "22687ea53329db8a338f7effbee6f2405e50e")
	reqest.Header.Add("Authorization", "Bearer pkpzbvtHrHcAxPH9x3UjkXjAwpQsrrTEcF1qjulO")
	reqest.Header.Add("X-Auth-User-Service-Key", "v1.0-268fb96431256228391ab1bc-a2664f8a21d34fb0cceff16d76dc711d4a720138e838b19519b1e929e6a24a9137d296866598fed16b67ff53366d74c3420f5161764052fee8a266fff2f2e32d043ddabd7ac967")
	reqest.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	sitemap, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", sitemap)
}
