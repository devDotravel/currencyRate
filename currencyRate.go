package main

import (
	"currency/config"
	"currency/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

type currencyRatesResp struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     rates
}

type rates struct {
	GBP float64
	USD float64
}

type currencyReverseRates struct {
	GBP float64
	USD float64
}

func main() {
	utils.SetDirectoryTree()

	urlRequest := config.APIURL + "/latest?access_key=" + string(config.APIKEY) + "&symbols=GBP,USD"

	// request and make a struct out of response
	res, err := http.Get(urlRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var currencyRates currencyRatesResp
	json.Unmarshal(body, &currencyRates)

	// new struct with reverse logic
	reverseRates := currencyReverseRates{
		GBP: math.Ceil(100*1/(currencyRates.Rates.GBP-config.ConversionFee)) / 100,
		USD: math.Ceil(100*1/(currencyRates.Rates.USD-config.ConversionFee)) / 100,
	}

	// outputs to file
	outputContent, err := json.Marshal(reverseRates)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(config.OutputFilePath, outputContent, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
