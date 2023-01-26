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
}

type Service struct {
	host string
}

func (svc *Service) Start() error {
	fmt.Println("Connecting to", svc.host)
	return nil
}
