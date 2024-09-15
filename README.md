# ccwc

## Coding Challenge #1 - WC tool

This repository holds the source code completing the first challenge from CodingChallenges.fyi.

This tool is called the wc tool commonly available on linux/unix based terminals to count bytes/characters/words/lines in a text file.

## Usage

Run the following command to create an executable go file:

`make build`

Run the tool using:

`./ccwc [-c|-m|-w|-l] <filename>`

OR

`cat <filename> | ./ccwc [-c|-m|-w|-l]`

## Definitions of flag(s)

    -c : this counts the number of bytes
    -m : this counts the number of characters
    -w : this counts the number of words
    -l : this counts the number of lines
    if no flag is provided : you will get the counts in this order - <lines> <word> <bytes>
