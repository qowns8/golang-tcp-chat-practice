package util

import (
	"strings"
	"log"
	"fmt"
)

type HttpData struct {
	method string
	header map[string]string
	body string
}

func (h * HttpData) bodyParse(str string) string {
	sp := strings.Split(str, "\r\n\r\n")
	if len(sp) < 2 {
		log.Printf("body error: %s", str)
	} else {
		return sp[1]
	}
	return ""
}

func (h * HttpData) methodParse(str string) string {
	if strings.Contains(str, "GET") {
		return "GET"
	} else if strings.Contains(str, "POST") {
		return "POST"
	} else if strings.Contains(str, "PUT") {
		return "PUT"
	} else if strings.Contains(str, "DELETE") {
		return "DELTE"
	} else {
		return ""
	}
}

func (h * HttpData) headerParse(bytes []byte) map[string]string {
	result := make([]string, 30)
	lineNumber := 0
	temp := false

	if bytes[0] == 0 {
		fmt.Println("ps: splitHeader: Couldn't get httpheader, zero filter")
		result[0] = string("-1")
		return nil
	}

	for index, element := range bytes {
		if element == '\r' {
			if bytes[index+1] == '\n' {
				temp = true
			}
		}

		if temp != true {
			result[lineNumber] += string(element)
		}

		if element == '\n' {
			temp = false
			lineNumber++
		}
	}

	resultMap := make(map[string]string, 30)
	for index, item := range result{
		sp := strings.Split(item, ": ")
		if len(sp) == 2 {
			key := sp[0]
			value := sp[1]
			resultMap[key] = value
		} else if index == 0 {
			println(item)
		}
	}

	return resultMap
}
