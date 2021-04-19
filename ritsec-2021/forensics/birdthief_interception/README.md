# Birdthief Interception - 200
## Solved by: oneNutW0nder
## 46 Solves

> Read the slide deck for more information

- Another challenge in the Birdthief series. This one handles a PCAP file of someone trying to access a drone. Our job is to figure out how this happened to get the flag:

![](Pasted%20image%2020210419145645.png)

- Viewing the PCAP in wireshark we can see some suspicious DNS traffic to a `drone` address.

![](Pasted%20image%2020210419145810.png)

- Now that we have an idea of what we are looking for, I just started going through all of the TCP stream in wireshark until I found this:

![](Pasted%20image%2020210419150225.png)

- This is the login process for the drone and the flag is Base32 encoded:

![](Pasted%20image%2020210419150336.png)