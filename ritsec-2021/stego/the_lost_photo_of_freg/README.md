# The Lost Photo of Freg - 350 
## Solved by: DeadlyKitten
## 26 Solves

> My prized possession has been corrupted! Help me recover it and get to the bottom of this!


- This was also a cool challenge. We were given this image:

![](Pasted%20image%2020210419143259.png)

- It very clearly has something wrong with it and it is our job to figure that out. Early on we had the idea that it could be a _Piet_ script embedded in the file (_Piet_ is a pixel based esolang). With this theory we proceeded in extracting the data from the image and got the following image in `meep3.png`.

- if you view this image it is a valid _Piet_ script! All that is left is to load it into an online interpreter and you get the flag! Below is the python script we used to extract the _Piet_ script:

```python
from PIL import Image

def decodeImage():
    image = Image.open('freg.png')
    pxs = image.load()
    im = Image.new(mode = "RGB", size = (35, 50))
    newPixels = im.load()
    for x in range(35):
        for y in range(50):
            newPixels[x,y] = pxs[x*4,y*3]
    im.show()
    im.save('meep.png')

if __name__ == '__main__':
    decodeImage()
```