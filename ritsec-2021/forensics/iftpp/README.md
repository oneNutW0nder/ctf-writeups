# IFTPP - 500
## Solved by: Scuzz3y
## 23 Solves

> Dang that's a big ping

- This challenge was very unique. The challenge gives us a PCAP file which contains one HTTP conversation and a lot of ICMP messages. If we look at the HTTP traffic, we can see a file `rfc.txt` was requested:

![](Pasted%20image%2020210419174848.png)

- If copy and paste the HTTP response into a text editor for easier reading we can see that this is actually an RFC for the _Insecure File Transfer Protocol over Ping_. This is not a real RFC which is obvious almost immediately from reading the RFC. The goal of the challenge became very obvious at this point. This PCAP contains a file transfer using IFTPP. Since we have the RFC we can read it in order to figure out how to extract the file data and presumably get the flag.

- The RFC contains some snippets of Go code for things like the _shared key generation_. It also references the use of _protobuf_ which is a serialization protocol. The RFC also provides us with the protobuf structure. Since we have some snippets of code already, we wrote our solution in Go which can be found in `main.go`. It parses the PCAP and the file data. The resulting file is a JPG which is the flag!

