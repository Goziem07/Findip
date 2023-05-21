package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func printBanner() {
	banner := `
 _______ _           _ _       
(_______|_)         | (_)      
 _____   _ ____   __| |_ ____  
|  ___) | |  _ \ / _  | |  _ \ 
| |     | | | | ( (_| | | |_| |
|_|     |_|_| |_|\____|_|  __/ 
                        |_|   by g0ziem
	`
	fmt.Println(banner)
}

func getIPs(domains []string, nodomain bool) []string {
	ips := []string{}
	for _, domain := range domains {
		ipAddresses, err := net.LookupIP(domain)
		if err == nil && len(ipAddresses) > 0 {
			ipAddress := ipAddresses[0].String()
			if !nodomain {
				ips = append(ips, fmt.Sprintf("%s: %s", domain, ipAddress))
			} else {
				ips = append(ips, ipAddress)
			}
		}
	}
	return ips
}

func main() {
	listFile := flag.String("l", "", "File containing a list of domain names")
	noDomain := flag.Bool("n", false, "Exclude domain name from output")
	outputFile := flag.String("o", "", "Output file path to save the results")
	flag.Parse()

	// Print the banner to show that the tool is running
	printBanner()

	// Read the list of domains from the input file
	file, err := os.Open(*listFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var domains []string
	for scanner.Scan() {
		domain := scanner.Text()
		domains = append(domains, domain)
	}
	if scanner.Err() != nil {
		fmt.Printf("Failed to read file: %v\n", scanner.Err())
		os.Exit(1)
	}

	// Resolve each domain name to its corresponding IP address
	ips := getIPs(domains, *noDomain)

	// Print out all the IPs found
	for _, ip := range ips {
		fmt.Println(ip)
	}

	// Print the total number of IPs found
	fmt.Printf("Total number of IPs found: %d\n", len(ips))

	// Save the results to a file if specified
	if *outputFile != "" {
		output, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Failed to create output file: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()

		for _, ip := range ips {
			fmt.Fprintln(output, ip)
		}
	}
}
