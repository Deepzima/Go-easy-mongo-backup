package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"time"
)

type Config struct {
	Path      string        `json:"backupPath"`
	Frequency string        `json:"frequency"`
	StartAt   time.Time     `json:"startAt"`
	Email     string        `json:"email"`
}

func main()  {
	configFile := getConfig()
	fmt.Println(configFile)
}

func getConfig() []Config {
	raw, err := ioutil.ReadFile("/home/andream16/go/src/Go-easy-mongo-backup/config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Config
	json.Unmarshal(raw, &c)
	fmt.Printf("%v", c)
	return c
}