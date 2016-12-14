package main

import (
        "fmt"
        "os"
        "net"
)

//Lookup functions
func aRecord(query string) {
    ipRecord, err := net.LookupIP(query)
    fmt.Printf("\n[+] A Record\n")
    for i := 0; i < len(ipRecord); i++ {
        fmt.Printf("IP : %s \n", ipRecord[i])
    }
        
        if err != nil {
                //panic(err)
        }
}

func cnameRecord(query string) {
    cnameRecord, err := net.LookupCNAME(query)
    fmt.Printf("\n[+] CNAME Records\n")
    fmt.Println(cnameRecord, err)
    if err != nil {
        //panic(err)
    }
}

func mxRecord(query string) {
        mxRecord, err := net.LookupMX(query)
        fmt.Printf("\n[+] MX Records\n")
        for i := 0; i < len(mxRecord); i++ {
            fmt.Printf("Host : %s Priority : %d \n", mxRecord[i].Host, mxRecord[i].Pref)
        }

        if err != nil {
           //panic(err)
        }
}

func txtRecord(query string) {
        txtRecord, err := net.LookupTXT(query)
        fmt.Printf("\n[+] TXT Record(s)\n")
        for i := 0; i < len(txtRecord); i++ {
            fmt.Printf("#%d : %s \n", i, txtRecord[i])
        }
        
        if err != nil {
                //panic(err)
        }
}

func srvRecord(query string) {
        fmt.Printf("\n[+] SRV Record(s)\n")
        services := [...]string{"sipfederationtls", "autodiscover", "tls", "_sip", "xmpp-server", "VLMCS"}
        for _, service := range services {
            cname , addrs, err := net.LookupSRV(service, "tcp", query)
            
            if err != nil {
                    //panic(err)
            }
            for _, record := range addrs {
                fmt.Printf("Target: %s:%d\n",cname, record.Target, record.Port)
            }
        }

}

func nsRecord(query string) {
        nsRecord, err := net.LookupNS(query)
        fmt.Printf("\n[+] NS Records\n")
        for i := 0; i < len(nsRecord); i++ {
            fmt.Printf("NS : %s \n", nsRecord[i].Host)
        }
        
        if err != nil {
                //panic(err)
        }
}

func main() {
        //Query argument (./digall <arg>)
        if len(os.Args) > 1 {
                // input argument
                query := os.Args[1]
                // variable for A record subdomain
                www := fmt.Sprintf("www.")
                //Run query functions with info
                fmt.Printf("\n[ digall: Starting DNS queries ]\n")
                aRecord(query)
                fmt.Printf("\n[!] with www.")
                aRecord(www+query)
                cnameRecord(www+query)
                mxRecord(query)
                txtRecord(query)
                srvRecord(query)
                fmt.Printf("\n")
                nsRecord(query)
                fmt.Printf("\n[!] Done\n")

                } else {
                fmt.Print("ERROR: Not a valid argument\n")
                fmt.Println("Usage: ./digall <domain>\n[!] QUITTING")
        }
}
