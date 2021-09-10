package main

import (
	"encoding/json"
	"fmt"
	"time"

	"k8s.io/klog/v2"
)

type MessTime struct {
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
	Second    int    `json:"second"`
	Zone      string `json:"zone"`
	Deviation int    `json:"deviation"`
}

func main() {
	TimeComputation()
}

func TimeComputation() {
	var t = time.Now()
	var ti MessTime
	mt, mt2 := t.Zone()
	month := int(t.Month())

	ti = MessTime{t.Year(), month, t.Day(), t.Hour(),
		t.Minute(), t.Second(), mt, mt2 / 3600}

	bytes, err := json.Marshal(ti)
	if err != nil {
		klog.Errorf("err: %v", err)
	}
	fmt.Println(string(bytes))
}
