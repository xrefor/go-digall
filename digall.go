
package main

import (
    "fmt"
    "github.com/briandowns/spinner"
    "github.com/fatih/color"
    "net"
    "os"
    "time"
)

// DNSRecordFetcher holds methods to fetch various DNS record types.
type DNSRecordFetcher struct {
    domain string
    s      *spinner.Spinner
}

// NewDNSRecordFetcher initializes a DNSRecordFetcher for the specified domain with a spinner.
func NewDNSRecordFetcher(domain string) *DNSRecordFetcher {
    s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Use a clean spinner character set
    s.Prefix = "Loading... "
    return &DNSRecordFetcher{domain: domain, s: s}
}

// FetchARecords fetches A records (IPv4 addresses) for the domain.
func (d *DNSRecordFetcher) FetchARecords() {
    d.s.Start()
    if records, err := net.LookupIP(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgGreen).Println("A Records:")
        for _, ip := range records {
            if ip.To4() != nil {
                fmt.Println(ip)
            }
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchAAAARecords fetches AAAA records (IPv6 addresses) for the domain.
func (d *DNSRecordFetcher) FetchAAAARecords() {
    d.s.Start()
    if records, err := net.LookupIP(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgCyan).Println("AAAA Records:")
        for _, ip := range records {
            if ip.To16() != nil && ip.To4() == nil {
                fmt.Println(ip)
            }
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchMXRecords fetches MX records for the domain.
func (d *DNSRecordFetcher) FetchMXRecords() {
    d.s.Start()
    if records, err := net.LookupMX(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgYellow).Println("MX Records:")
        for _, mx := range records {
            fmt.Printf("%s %d\n", mx.Host, mx.Pref)
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchNSRecords fetches NS records for the domain.
func (d *DNSRecordFetcher) FetchNSRecords() {
    d.s.Start()
    if records, err := net.LookupNS(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgBlue).Println("NS Records:")
        for _, ns := range records {
            fmt.Println(ns.Host)
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchTXTRecords fetches TXT records for the domain.
func (d *DNSRecordFetcher) FetchTXTRecords() {
    d.s.Start()
    if records, err := net.LookupTXT(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgMagenta).Println("TXT Records:")
        for _, txt := range records {
            fmt.Println(txt)
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchCAARecords fetches CAA records for the domain.
func (d *DNSRecordFetcher) FetchCAARecords() {
    d.s.Start()
    if records, err := net.LookupTXT(d.domain); err == nil && len(records) > 0 {
        d.s.Stop()
        color.New(color.FgHiYellow).Println("CAA Records:")
        for _, caa := range records {
            fmt.Println(caa)
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchCNAME fetches the CNAME record for the domain.
func (d *DNSRecordFetcher) FetchCNAME() {
    d.s.Start()
    if cname, err := net.LookupCNAME(d.domain); err == nil && cname != d.domain {
        d.s.Stop()
        color.New(color.FgHiGreen).Println("CNAME Record:", cname)
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchPTRRecords fetches PTR records (reverse DNS lookup) for the domain.
func (d *DNSRecordFetcher) FetchPTRRecords() {
    d.s.Start()
    ips, err := net.LookupIP(d.domain)
    if err == nil && len(ips) > 0 {
        d.s.Stop()
        color.New(color.FgRed).Println("PTR Records:")
        for _, ip := range ips {
            if names, err := net.LookupAddr(ip.String()); err == nil && len(names) > 0 {
                for _, name := range names {
                    fmt.Println(name)
                }
            }
        }
        fmt.Println()
    } else {
        d.s.Stop()
    }
}

// FetchAllRecords fetches all DNS records and displays them if present.
func (d *DNSRecordFetcher) FetchAllRecords() {
    fmt.Printf("Fetching DNS records for domain: %s\n\n", d.domain)
    d.FetchARecords()
    d.FetchAAAARecords()
    d.FetchMXRecords()
    d.FetchNSRecords()
    d.FetchTXTRecords()
    d.FetchCAARecords()
    d.FetchCNAME()
    d.FetchPTRRecords()
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go-digall <domain>")
        os.Exit(1)
    }
    domain := os.Args[1]
    fetcher := NewDNSRecordFetcher(domain)
    fetcher.FetchAllRecords()
}
