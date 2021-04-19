# IceID - 350
## Solved by: oneNutW0nder
## 109 Solves

> NOTE: this challenge builds upon BegineersRITSEC.html, and that challenge must be completed first.
Stepping it up to IceID/Bokbot, this challenge is like the previous challenge but requires some ability to read and understand coding in addition to some additional decoding skills, there are two flags in this document. (Flag 1/2)

- The entire series of these challenges were very strange and some were broken up until the end of the competition. This challenge however, builds off one of the documents that was attached to the Outlook email message from the _RITSEC Beginners HTML_ challenge. 

- First, it is important to remember that `.doc` files are really just archives. This means you can do `unzip MY_DOCUMENT.doc` and get a bunch of files out. This includes all of the VB macros that may or may not included in the document:

![](Pasted%20image%2020210419184310.png)

- Considering this challenge references malware based on office macros we should probably start there. I like to run strings on things right away because it catches a lot of low hanging fruit. 

![](Pasted%20image%2020210419184539.png)

- Sure enough! There is a flag that is encoded with ROT13.

![](Pasted%20image%2020210419184628.png)