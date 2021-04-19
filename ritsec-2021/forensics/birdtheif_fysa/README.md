# Birdthief FYSA - 100
## Solved by: oneNutW0nder
## 254 Solves

> No description

- This challenge was the first in a series that deals with a briefing document about a crashed drone. This briefing contains the hints and information for all of the challenges in the series.

- For this challenge we can see that we need to _get as much as we can_ from this briefing.

![](Pasted%20image%2020210419145344.png)

- Running `strings` doesn't get us anywhere so I decided to run `binwalk`. This will extract all of the attachments and embedded files within the PDF document. After this one of the extracted images has the flag for us!

![](Pasted%20image%2020210419145439.png)