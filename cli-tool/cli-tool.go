package main

import (
	"flag"
	"fmt"
)

func main() {
	baseUrl := flag.String("url", "https://github.com/vrnsky", "URL")
	namespace := flag.String("namespace", "minikube", "Namespace of K8s")
	serviceName := flag.String("service", "mesh-service", "Name of the service")

	fmt.Printf("Start checking services. Base url: %v\n", *baseUrl)
	fmt.Printf("Namespace: %v\n", *namespace)
	fmt.Printf("Checking service [%v]\n", *serviceName)
}
