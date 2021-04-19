# In Scope - 382
## Solved by: _Team Effort_
## 7 Solves

> Read the slide deck for more information

- This was a cool challenge! It was part of the _Birdthief_ briefing packet which gave some hints as to what this `.wav` file was:

![](Pasted%20image%2020210419142802.png)

- There is a heavy use of italics around the word _scope_. Using `binwalk` and other basic stego techniques on the file yielded no results. Then it hit me! What if this is an oscilliscope waveform!!! It turns out that you can actually encode an image in a `.wav` file for an oscilliscope. 

- We found this site: https://dood.al/oscilloscope/ which allows you to upload a `.wav` file and play it for the oscilliscope. It took some tinkering with the site and some time to decipher the flag but in the end you can watch the flag be drawn out for you using the above mentioned site.