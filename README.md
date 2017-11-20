# go-digall

My first golang project learning Golang. Does domain lookup of A, MX, CNAME, TXT, SRV and NS records, all in one go.

example usage: digall github.com
[+] Starting DNS queries
[+] A Records
[*] github.com
IP : 192.30.253.112 
IP : 192.30.253.113 

[*] www.github.com
IP : 192.30.253.112 
IP : 192.30.253.113 

[+] CNAME Record(www.github.com)
github.com.
[+] MX Records
Host : aspmx.l.google.com. Priority : 1 
Host : alt1.aspmx.l.google.com. Priority : 5 
Host : alt2.aspmx.l.google.com. Priority : 5 
Host : alt4.aspmx.l.google.com. Priority : 10 
Host : alt3.aspmx.l.google.com. Priority : 10 

[+] TXT Record(s)
#0 : docusign=087098e3-3d46-47b7-9b4e-8a23028154cd 
#1 : v=spf1 ip4:192.30.252.0/22 ip4:208.74.204.0/22 ip4:46.19.168.0/23 include:_spf.google.com include:esp.github.com include:_spf.createsend.com include:mail.zendesk.com include:servers.mcsv.net ~all 

[+] SRV Record(s)

[+] NS Records
NS : ns-1707.awsdns-21.co.uk. 
NS : ns-520.awsdns-01.net. 
NS : ns1.p16.dynect.net. 
NS : ns2.p16.dynect.net. 
NS : ns4.p16.dynect.net. 
NS : ns-421.awsdns-52.com. 
NS : ns-1283.awsdns-32.org. 
NS : ns3.p16.dynect.net.

