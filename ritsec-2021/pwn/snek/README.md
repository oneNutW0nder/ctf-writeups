# Snek - 100
## Solved by: oneNutW0nder
## 316 Solves

> No step on snek

- The file `snek` given for the challenge is Python 3.7 byte code. This is essentially pseudo compiled python that can be undone. The catch here is that you need Python 3.7 or 3.8 in order to use https://github.com/rocky/python-decompile3. Using this tool yields the following script from the byte code:

```python
# decompyle3 version 3.3.2
# Python bytecode 3.7 (3394)
# Decompiled from: Python 3.8.4 (default, Jul 13 2020, 21:16:07) 
# [GCC 9.3.0]
# Embedded file name: snek.py
# Compiled at: 2021-04-08 02:24:05
# Size of source mod 2**32: 834 bytes
"""
Written for RITSEC CTF 2021
Author: knif3
Flag: RITSEC{}

TODO: Finish this challenge
"""

class d(object):

    def __init__(self, password):
        self.password = password.encode()
        self.decrypt = [97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 95, 82, 83, 123, 97, 108, 108, 95, 104, 105, 36, 36, 95, 97, 110, 100, 95, 110, 48, 95, 98, 105, 116, 51, 125]

    def eq(self, other):
        if self.password == bytes(self.decrypt):
            print('!flag')
            return True
        return False


x = input('Enter my name: ')
a = d(x)
if a == x:
    print('IS_THIS_THE_FLAG??')
    print('NOPE')
else:
    print('WRONG')
# okay decompiling snek.pyc
```

- Immediately we can see that `self.decrypt` is an array of numbers that all fall within the printable ASCII range. We can quickly see what characters are in this array by adding the following python code directly after the `self.decrypt = []` line and running the script:

```python
for val in self.decrypt:
	print(chr(val), end="")
```

![](Pasted%20image%2020210418000028.png)