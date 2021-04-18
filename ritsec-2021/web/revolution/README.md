# Revolution - 250
## Solved by: oneNutW0nder
## 29 Solves

> The robots are taking over. They are posting their propaganda everywhere. Go here to find out more about it. 
> You will need a valuable piece of information from the 'robots' challenge.

> All the other information you should need is on the root page. You need to craft a message to send to the leaders of the revolution to let them know your safe to talk to. They have made it sort of cryptic so that not everyone can access it. You will need to use the words on the page in order to figure out how to get a response from them.

> They expect a special type of request and only have the ability to read plain text from a special agent. ONLY SEND PLAIN TEXT DATA.

- First, I must preface this challenge by saying that this challenge was difficult in that it required some obscene guess work. The solution was not intuitive nor did it follow any common web vulnerabilities or miscommunications. 

- Moving into the challenge, we are presented with a web page that talks about four things that robots claim they are:

- Robots are _Friendly_

![](Pasted%20image%2020210417200219.png)

- Robots are _Caring_

![](Pasted%20image%2020210417200311.png)

- Robots have _Laws_

![](Pasted%20image%2020210417200340.png)

- Robots _Protect_

![](Pasted%20image%2020210417200357.png)

- At the bottom of the page there is some more text about the _revolution_

![](Pasted%20image%2020210417200414.png)

- Based on this text at the bottom, we are need to send a special _crafted_ message to the server and it will presumably give us the flag in response to the correct message. To start we remember the challenge text that says _you will need a valuable piece of information from the robots challenge_. Recall from the _robots_ challenge that there was a `User-Agent` string at the top of the `/robots.txt` file.

![](Pasted%20image%2020210417200911.png)

- Unfortunately, sending a request to the challenge page with this `User-Agent` set does not give us the flag. It also does not change the page at all. Now for some guessing. At this point I assumed that there was a different endpoint that would accept the request with the `User-Agent`. This assumption, along with the challenge info telling me to look and read carefully, I started sending requests to endpoints that were words on the page. I tried sending requests to endpoints like `/laws`, `/caring`, `/protect`, and `/friendly`. None of these endpoints yielded anything. I finally sent a request to the `/revolution` endpoint which responded with a `405 Method Not Supported` response. This means that the server understood my request, the endpoint exists, but the `GET` method is not supported. Since `GET` is not supported, I sent an `OPTIONS` request to see what the server will allow for that endpoint. 

![](Pasted%20image%2020210417201723.png)

- Sending the `OPTIONS` request:

![](Pasted%20image%2020210417201810.png)

- We can see that the server responded with the allowed _methods_ for the `/revolution` endpoint. Once of them is not part of the HTTP RFC and that is `UNLOCK`. Using this method along with the same user agent, we get the following response:

![](Pasted%20image%2020210417201949.png)

- Some trolling by the challenge author here because the web server responded with a `200 OK`  message, but the HTML returned says `404 ;)`.  Unfortunately, this is where even more guessing comes into play. It wasn't until after some time of asking the challenge author questions and reading each hint as it came out that the solution revealed itself. In order to get the flag, you must send an `UNLOCK` request with the user agent found from the `robots` challenge and the four pieces of propaganda in the request body. However, `cURL` will automatically add the `Content-Type` header if a message body is present. It will set the content type to `application/x-www-form-urlencoded` which will not give you the flag.

![](Pasted%20image%2020210417202352.png)

- Finally, we find out that we need to send _plain text data_ to the server. This means that we need to change the `Content-Type` header to by `text/plain` instead of form data. Doing this yields us the flag!

![](Pasted%20image%2020210417202512.png)
