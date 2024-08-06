package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/likexian/whois"
)

func scanPort(protocol, hostname string, port int) bool {
    address := hostname + ":" + strconv.Itoa(port)
    conn, err := net.DialTimeout(protocol, address, 1*time.Second)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

func main() {

	popularPorts := []int{80, 443, 22, 21, 25, 110, 23, 53, 123, 3389}

	outputFile, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	//? Defer means execute it end of the function that's belong to
	defer outputFile.Close()

	var domain string
	fmt.Print("Type a domain: ")
	fmt.Scan(&domain)

	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {

			fmt.Fprintln(outputFile, "IPv4: ", ipv4)
			fmt.Fprintln(outputFile,"")
		}
	}

	result, err := whois.Whois(domain)
	if err == nil {
		fmt.Fprintln(outputFile,result)
	}

	for _, port := range popularPorts {
        if scanPort("tcp", domain, port) {
            fmt.Printf("Port %d is open on %s\n", port, domain)
            fmt.Fprintf(outputFile, "Port %d is open on %s\n", port, domain)
        } else {
            fmt.Printf("Port %d is closed on %s\n", port, domain)
            fmt.Fprintf(outputFile, "Port %d is closed on %s\n", port, domain)
        }
    }

}
