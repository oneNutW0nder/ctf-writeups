# Memedrive - 275
## Solved by: DeadlyKitten
## 31 Solves

> The best Android app for viewing memes!

- This challenge provided an APK file. This is an Android application that we can unpack and decompile using `apktool d memedrive.apk`.  This will unpack everything and allow us to start digging through the source code. 

- I am no expert on mobile applications or Java but something in this application was fishy. First there were a bunch of memes of course in the `/resources/res/drawable` path that were entertaining. There was also this picture included with them:

![](Pasted%20image%2020210419135330.png)

- This is a hint for the challenge as we will see later on. If we go have a look at the `InitStuff.java` file we can see some interesting things:

![](Pasted%20image%2020210419135434.png)

- The first thing is that this function is called `doInBackground()` which is always interesting. Second, it handles writing what looks like usernames and salted passwords to a file called `etc-password`. It also writes some zeros to the files `index.log` and `key.log`. This very interesting because we already know that XOR is present in this challenge and the file is called `key.log` so we will keep that in mind. Another interesting ting is this Github link. It takes us to a file called `flags.txt` and is a massive file of what looks like randome garbage:

![](Pasted%20image%2020210419135704.png)

- The last piece of the puzzle lies in the file `ShowingX.java` where `X` is the number for a picture. These files look like this:

![](Pasted%20image%2020210419135748.png)

- See anything familiar? Each of these files opens up `index.log` and `key.log` and writes some data into them! What if `index.log` contains the indexes for the important lines in `flags.txt` and `key.log` contains possible keys to XOR against those lines! If we do this and Base64 decode the `flags.txt` lines and XOR them with the keys from `key.log` we get the flag! Below is the script used to solve this challenge:

```python
# Pulled from APK
# https://raw.githubusercontent.com/yung-g4ngst3r360noscope/gimmeThatData/main/flags.txt
import base64

keys = ['LN3M99BX', 'PQ4LL22C', 'WV94J7ZH', 'GT6BW30K', 'N28UB11M', 'FQ7FV5K2', '7NFSK27C', '1CZ7KHR5', 'NN52DOMW', 'JWVE66PI']

with open('flags.txt', 'r') as f:
    flags = f.readlines()

for l in flags:
    for k in keys:
        dec = base64.b64decode(l)
        xored = []
        for i in range(len(dec)):
            xored_value = dec[i] ^ ord(k[i%len(k)])
            xored.append(chr(xored_value))
        meep = ''.join(xored)
        if meep.startswith('RS'):
            print(meep)
```