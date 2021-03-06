package main

import (
	"log"
	jamf "jamf-api-client"
	"fmt"
	"os"
)

const url = "https://moneyforward.jamfcloud.com/JSSResource"

func main() {
	if os.Getenv("USER") == "" || os.Getenv("PASSWORD") == "" {
		log.Fatalf("Must specify both USER and PASSWORD")
	}

	c, err := jamf.NewClient(os.Getenv("USER"), os.Getenv("PASSWORD"), url, true)
	if err != nil {
		log.Fatalf("[ERROR] Failed to create a client: %v", err)
	}

	service := &jamf.ComputerService{Client:c}
	computers, err := service.GetComputers()

	if err != nil {
		log.Printf("failed to fetch computers: %v", err)
	}

	jamf.SortByComputerName(computers)
	for _, computer := range computers {
		fmt.Println(fmt.Sprintf("%s, %s", computer.FullName, computer.ComputerName))
	}
}
