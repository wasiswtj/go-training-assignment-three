package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func DoPostAPI() {
	min := 1
	max := 100
	randomWind := rand.Intn(max-min) + min
	randomWater := rand.Intn(max-min) + min

	payload := map[string]int{
		"wind":  randomWind,
		"water": randomWater,
	}

	payloadBuff, _ := json.Marshal(payload)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:1323/update", bytes.NewBuffer(payloadBuff))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// var response model.ResponseAi
	// json.Unmarshal(body, &response)
}
