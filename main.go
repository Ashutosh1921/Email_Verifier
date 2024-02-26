package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")
	for scanner.Scan() {
		Checkdomain(scanner.Text())
	}
	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	// till here we have just read the input given by the user

}

func Checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	// if mxrecords is not empty then it means that domain has mx records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Errors:%v\n", err)
	}
	// now we find spf records in given string
	for _, records := range txtRecords {
		if strings.HasPrefix(records, "v=spf1") {
			hasSPF = true
			spfRecord = records
			break
		}
	}
	// now search for the dmar records
	dmarcRec, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range dmarcRec {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// now printing all the values
	fmt.Printf("%v,%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, hasDMARC, mxRecords, dmarcRecord, spfRecord)

}
