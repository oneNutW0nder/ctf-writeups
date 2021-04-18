# Birdthief Downlink - 100

## Solved by: 1cysw0rdk0
## 3 Solves

> Read the slide deck for more information

- The _Birdthief_ series of challenges revolved around a PDF document that was a briefing about a drone that crash landed and needed to be recovered. All of these challenges revolved around reading this briefing for more information and subtle hints about the challenges.

- For this challenge the briefing has the following information:

![](Pasted%20image%2020210417224754.png)

- As the slide suggests, there is something wrong with the data provided. This suggests that there is some sort of parity or checksum going on with the data. We considered Hamming codes for a time, but that did not lead anywhere. In the end, this challenge was simply checking for a parity bit. The following is our script to get the flag:

```python
def checkParity(byte):
    parity = True
    for bit in byte:
        if bit == '1':
            parity = not parity

    return parity


for line in open('temp.txt'):
    for byte in line.split():
        if(checkParity(byte)):
            print(byte[:-1])
```

- First we had to take the challenge file and get all 8 bits per each byte in the corrupted data. This gives us `tmp.txt` which the above script operates on. This script will output the non-corrupted bytes which you can then convert back to ASCII to get the flag:

![](Pasted%20image%2020210417230230.png)

