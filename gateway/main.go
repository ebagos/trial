package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

const (
	configFileLabel   = "CONFIG_FILE"
	portLabel         = "PORT"
	configFileDefault = "config.json"
)

type configType struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func main() {
	// 環境変数から自サーバのポート番号を得る
	myPort := os.Getenv("portLabel")
	if myPort == "" {
		myPort = ":8080"
	}
	// 環境変数からコンフィグファイルのパスを得る
	configFile := os.Getenv(configFileLabel)
	if configFile == "" {
		configFile = configFileDefault
	}
	// configFile読み込み
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	// JSONに変換
	var configData []configType
	err = json.Unmarshal(data, &configData)
	if err != nil {
		log.Fatal(err.Error())
	}

	// リバースプロキシとして動作させる部分（実はもう少し考えないと遅い？）
	reverseProxy := new(httputil.ReverseProxy)
	reverseProxy.Director = func(req *http.Request) {
		for _, cfg := range configData {
			if strings.HasPrefix(req.URL.Path, cfg.Path) {
				target, err := url.Parse(cfg.Url)
				if err != nil {
					log.Fatal("parse url:", err)
				}
				req.URL.Scheme = "http"
				req.URL.Host = target.Host
			}
		}
	}

	err = http.ListenAndServe(myPort, reverseProxy)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
