# Daily Coding Problem: Problem #385 [Medium]

This problem was asked by Apple.

You are given a hexadecimal-encoded string that has been XOR'd against a single char.

Decrypt the message. For example, given the string:

7a575e5e5d12455d405e561254405d5f1276535b5e4b12715d565b5c551262405d505e575f

You should be able to decrypt it and get:

```
Hello world from Daily Coding Problem
```

## Build and Run

There's a decoder:

```sh
$ go build decode.go
$ ./decode -f example.in
```

The decoder also has a flag to allow more than 0 non-ASCII bytes 
in the possible cleartext.

I wrote an encoder to see if I understood the problem statement.

```sh
$ go build encode.go
$ echo "now is the time for all good men to come to" > input
$ ./encode 45 input > encoded.txt
$ ./decode -f encoded.txt -e 1 
...
Best key 45
now is the time for all good men to come to
$
```

## Analysis

This is a crypto question,
which is somewhat unfortunate for the candidate,
as crypto problems always always always have the weirdest bugs,
and are finicky to get right.

[Key elimination](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher#Key_elimination)
requires a known plaintext.
While the problem statement does give an entire plaintext
for the example ciphertext,
the problem statement doesn't say you've got one.
Other than that, key elimination would work well with single-byte keys.

Only 256 single-byte XOR keys exist (values 0 through 255),
so it's computationally feasible to iterate through all of them
to get potential cleartext.
It's sometimes possible to find multiple XOR keys that decode the
ciphertext to (mostly) ASCII bytes,
so it's not possible to just say "whatever key byte" decodes
every ciphertext byte to ASCII.

Luckily, I have some [classic cipher code](https://github.com/bediger4000/vigenere-ciphering-deciphering) from a while back.
It includes a way to compare possible cleartext with English
character frequencies.
By picking the single-byte key that yields cleartext
that has the closest similarity to English character frequencies,
we almost always find the single byte key.

Unfortunately, I wrote the code to throw out possible key bytes
that yield too many errors,
so I kludged in a flag that will cause it to consider
only similarity to English character frequencies.

```sh
$ ./decode -f encoded.txt -i
...
Best key 45
now is the time for all good men to come to
```

This seems to work well.

This really isn't a problem for an entry- or even junior-level job candidate.
It's too finicky to get working code,
and the obvious criteria (lowest number of non-ASCII bytes in possible cleartext)
gives mediocre results.
Unfortunately, the character frequency data isn't amenable to
carrying around in a human memory,
and the dot-product calculation isn't easy to get correct.
If the interviewer will accept some handwaving
("and here's where we'd include an English character frequence count")
it's maybe a good medium to advanced-level problem.
But even a decent programmer who hasn't done some cipher-related
programming will make dreadful mistakes.

Maybe the interviewers accepted not-quite-perfect programs
just so they could see the candidates write some code.

## Around the web

Ha ha!, This is [crypto-pals challenge 3](https://cryptopals.com/sets/1/challenges/3).
Which I had [already done](https://cryptopals.com/sets/1/challenges/3) when I did this Daily Coding Problem.
