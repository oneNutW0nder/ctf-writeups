# Sessions - 100
## Solved by: oneNutW0nder
## 302 Solves

> Find the flag.

- Due to the low point value of this challenge and the name of it, I immediately assume that the flag will have something to do with the session token/cookie of the web application. Going to the challenge page we see the following:

![](Pasted%20image%2020210417202717.png)

- Trying some default credentials like `admin:admin` does not yield anything of interest. Usually I check the source of the web page to see if there is anything interesting hidden there.

![](Pasted%20image%2020210417202839.png)

- Sure enough! Some developer left us some juicy credentials in the comments of the HTML: `iroh:iroh`. When we log in we see the following:

![](Pasted%20image%2020210417202929.png)

- This looks like it is a tribute to Iroh from _The Last Airbender_. Now that we are logged in let's take a look at the session token/cookie. During the CTF I used BurpSuite to view the requests, but for the write up we will use the JavaScript console in the Firefox browser.

- Viewing the cookie with `document.cookie` gives the following:

![](Pasted%20image%2020210417203134.png)

- This looks like Base64 again! Let's decode it for the flag!

![](Pasted%20image%2020210417203247.png)