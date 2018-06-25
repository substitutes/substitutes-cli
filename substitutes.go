package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// Substitutes struct for holding data
type substitutes struct {
	Meta struct {
		Date     string `json:"date"`
		Class    string `json:"class"`
		Extended bool   `json:"extended"`
	} `json:"meta"`
	Substitutes []struct {
		Date      string `json:"date"`
		Hour      string `json:"hour"`
		Day       string `json:"day"`
		Teacher   string `json:"teacher"`
		Time      string `json:"time"`
		Subject   string `json:"subject"`
		Type      string `json:"type"`
		Notes     string `json:"notes"`
		Classes   string `json:"classes"`
		Room      string `json:"room"`
		After     string `json:"after"`
		Cancelled bool   `json:"cancelled"`
		New       bool   `json:"new"`
		Reason    string `json:"reason"`
		Counter   string `json:"counter"`
	} `json:"substitutes"`
}

// Credentials struct for importing credentials
type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func buildURL(path string) string {
	return viper.GetString("server") + "/api" + path
}

func main() {
	viper.SetDefault("server", "https://v.uff.space")
	viper.SetDefault("class", "10")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.substitutes/")
	viper.AddConfigPath("/etc/substitutes/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read configuration file: " + err.Error())
	}

	res, err := request(buildURL("/c/" + viper.GetString("class")))
	if err != nil {
		log.Fatal("Failed to get data: " + err.Error())
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed to read data: " + err.Error())
	}
	var s substitutes
	if json.Unmarshal(data, &s) != nil {
		log.Fatal("Failed to unmarshal data")
	}
	fmt.Printf("- Class: %s\n- Date: %s\n", s.Meta.Class, s.Meta.Date)
}
