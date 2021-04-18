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
