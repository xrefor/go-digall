package main

import (
        "fmt"
        "os"
        "net"
        "github.com/fatih/color"
        "github.com/likexian/whois-go"
        "regexp"
        "strings"
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
            fmt.Println("[+] ID: ", npri)
        } else if len(orgnr) > 0 {
            fmt.Println("[+] Org.nr:",orgnr)
        } else {
            color.Red("[-] ERROR. Check if org is deleted.")
        }
        return
 }


//Lookup functions
func aRecord(query string) {
        ipRecord, err := net.LookupIP(query)
    color.Cyan("[+] A Record\n")
    for i := 0; i < len(ipRecord); i++ {
        fmt.Printf("IP : %s \n", ipRecord[i])
    }
        if err != nil {
                //panic(err)
        }
        return
}

func cnameRecord(query string) {
        cnameRecord, err := net.LookupCNAME(query)
        color.Cyan("\n[+] CNAME Record("+query+")")
        fmt.Println(cnameRecord, err)
        if err != nil {
                //panic(err)
        }
        return
}

func mxRecord(query string) {
        mxRecord, err := net.LookupMX(query)
        color.Cyan("\n[+] MX Records\n")
        for i := 0; i < len(mxRecord); i++ {
                fmt.Printf("Host : %s Priority : %d \n", mxRecord[i].Host, mxRecord[i].Pref)
        }

        if err != nil {
                //panic(err)
        }
        return
}

func txtRecord(query string) {
        txtRecord, err := net.LookupTXT(query)
        color.Cyan("\n[+] TXT Record(s)\n")
        for i := 0; i < len(txtRecord); i++ {
                fmt.Printf("#%d : %s \n", i, txtRecord[i])
        }
        if err != nil {
                //panic(err)
        }
        return
}

func srvRecord(query string) {
        color.Cyan("\n[+] SRV Record(s)\n")
        services := [...]string{"sipfederationtls", "autodiscover", "tls", "_sip", "xmpp-server", "VLMCS"}
        for _, service := range services {
                cname, addrs, err := net.LookupSRV(service, "tcp", query)
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
                //for _, record := range addrs {
                //fmt.Printf("Target: %s:%d\n",cname, record.Target, record.Port,)
                //}
        }
        return
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
        return
}

func main() {
        //Query argument (./digall <arg>)
        if len(os.Args) > 1 {
                // input argument
                query := os.Args[1]
                // variable for A record subdomain
                www := fmt.Sprintf("www.")
                //Run whois query
                color.Green("[ DIGALL ]")
                color.Yellow("\n[+] Checking whois information")
                color.Yellow("[!] Currently only works for .no domains\n")
                whoisQuery(query)
                //Run query functions with info
                color.Yellow("\n[+] Starting DNS queries\n")
                aRecord(query)
                color.Magenta("\n[!] with www. :")
                aRecord(www+query)
                cnameRecord(www+query)
                mxRecord(query)
                txtRecord(query)
                srvRecord(query)
                fmt.Printf("\n")
                nsRecord(query)
                color.Green("\n[+] Done\n")

                } else {
                color.Red("[-] ERROR: Not a valid argument\n")
                color.Yellow("[!] Usage: ./digall <domain>\n[!] QUITTING")
        }
}
