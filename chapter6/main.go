package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"golang-a-to-z/chapter6/ecbank"
	"golang-a-to-z/chapter6/money"
)

func main() {
	// read currencies from the input
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	// parse flags
	flag.Parse()

	// parse the target currency
	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse target currency %q: %s.\n", *to, err.Error())
		os.Exit(1)
	}

	amount := parseAmount(from)

	rates := ecbank.NewClient(30 * time.Second)

	// convert the amount from the source currency to the target with the current exchange rate
	convertedAmount, err := money.Convert(amount, toCurrency, rates)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %s to %s: %s.\n", amount, toCurrency, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s = %s\n", amount, convertedAmount)
}

func parseAmount(from *string) money.Amount {
	// parse the source currency
	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *from, err.Error())
		os.Exit(1)
	}

	// read the argument
	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	// parse the quantity to convert
	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err.Error())
		os.Exit(1)
	}

	// transform value into an amount with its currency
	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return amount
}
