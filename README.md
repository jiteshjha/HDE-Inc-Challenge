# HDE-Inc-Challenge

## Description

### Challenge #1


We want to calculate a sum of squares of some integers, excepting negatives
The first line of the input will be an integer N (1 <= N <= 100)
Each of the following N test cases consists of one line containing an integer X (0 < X <= 100),
followed by X integers (Yn, -100 <= Yn <= 100) space-separated on the next line

For each test case, calculate the sum of squares of the integers excepting negatives,

and print the calculated sum to the output. No blank line between test cases

(Take input from standard input, and output to standard output)


### Challenge #2

Then, make an HTTP POST request to the URL http://hdegip.appspot.com/challenge/003/endpoint 
which contains the JSON string as a body part.

Content-Type: of the request must be "application/json".
The URL is protected by HTTP Basic Authentication, which is explained on Chapter 2 of RFC2617, so you have to provide an Authorization: header field in your POST request
For the "userid" of HTTP Basic Authentication, use the same email address you put in the JSON string.
For the "password", provide an 10-digit time-based one time password conforming to RFC6238 TOTP.
 
You have to read RFC6238 (and the errata too!) and get a correct one time password by yourself.
TOTP's "Time Step X" is 30 seconds. "T0" is 0.
Use HMAC-SHA-512 for the hash function, instead of the default HMAC-SHA-1.
Token shared secret is the userid followed by ASCII string value "HDECHALLENGE003" (not including double quotations).
 
For example, if the userid is "ninja@example.com", the token shared secret is "ninja@example.comHDECHALLENGE003".
For example, if the userid is "ninjasamuraisumotorishogun@example.com", the token shared secret is "ninjasamuraisumotorishogun@example.comHDECHALLENGE003"

If your POST request succeeds, the server returns HTTP status code 200.

