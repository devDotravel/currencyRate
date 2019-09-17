package main

import (
	"currencyRate/config"
	"currencyRate/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
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
	AED float64
}

type currencyReverseRates struct {
	GBP float64
	USD float64
	AED float64
}

func main() {
	utils.SetDirectoryTree()

	urlRequest := config.APIURL + "/latest?access_key=" + string(config.APIKEY) + "&symbols=GBP,USD,AED"

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

	// math and formatting to get values of currencies in EUR
	GBPinEURStr := fmt.Sprintf("%.2f", (math.Ceil(100*1/(currencyRates.Rates.GBP))/100)+config.ConversionFee)
	GBPinEUR, _ := strconv.ParseFloat(GBPinEURStr, 2)
	USDinEURStr := fmt.Sprintf("%.2f", (math.Ceil(100*1/(currencyRates.Rates.USD))/100)+config.ConversionFee)
	USDinEUR, _ := strconv.ParseFloat(USDinEURStr, 2)
	AEDinEURStr := fmt.Sprintf("%.2f", (math.Ceil(100*1/(currencyRates.Rates.AED))/100)+config.ConversionFee)
	AEDinEUR, _ := strconv.ParseFloat(AEDinEURStr, 2)

	// new struct with reverse logic
	reverseRates := currencyReverseRates{
		GBP: GBPinEUR,
		USD: USDinEUR,
		AED: AEDinEUR,
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
