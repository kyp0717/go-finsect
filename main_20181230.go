package main

import (
	"fmt"
	"github.com/piquette/finance-go/equity"
	"time"
)

func nowx() int64 {
	return time.Now().Unix()
}

type Stock struct {
	symbol string
	price  float64
	time   int
}

func getPrice(sym string) Stock {

	q, err := equity.Get(sym)
	if err != nil {
		// Uh-oh.
		panic(err)
	}
	s := Stock{symbol: q.Quote.Symbol,
		price: q.Quote.RegularMarketPrice,
		time:  q.Quote.RegularMarketTime}
	return s
}

func getSnapShot(respond chan<- []Stock, sym []string) {
	var s []Stock
	for i := 0; i < 5; i++ {
		x := getPrice(sym)
		//aapl := <-chTech
		s = append(s, x)
		time.Sleep(500 * time.Millisecond)

	}
	respond <- s
}

var spider = []string{"XLE", "XLB", "XLF"}

func main() {
	chTech := make(chan []Stock)
	var ss [][]Stock
	for _, sp := range spider {
		go getMinute(chTech, sp)
		result := <-chTech
		ss = append(ss, result)

	}
	time.Sleep(6 * time.Second)
	fmt.Println(ss)
}
