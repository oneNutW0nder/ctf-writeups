# Fleshwound - 200
## Solved by: DeadlyKitten
## 71 Solves

> Tis but a....

- If we continue the information give in the challenge description _Tis but a.... scratch_ we get a massive hint! This turns out to be a JSON file that is a SCRATCH program! We can go upload this to the scratch site online and see what it is doing.

![](Pasted%20image%2020210419141455.png)

- This is what we see when we upload the JSON file. I don't know about you but Scratch is weird. There is separate code for each sprite and each sprite has stages. So we click on the _drum_ sprite and see more code, then we select the stages for the _drum_ sprite:

![](Pasted%20image%2020210419141725.png)

- This is the resulting code that we get at this point:

![](Pasted%20image%2020210419141754.png)

- I am no scratch expert but essentially you solve this challenge by modifying a function call in this block:

![](Pasted%20image%2020210419141819.png)

- At the bottom of this block the code calls the function `distract_me`. If we change this to call the `finish` function instead and execute the program we can get the flag! 

![](Pasted%20image%2020210419141901.png)

- Make sure to give the input `gib flag` which can be derrived from the top of the block in above picture. The flag will then be printed above the _drum_ sprite in the top right corner:

![](Pasted%20image%2020210419142059.png)