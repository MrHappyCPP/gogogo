package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/*buf, err := ioutil.ReadFile("slash.json")
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}*/

	//fmt.Printf("got bytes %d\n", len(buf))

	/*links, err := getServiceLinks(buf)
	fmt.Printf("got Links %d\n", len(links))

	for _, v := range links {
		fmt.Println(v)
	}*/

	getOverviewFromSlash()
}

func getOverviewFromSlash() (err error) {
	resp, err := http.Get("http://slash.service.prod.rewe-digital.com/services")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		//bodyString := string(bodyBytes)
		//fmt.Println("Got from Slash", bodyString)
		links, err := getServiceLinks(bodyBytes)
		fmt.Println("Got Links", len(links))
	}
	return nil
}

func getServiceLinks(buf []byte) (links []string, err error) {
	var result []string
	var parsedJson map[string]interface{}

	//here we get the "services" element from the json
	err = json.Unmarshal([]byte(buf), &parsedJson)
	if err != nil {
		panic(err)
	}

	//now call some one who can extract the elements from "services" element
	services := parsedJson["services"].(interface{})
	fmt.Println(services)

	for _, v := range services.([]interface{}) {
		//fmt.Println(reflect.TypeOf(v))
		lolo := v.(map[string]interface{})
		//fmt.Println(reflect.TypeOf(lolo))
		v := lolo["link"]
		//fmt.Println(reflect.TypeOf(v))
		//fmt.Println(v)
		result = append(result, v.(string))
	}

	return result, nil
}
