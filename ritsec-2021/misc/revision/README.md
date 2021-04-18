# Revision - 200
## Solved by: oneNutW0nder
## 75 Solves

> ... They aren’t necessarily obvious but are helpful to know.

- Cloning the given repository gives us the following files:

![](Pasted%20image%2020210417222255.png)

- Nothing pops out immediately so that probably means there is something going on in the history. If we inspect the `git log` there are a ton of commits.

![](Pasted%20image%2020210417223056.png)

- When inspecting large histories like this I like to use some sort of graphical interface for viewing git logs. I prefer the _git graph_ view in _Visual Studio Code_. This makes it much easier to see what is happening between branches and commits. 

![](Pasted%20image%2020210417223405.png)

- If we scroll down in the _git graph_ we can see some suspicious commits done my `knif3` who is the author of the challenge:

![](Pasted%20image%2020210417223529.png)

- Viewing the file that these commits modify we can see `flag.txt`:

![](Pasted%20image%2020210417223638.png)

- Unfortunately, each commit shows the removal and addition of a single character of the flag:

![](Pasted%20image%2020210417223839.png)

- Being that I could not think of a fast way to script this as I am not all that familiar with git commands, I just went through and viewed the file history of each commit and slowly built the flag character by character.

`RS{I_h0p3_y0u_scr1pted_th0s3_git_c0ms}`