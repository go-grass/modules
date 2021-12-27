package modules

import "rsc.io/quote"

func ReadQuotes() string {
	return quote.Glass()
}

func SayHello() string {
	return quote.Hello()
}

func HelloGo() string {
	return quote.Go()
}