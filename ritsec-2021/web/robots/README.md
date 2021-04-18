# Robots - 100
## Solved by: oneNutW0nder
## 500 Solves

> Robots are taking over. Find out more.

- Here is the web page that we are presented with upon reaching the challenge site:

![](Pasted%20image%2020210417194943.png)

- Due to the name of the challenge being _robots_ it is a pretty good assumption that the flag will either be in or have something to do with the `/robots.txt` file. Navigating to this file gives us a large amount of paths that are allowed and disallowed. 

![](Pasted%20image%2020210417195059.png)

- Using `CTRL+F` to quickly search the page for `flag` we find this entry:

![](Pasted%20image%2020210417195129.png)

- The path looks like it is Base64 encoded. Decoding it yields the flag!

![](Pasted%20image%2020210417195247.png)