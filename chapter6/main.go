package main

import (
	"flag"
	"fmt"
	"golang-a-to-z/chapter6/money"
	"os"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Unable to parse source currency %q: %s.\n", *from, err.Error())
	}

	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse target currency %q: %s\n", *to, err.Error())
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse value %q: %s\n", value, err.Error())
	}

	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Amount:", amount, "; Currency:", toCurrency)
	convertedAmount, err := money.Convert(amount, toCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %s to %s: %s\n", amount, toCurrency, err.Error())
	}

	fmt.Printf("%s = %s\n", amount, convertedAmount)
}
