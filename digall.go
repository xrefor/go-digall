package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/likexian/whois-go"
)

func whoisQuery(query string) {
	result, err := whois.Whois(query)
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile("[0-9]{9}")
	np := regexp.MustCompile("N.PRI.[0-9]{8}")
	orggrep := re.FindAllString(result, -1)
	nprigrep := np.FindAllString(result, -1)
	orgnr := strings.Trim(fmt.Sprint(orggrep), "[]")
	npri := strings.Trim(fmt.Sprint(nprigrep), "[]")

	if len(npri) > 0 {
		fmt.Println("[+] ID:", npri)
	} else if len(orgnr) > 0 {
		fmt.Println("[+] Org.nr:", orgnr)
	} else {
		color.Red("[-] ERROR. Check if org is deleted.")
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
	fmt.Println(cnameRecord, err)
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
	services := [...]string{"sipfederationtls", "autodiscover", "sip", "tls", "tcp", "xmpp-server", "VLMCS", "pbx","h323ls"
			       "ts3","stun"}
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

func main() {
	//Query argument (digall <arg>)
	if len(os.Args) > 1 {
		// input argument
		query := os.Args[1]
		// variable for A record subdomain
		www := "www."
		//Used to check protocoll of SRV Record (See below)
		proto := [3]string{"tcp", "tls", "udp"}
		//Run whois query with info
		color.Green("[ DIGALL ]")
		color.Yellow("\n[+] Checking whois information")
		color.Yellow("[!] Currently only works for .no domains\n")
		whoisQuery(query)
		//Run query functions and print info
		//----------------------------------
		//A Records
		color.Yellow("\n[+] Starting DNS queries\n")
		color.Cyan("[+] A Records\n")
		aRecord(query)
		color.Magenta("\n[!] with www. :")
		aRecord(www + query)
		//CNAME Records
		color.Cyan("\n[+] CNAME Record(" + query + ")")
		cnameRecord(www + query)
		//MX Records
		color.Cyan("\n[+] MX Records\n")
		mxRecord(query)
		//TXT Records
		color.Cyan("\n[+] TXT Record(s)\n")
		txtRecord(query)
		//SRV Records using slice proto to check protocoll tcp and tls
		color.Cyan("\n[+] SRV Record(s)\n")
		srvRecord(query, proto[0])
		srvRecord(query, proto[1])
		srvRecord(query, proto[2])
		fmt.Printf("\n")
		//NS Records
		nsRecord(query)
		color.Green("\n[+] Done\n")

	} else {
		color.Red("[-] ERROR: Not a valid argument\n")
		color.Yellow("[!] Usage: ./digall <domain>\n[!] QUITTING")
	}
}
