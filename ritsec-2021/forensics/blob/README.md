# Blob - 200
## Solved by: DeadlyKitten
## 141 Solves

> Ha. Blob. Did you get the reference?

- This is another git based challenge. After cloning the repository, we had a pretty good idea that it had something to do with git objects/blobs. I will let the reader do further research on their own into this matter as we will not be discussing it here.

- Basically we want to list all _git objects_ like this:

![](Pasted%20image%2020210419150753.png)

- Now we can take the ID of the `flag.txt` file and get the flag by printing the contents of each file object:

![](Pasted%20image%2020210419151021.png)