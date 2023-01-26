package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	s := &Service{host: "localhost"}
	err := s.Start()
	if err != nil {
		fmt.Println(err)
	}
	secondService := &Service{host: "google.com"}
	secondErr := (*Service).Start(secondService)
	if secondErr != nil {
		fmt.Println(secondErr)
	}
	var a number = 5
	fmt.Println(a.square())
}

// Service - it is definition of structure or class in terms of Java/**
type Service struct {
	host string
}

// Start Start() it is method of Service structure -/**
func (svc *Service) Start() error {
	fmt.Println("Connecting to", svc.host)
	return nil
}

type number int

func (n number) square() number {
	return n * n
}
