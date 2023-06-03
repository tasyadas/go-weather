package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(15 * time.Second)
		randomize()
	}
}

func randomize() {
	var (
		waterStatus string
		windStatus  string
	)
	elNumber := map[string]int{}

	elNumber["water"] = rand.Intn(100)
	elNumber["wind"] = rand.Intn(100)

	if elNumber["water"] <= 5 {
		waterStatus = "aman"
	}

	if elNumber["water"] >= 6 && elNumber["water"] <= 8 {
		waterStatus = "siaga"
	}

	if elNumber["water"] > 8 {
		waterStatus = "bahaya"
	}

	if elNumber["wind"] <= 6 {
		windStatus = "aman"
	}

	if elNumber["wind"] >= 7 && elNumber["wind"] <= 15 {
		windStatus = "siaga"
	}

	if elNumber["wind"] > 15 {
		windStatus = "bahaya"
	}

	fmt.Println(elNumber)
	fmt.Printf("status water : %s \n", waterStatus)
	fmt.Printf("status wind : %s \n\n\n", windStatus)

	data := map[string]interface{}{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8100/record", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
