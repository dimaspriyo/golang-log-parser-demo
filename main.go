package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	_ "github.com/davecgh/go-spew/spew"
)

type Log struct {
	IP          string `json:"ip"`
	Timestamp   string `json:"timestamp"`
	HttpMethod  string `json:"http_method"`
	Parameter   string `json:"parameter"`
	HttpVersion string `json:"http_version"`
	Status      string `json:"status"`
	Byte        string `json:"byte"`
	URL         string `json:"url"`
	UserAgent   string `json:"user_agent"`
}

func main() {

	if os.Args[1] == "-h" {
		fmt.Println(`
		-t  Format output [json , PlainText]
		-o  Outpath path file

		Examples: 
		./main /home/dimas/personal/LogicalTest/nginx-access.log -t json
		./main /home/dimas/personal/LogicalTest/nginx-access.log -t json -o /blablabla/output.txt 
		go run main.go /blablabla/nginx-access.log -t json -> Print log output with JSON Format
		go run main.go /blablabla/nginx-access.log -t PlainText -o /blablabla/output.txt -> Output log to a file with PlainText Format
		`)
	} else {
		res, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err.Error())
		}

		regPattern := regexp.MustCompile(`(?P<ip>.+)-(.+)-(.+)\[(?P<timestamp>.+)]\s\"(?P<http_method>.+?)\s(?P<parameter>.+?)\s(?P<http_version>.+?)\"\s(?P<status>.+?)\s(?P<byte>.+?)\s"(?P<url>.+?)"\s"(?P<user_agent>.+?)"`)

		temp := make(map[string]interface{})
		match := regPattern.FindStringSubmatch(string(res))
		for k, v := range regPattern.SubexpNames() {
			if v == "" {
				continue
			}
			temp[v] = match[k]
		}

		jsonRes, err := json.Marshal(temp)
		if err != nil {
			panic(err.Error())
		}

		log := Log{}
		err = json.Unmarshal(jsonRes, &log)
		if err != nil {
			panic(err.Error())
		}

		listArgs := os.Args[2:]

		formatFlag := "PlainText"
		outputFlag := "no"

		formatPosition := 0
		outputPosition := 0

		for k, v := range listArgs {
			if v == "-t" {
				formatPosition = k + 1
				formatFlag = listArgs[formatPosition]
			}

			if v == "-o" {
				outputPosition = k + 1
				outputFlag = listArgs[outputPosition]
			}
		}

		response := ""

		if formatFlag == "json" {
			response = string(jsonRes)
		} else {
			response = string(res)
		}

		if outputFlag == "no" {
			fmt.Println(response)
		} else {
			file, err := os.Create(outputFlag)

			if err != nil {
				return
			}
			defer file.Close()

			err = ioutil.WriteFile(outputFlag, []byte(response), 0644)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}
