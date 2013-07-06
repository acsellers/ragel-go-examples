# [Ragel](http://www.complang.org/ragel/) Golang Examples

As I can't find good examples for Ragel code in Go1 style, and
I'm going to be implementing a bunch of different parser type 
things to learn Ragel, I decided to go ahead and put my practice 
code up on Github for the next person to learn from (hopefully).
I'm not associated with the people behind Ragel in any form.

## Current Examples

All of the current code is 100% what I wrote excepting the Atoi, 
which was from an official ragel example, but was modified to be 
current Go style. I'm not sure how many examples I'm going to 
write, but current goal is 50. The code I use for ragel is 

    ragel -Z -G2 -o <name>.go <name>.rl

* Atoi - Convert an integer in a string to an int64
* Atof - Convert a float in a string to a float64
* HexDecode - Convert a hex string into an int64
* Comma - Separate a string into values like a csv

## Licensing

Any code is available under the 2-clause BSD License, a copy
of which is included as the file LICENSE.
