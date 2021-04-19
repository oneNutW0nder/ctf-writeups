# Inception Dream 4 - 200
## Solved by: oneNutW0nder
## 162 Solves

> Note: This challenge builds off of InceptionCTF: Dream 3
Don’t lose yourself within the dreams, it’s critical to have your totem. Take a close look at the facts of the file presented to you. Please note the flag is marked with an “RITSEC=” rather than {} due to encoding limitations.
The flag will need to be converted to RITSEC{}

- This challenge gives us a `PasswordPexe.hta` file. The first thing that I did was run `strings` on it and saw the following:

![](Pasted%20image%2020210419144050.png)

- This JavaScript is performing a `document.write()` with a bunch of URL encoded data. This means it is writing HTML most likely to a web page. If we pull out this data and create a file from the HTML we get the following:

![](Pasted%20image%2020210419144200.png)

- Being the hackers that we are, we don't care about what the page looks like so we check the source:

![](Pasted%20image%2020210419144229.png)