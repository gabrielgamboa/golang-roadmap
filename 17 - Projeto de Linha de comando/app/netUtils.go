package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Found %d IP's for %s host: \n", len(ips), host)

	for _, ip := range ips {
		fmt.Println("* ", ip)
	}
}

func searchServerNames(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Found %d Server Names for %s host: \n", len(servers), host)

	for _, server := range servers {
		fmt.Println("* ", server.Host)
	}
}
