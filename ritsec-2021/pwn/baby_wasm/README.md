# Baby WASM
## Solved by: oneNutW0nder
## 5 Solves

> Baby WASM
This is a JS engine pwn challenge.
V8 version: 8.7.220
V8 commit: 0d81cd72688512abcbe1601015baee390c484a6a
glibc version: 2.31
OS: Ubuntu 20.04.2 - 64 bit
Provided challenge files: V8 debug and release builds, diff and patch files, challenge setup files, libc.so.6
The server will be running the release build.

- First, I would like to point out that this challenge had an unintended solution. If you wish to read a write up that involves the intended solution for this challenge please go to: https://github.com/ducphanduyagentp/ctf-challs

- When I saw teams solve this challenge super fast from the start of the CTF I became interested to see if there was an unintended solution. When connecting the server, we are prompted for a URL to our JavaScript payload. The challenge server will then go out to the supplied link, pull down our JavaScript, and execute it. Obviously the intended solution involves exploiting the V8 JS engine, but that is way beyond me. Knowing that there might be a really simple solution I began browsing for ways to read or execute system commands with V8/JavaScript. I eventually stumbled upon this link that shows how to read and print file contents from the file system: https://v8.dev/docs/d8

- Using the information on this link I created the following script:

```js
const flag = read("flag.txt")
console.log(flag)
```

- I hosted this on a public server so that the challenge server could reach it and provided its URL. Luckily I was given the flag!

![](Pasted%20image%2020210418000756.png)