# Parcel - 200
## Solved by: oneNutW0nder
## 176 Solves

> That's a lot of magick

- This challenge gave a file that contained the HTTP body from many file uploads. All of the data was Base64 encoded as well. All of this data turns out to be many different images. My solution was the unintended one and actually a consequence of ImageMagik embedding metadata into the images.

- My solution was to decode a block and I was given the flag:

![](Pasted%20image%2020210419175907.png)

- The real solution involved decoding all of the data into individual images and piecing them together like a puzzle to get the flag!