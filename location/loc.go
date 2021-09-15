package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User_agent struct {
	Product   string `json:"product"`
	Version   string `json:"version"`
	Raw_value string `json:"raw_value"`
}

type IpInformation struct {
	Ip          string     `json:"ip"`
	Ip_decimal  int        `json:"ip_decimal"`
	Country     string     `json:"country"`
	Country_iso string     `json:"country_iso"`
	Country_eu  bool       `json:"country_eu"`
	Region_name string     `json:"region_name"`
	Region_code string     `json:"region_code"`
	City        string     `json:"city"`
	Latitude    float32    `json:"latitude"`
	Longitude   float32    `json:"longitude"`
	Time_zone   string     `json:"time_zone"`
	Asn         string     `json:"asn"`
	Asn_org     string     `json:"asn_org"`
	User_agent  User_agent `json:"user_agent"`
}

func main() {
	httpGet()
}
func httpGet() {

	resp, _ := http.Get("http://ifconfig.co/json")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	obj := &IpInformation{}
	_ = json.Unmarshal([]byte(body), obj)

	fmt.Printf("this is boj: %+v ", obj)
}
