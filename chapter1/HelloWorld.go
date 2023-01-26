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

	var resource = resource{
		path: "/Users/vrnsky/Desktop",
		data: nil,
	}
	resource.deny()

	myCar := &car{
		engine{volume: 25},
	}

	fmt.Println(myCar.engine.volume)
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

// Example of interface declaration
type clickable interface {
	Click(s string) (n string, err error)
}

type resource struct {
	path string
	data []byte
}

func (r *resource) deny() {
	fmt.Println("Access to", r.path, "is restricted")
}

type car struct {
	engine engine
}

type engine struct {
	volume float64
}
