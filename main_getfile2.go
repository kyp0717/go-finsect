package main

import (
	//"bufio"
	//"encoding/csv"
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	//"os"
)

type Exchange struct {
}

type Nasdaq struct {
	Symbol    string
	Name      string
	LastSale  float32
	MarketCap int64
	ADRTSO    string
	IPOyear   string
	Sector    string
	Industry  string
	Link      string
}

func main() {
	fileUrl := "https://www.nasdaq.com/screening/companies-by-industry.aspx?exchange=NASDAQ&render=download"

	ndqcsv, err := DownloadFile("nasdaq.csv", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(ndqcsv)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil

}
