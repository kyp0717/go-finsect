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

type Stocks []Stock
type PriceDelta []float64

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

func getSnapShot(sym []string) Stocks {
	// This channel is buffer since we do not want any blocking
	chStk := make(chan Stock, len(sym))

	var ss Stocks
	var wg sync.WaitGroup
	wg.Add(len(sym))
	for _, x := range sym {
		go reqStkPrice(chStk, &wg, x)
	}
	wg.Wait()
	close(chStk)
	for i := 1; i <= len(sym); i++ {
		ss = append(ss, <-chStk)
	}
	return ss

}

func Find(stk string, stks Stocks) (Stock, error) {
	for _, s := range stks {
		if stk == s.symbol {
			return s, nil
		}
	}
	return Stock{}, nil
}

func GetDelta(sl []string) PriceDelta {
	var pdiffs PriceDelta
	s1 := getSnapShot(sl)
	time.Sleep(15 * time.Second)
	s2 := getSnapShot(sl)
	for _, stk := range sl {
		x1, _ := Find(stk, s1)
		x2, _ := Find(stk, s2)
		pdiff := x2.price - x1.price
		pdiffs = append(pdiffs, pdiff)
	}
	return pdiffs

}

var spider = []string{"XLE", "XLB", "XLF"}

func main() {
	s := getSnapShot(spider)
	fmt.Println(s)
	fmt.Println(Find("XLE", s))

	pd := GetDelta(spider)
	fmt.Println(pd)

}
