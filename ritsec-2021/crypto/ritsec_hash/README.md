# RITSEC Hash - 250
## Solved by: DeadlyKitten
## 56 Solves

> Hmmm.. we found this hash along with a white paper explaining this custom hashing algorithm.

> Can you break it for us?

> hash : 435818055906

- The premise of this challenge is to read the _white paper_ containing information about the _RITSEC Hashing Algorithm_, implement it, and crack the provided hash.

- The following is our solution code:

```cpp
#include <string>
#include <vector>
#include <iostream>
#include <iomanip>
#include <fstream>

std::vector<uint8_t> h0{'R', 'I', 'T', 'S', 'E', 'C'};

std::vector<uint8_t> Run_Round(std::vector<uint8_t> Hi, uint8_t Xi, uint8_t round_num)
{
    std::vector<uint8_t> temp(6, 0);
    uint8_t a = Hi[0];
    uint8_t b = Hi[1];
    uint8_t c = Hi[2];
    uint8_t d = Hi[3];
    uint8_t e = Hi[4];
    uint8_t f = Hi[5];

    temp[0] = ((((((c ^ e) & f) + b) + (d << 2)) + Xi) + round_num);
    temp[1] = a;
    temp[2] = d << 2;
    temp[3] = b >> 5;
    temp[4] = a + f;
    temp[5] = d;

    return temp;
}

int main()
{
    std::string goal = "435818055906";
    std::ifstream ifs("rockyou.txt");
    std::string input;
    while (std::getline(ifs, input))
    {
        std::vector<uint8_t> h(h0.begin(), h0.end());

        for (uint8_t c : input)
        {
            for (uint8_t i = 0; i < 13; i++)
            {
                h = Run_Round(h, c, i);
            }
        }
        std::stringstream ss;
        ss << std::hex;
        for (uint8_t c : h)
        {
            ss << std::setfill('0') << std::setw(2) << (int)c;
        }
        if (goal == ss.str())
        {
            std::cout << input << std::endl;
        }
    }
}
```

- Once you have successfully implemented the hashing algorithm simply go through the `rockyou.txt` file, hash each password with the RITSEC hash algorithm, and check to see if the resulting hash matches the one given in the challenge info.