package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func processCustomerData(filePath string) map[string]int {
	emailDomains := make(map[string]int)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	for _, row := range records {
		if len(row) > 2 {
			email := row[2]

			splitEmail := strings.Split(email, "@")
			if len(splitEmail) >= 2 {
				domain := splitEmail[1]
				emailDomains[domain]++
			} else {
				log.Println("An email doesn't contain '@'")

			}
		}
	}

	return emailDomains
}

func main() {
	filePath := "customers.csv"
	result := processCustomerData(filePath)

	// Convert the map to a slice for sorting
	var sortedResults []string
	for domain, count := range result {
		sortedResults = append(sortedResults, fmt.Sprintf("%s: %d customers", domain, count))
	}
	sort.Strings(sortedResults)

	fmt.Println("Email Domains and Number of Customers:")
	for _, entry := range sortedResults {
		fmt.Println(entry)
	}
}
