# Answers - 100
## Solved by: 1cysw0rdk0 
## 54 Solves

> Lookup _this_
> `answers.ritsec.club:53/udp`

- This challenge was very interesting. The challenge info immediately screams DNS look ups. However, doing a lookup for this domain yields the following:

![](Pasted%20image%2020210417230603.png)

- This shows that `answers.ritsec.club` is a CNAME record for `challenges1.ritsec.club`. We queried for many other record types but could not find anything interesting. We even attempted a zone transfer and that obviously didn't work either. However, if we lookup `answers.ritsec.club` by sending our query to `challenges1.ritsec.club` we finally get some interesting results:

![](Pasted%20image%2020210417230933.png)

- If we continue _digging_ through these records, we keep getting new random CNAME records:

![](Pasted%20image%2020210417231057.png)

- It became apparent very quickly that we need to script this in order to get the flag. The following the script that we used:

```python
#!/bin/env python3
# Requests the cname records using dig
# Adds all cnames to a list
# coninues to add cname records to list until one of them is a txt record, and stands out from the rest
# dig txt request that record
# flag

import subprocess

cnames = []

answer = subprocess.check_output(['dig', 'answers.ritsec.club', '@answers.ritsec.club'])

flag = False
for line in answer.splitlines():
    if (line != b';; ANSWER SECTION:' and not flag):
        continue
    
    if (line == b''):
        break
    flag = True

    if (line == b';; ANSWER SECTION:'):
        continue
    cnames.append(line.split()[-1])


while (True):

    host = cnames.pop()

    print(b'Requesting ' + host)
    answer = subprocess.check_output(['dig', host, '@answers.ritsec.club'])
    flag = False
    for line in answer.splitlines():
        if (line != b';; ANSWER SECTION:' and not flag):
            continue
    
        if (line == b''):
            break
        flag = True

        if (line == b';; ANSWER SECTION:'):
            continue

        if (line.split()[-1] == b'answers.ritsec.club.'):
            continue
        cnames.append(line.split()[-1])
        continue

print('finished')
```


- The script will continue following the records and printing the results. There will eventually be a record that obviously has a different name:

![](Pasted%20image%2020210417231627.png)

- This record name also indicates that there might be a _random txt record_. With this information we can request a TXT record from this domain to get the flag:

![](Pasted%20image%2020210417232201.png)