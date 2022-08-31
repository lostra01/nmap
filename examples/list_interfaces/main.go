package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lostra01/nmap/v2"
)

func main() {
	scanner, err := nmap.NewScanner()
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	interfaceList, err := scanner.GetInterfaceList()
	if err != nil {
		log.Fatalf("could not get interface list: %v", err)
	}

	bytes, err := json.MarshalIndent(interfaceList, "", "\t")
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}

	fmt.Println(string(bytes))
}
