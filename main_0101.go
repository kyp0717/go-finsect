package main

import (
	"fmt"
	"github.com/piquette/finance-go/equity"
	"sync"
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

func chkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func reqStkPrice(chStock chan<- Stock, wg *sync.WaitGroup, sym string) {
	defer wg.Done()
	q, err := equity.Get(sym)
	chkErr(err)

	s := Stock{symbol: q.Quote.Symbol,
		price: q.Quote.RegularMarketPrice,
		time:  q.Quote.RegularMarketTime}
	chStock <- s
}

func getSnapShot(sym []string) []Stock {
	// This channel is buffer since we do not want any blocking
	chStk := make(chan Stock, len(sym))

	var s []Stock
	var wg sync.WaitGroup
	wg.Add(len(sym))
	for _, x := range sym {
		go reqStkPrice(chStk, &wg, x)
	}
	wg.Wait()
	close(chStk)
	for i := 1; i <= len(sym); i++ {
		s = append(s, <-chStk)
	}
	return s

}

var spider = []string{"XLE", "XLB", "XLF"}

func main() {
	s := getSnapShot(spider)
	fmt.Println(s)

}
