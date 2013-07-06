# HexDecode

HexDecode will take a string and decode it into an uint64. It will not 
try to check if the input string will overrun the bounds of an int64.
There is an optional string of '0x' before the hex digits begin, it has
no effect on the result, but must not cause an error. The hex digits 
are the digits 0-9 and the letters a-f in upper or lower case.
