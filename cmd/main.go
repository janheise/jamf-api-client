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

	fmt.Println(computers)

	//req, err := http.NewRequest(http.MethodGet, "https://moneyforward.jamfcloud.com/JSSResource/computerreports/id/1", nil)
	//
	//if err != nil {
	//	log.Printf("failed to form a new request: %v", err)
	//}
	//
	//req.Header.Add("accept", "application/json")
	//req.Header.Add("authorization", "Basic YXVkaXRlcjpicElVXk9LUzF2NmI=")
	//
	//resp, err := http.DefaultClient.Do(req)
	//defer resp.Body.Close()
	//
	//if err != nil {
	//	log.Printf("failed requesting %v: %err", req, err)
	//}
	//
	//if hoge, err := httputil.DumpResponse(resp, true); err != nil {
	//	fmt.Println(hoge)
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Printf("Failed to read response body: %s", resp.Body)
	//}
	//
	//var report jamf.ComputerInfos
	//if err := json.Unmarshal(body, &report); err != nil {
	//	log.Println(fmt.Sprintf("Failed to decode json message from slack: %v -> %s",err, body))
	//} else {
	//	fmt.Println(report)
	//}
}
