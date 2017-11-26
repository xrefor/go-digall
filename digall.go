package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/net/idna"
	"net"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		//Query argument parsed and encoded to ASCII
		//in case of foreign letters(e.g æøå) usage: (digall <hostname>)
		u, err := url.Parse("http://" + os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		query, err := idna.ToASCII(u.Host)
		if err != nil {
			fmt.Println(err)
		}
		// variable for A record subdomain
		www := "www."
		//Used to check protocoll of SRV Record (See below)
		proto := [3]string{"tcp", "tls", "udp"}

		//Run query functions and print info
		//----------------------------------
		//A Records
		color.Yellow("[+] Starting DNS queries\n")
		color.Cyan("[+] A Records\n")
		color.Magenta("[*] " + query)
		aRecord(query)
		color.Magenta("[*] www." + query)
		aRecord(www + query)
		//CNAME Records
		color.Cyan("\n[+] CNAME Record(www." + query + ")")
		cnameRecord(www + query)
		//MX Records
		color.Cyan("\n[+] MX Records\n")
		mxRecord(query)
		//TXT Records
		color.Cyan("\n[+] TXT Record(s)\n")
		txtRecord(query)
		//SRV Records using slice proto to check protocoll tcp, udp and tls
		color.Cyan("\n[+] SRV Record(s)\n")
		srvRecord(query, proto[0])
		srvRecord(query, proto[1])
		srvRecord(query, proto[2])
		fmt.Printf("\n")
		//NS Records
		nsRecord(query)
		color.Green("[+] Done\n")

	} else {
		color.Red("[-] ERROR: Not a valid argument\n")
		color.Yellow("[!] Usage: ./digall <hostname>\n[!] QUITTING")
	}
}

//Lookup functions
func aRecord(query string) {
	ipRecord, err := net.LookupIP(query)
	for i := 0; i < len(ipRecord); i++ {
		fmt.Printf("IP : %s \n", ipRecord[i])
	}
	if err != nil {
		//panic(err)
	}
}

func cnameRecord(query string) {
	cnameRecord, err := net.LookupCNAME(query)
	fmt.Println(cnameRecord)
	if err != nil {
		//panic(err)
	}
}

func mxRecord(query string) {
	mxRecord, err := net.LookupMX(query)
	for i := 0; i < len(mxRecord); i++ {
		fmt.Printf("Host : %s Priority : %d \n", mxRecord[i].Host, mxRecord[i].Pref)
	}
	if err != nil {
		//panic(err)
	}
}

func txtRecord(query string) {
	txtRecord, err := net.LookupTXT(query)
	for i := 0; i < len(txtRecord); i++ {
		fmt.Printf("#%d : %s \n", i, txtRecord[i])
	}
	if err != nil {
		//panic(err)
	}
}

func srvRecord(query, proto string) {
	services := [...]string{"sipfederationtls", "autodiscover", "sip", "tls", "tcp", "xmpp-server", "VLMCS", "pbx", "h323ls",
		"ts3", "stun"}
	for _, service := range services {
		cname, addrs, err := net.LookupSRV(service, proto, query)
		for i := 0; i < len(addrs); i++ {
			fmt.Printf("addrs[%d].Target : %s \n", i, addrs[i].Target)
			fmt.Printf("addrs[%d].Port : %d \n", i, addrs[i].Port)
			fmt.Printf("addrs[%d].Priority : %d \n", i, addrs[i].Priority)
			fmt.Printf("addrs[%d].Weight : %d \n", i, addrs[i].Weight)
			color.Cyan("cname : %s \n", cname)
		}
		if err != nil {
			//panic(err)
		}
	}
}

func nsRecord(query string) {
	nsRecord, err := net.LookupNS(query)
	color.Cyan("[+] NS Records\n")
	for i := 0; i < len(nsRecord); i++ {
		fmt.Printf("NS : %s \n", nsRecord[i].Host)
	}
	if err != nil {
		//panic(err)
	}
}
