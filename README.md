# go-digall

My first golang project learning Golang. Does domain lookup of A, MX, CNAME, TXT, SRV and NS records, all in one go.

example usage:
digall google.com
```
[+] Starting DNS queries
[+] A Records
[*] google.com
IP : 2a00:1450:4001:81f::200e 
IP : 172.217.23.174 
[*] www.google.com
IP : 2a00:1450:4009:80c::2004 
IP : 216.58.204.4 

[+] CNAME Record(www.google.com)
www.google.com.

[+] MX Records
Host : aspmx.l.google.com. Priority : 10 
Host : alt1.aspmx.l.google.com. Priority : 20 
Host : alt2.aspmx.l.google.com. Priority : 30 
Host : alt3.aspmx.l.google.com. Priority : 40 
Host : alt4.aspmx.l.google.com. Priority : 50 

[+] TXT Record(s)
#0 : v=spf1 include:_spf.google.com ~all 

[+] SRV Record(s)
addrs[0].Target : xmpp-server.l.google.com. 
addrs[0].Port : 5269 
addrs[0].Priority : 5 
addrs[0].Weight : 0 
cname : _xmpp-server._tcp.google.com. 
addrs[1].Target : alt3.xmpp-server.l.google.com. 
addrs[1].Port : 5269 
addrs[1].Priority : 20 
addrs[1].Weight : 0 
cname : _xmpp-server._tcp.google.com. 
addrs[2].Target : alt4.xmpp-server.l.google.com. 
addrs[2].Port : 5269 
addrs[2].Priority : 20 
addrs[2].Weight : 0 
cname : _xmpp-server._tcp.google.com. 
addrs[3].Target : alt2.xmpp-server.l.google.com. 
addrs[3].Port : 5269 
addrs[3].Priority : 20 
addrs[3].Weight : 0 
cname : _xmpp-server._tcp.google.com. 
addrs[4].Target : alt1.xmpp-server.l.google.com. 
addrs[4].Port : 5269 
addrs[4].Priority : 20 
addrs[4].Weight : 0 
cname : _xmpp-server._tcp.google.com. 

[+] NS Records
NS : ns2.google.com. 
NS : ns4.google.com. 
NS : ns1.google.com. 
NS : ns3.google.com. 
[+] Done
```
