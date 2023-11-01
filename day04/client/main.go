package main

import (
	"crypto/x509"
	"day04/ex00/utils"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	caCertFile = "./cert/minica.pem"
	certFile   = "./cert/client/cert.pem"
	keyFile    = "./cert/client/key.pem"
)

var (
	money, count int64
	candyType    string
)

type requestBody struct {
	Money      int64  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int64  `json:"candyCount"`
}

func init() {
	flag.Int64Var(&money, "m", 0, "money")
	flag.Int64Var(&count, "c", 0, "candy's count")
	flag.StringVar(&candyType, "k", "", "candy's type")
	flag.Parse()
}

func main() {
	tlsCfg, _ := utils.GetTlsConfig(certFile, keyFile)

	certPool := x509.NewCertPool()

	caCertPEM, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatal(err)
	}

	certPool.AppendCertsFromPEM(caCertPEM)

	tlsCfg.RootCAs = certPool
	tlsCfg.InsecureSkipVerify = true

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsCfg,
		},
	}

	reqBody, _ := json.Marshal(requestBody{
		Money:      money,
		CandyType:  candyType,
		CandyCount: count,
	})

	resp, err := client.Post("https://127.0.0.1:3333/buy_candy", "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}
