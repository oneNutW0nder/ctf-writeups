# GandCrab Ursnif - 200
## Solved by: seans4099
## 117 Solves

> NOTE: this challenge builds upon BegineersRITSEC.html, and that challenge must be completed first.
GandCrab/Ursnif are dangerous types of campaigns and malware, macros are usually the entry point, see what you can find, there are two flags in this document. Flag1/2

- The entire series of these challenges were very strange and some were broken up until the end of the competition. This challenge however, builds off one of the documents that was attached to the Outlook email message from the _RITSEC Beginners HTML_ challenge. 

- First, it is important to remember that `.doc` files are really just archives. This means you can do `unzip MY_DOCUMENT.doc` and get a bunch of files out. This includes all of the VB macros that may or may not included in the document:

![](Pasted%20image%2020210419184715.png)

- Again, I like running strings because it catches some juicy things sometimes:

![](Pasted%20image%2020210419184804.png)

