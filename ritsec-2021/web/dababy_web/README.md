# Dababy - 150
## Solve by: oneNutW0nder
## 213 Solves

> Dababy wanted to share a message, but he seemed to put it too high up...

- Here is the website that we are presented with when going to the provided link.

![](Pasted%20image%2020210417194010.png)

- There are two hyperlinks on the page which bring us to interesting pages. The first link brings us to `/fun.php` which is shown below.

![](Pasted%20image%2020210417194238.png)

- If we enter a name the page will tell us that it is a cool name. We can also see that the page now has a parameter called `string` which is set to the name that we entered.

![](Pasted%20image%2020210417194446.png)

- Because of the low point value of the challenge, the fact that it is PHP, and the fact that the web application is essentially _echoing_ our name back to us, it is logical to try command injection on this page. First we will just try simple command injection with a semicolon: `; ls`

![](Pasted%20image%2020210417194630.png)

- We can see that this worked! `; ls` returned the contents of our current directory which is `dababy.sh`. We can use this to search for the flag! For some reason when completing this write up the `flag.txt` no longer exists. However, if it did exist, you could find it in `../flag.txt`. If we use this command injection to try and read the flag file we get the following error:

![](Pasted%20image%2020210418002706.png)

- It appears that we cannot use the `cat` command here. Let's visit the other link that was on the home page:

![](Pasted%20image%2020210418002749.png)

- On this page `/fun1.php` we can see that we are accessing a file in the GET parameters: 

![](Pasted%20image%2020210418002823.png)

- This is likely a sign that there is a Local File Inclusion (LFI) vulnerability. Let's try accessing `../../../../../etc/passwd`:

![](Pasted%20image%2020210418003206.png)

- The screen shot above shows the source code of the web page after trying to access `../../../../../etc/passwd` because it makes it much easier to see the results.

- This vulnerability allows us to read arbitrary files on the file system! Now we can read the flag at `../flag.txt`. Again, I am unable to show the flag because it does not exist in the post competition challenge for some reason.