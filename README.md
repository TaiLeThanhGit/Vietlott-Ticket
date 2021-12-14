# Vietlott-Ticket - Mega 6/45 - Power 6/55
## _This Golang source will generates six random numbers_

Their values are from 1 to 45 or from 1 to 55.
And their values are not in a list of numbers that saved in a file.
This file has many line, each line has six numbers that their value are from 1 to 45 or 1 to 55.

## Features

- Generate MEGA Ticket
- Generate POWER Ticket
- Import file that contains the list of unexpected numbers

## Build

Build the source on windows

```sh
 go build -o vietlott.exe vietlott.go
```
## How to run
Generate MEGA ticket

```sh
 vietlott.exe -t 1
```
>Output
```sh
Unexpected Numbers: []
Type:  MEGA
Number of Tickets:  1
1: 9 12 15 21 37 38

```
Generate five MEGA tickets

```sh
 vietlott.exe -t 1 -n 5
```
> Output
```sh
Unexpected Numbers: []
Type:  MEGA
Number of Tickets:  5
1: 5 11 18 28 31 36
2: 9 18 24 25 31 34
3: 9 16 18 20 21 37
4: 2 12 21 22 34 37
5: 1 13 25 27 30 41
```
Generate five MEGA tickets and the tickets numbers are not in a file
```sh
vietlott.exe -t 1 -n 5 -f unexpected_numbers.txt
```
> Output
```sh
Unexpected Numbers: [3 5 6 31 34 35 10 14 29 37 38 40 8 9 30 45 13 17 26 27]
Type:  MEGA
Number of Tickets:  5
1: 10 13 15 18 37 42
2: 2 9 12 30 32 41
3: 9 16 28 29 35 36
4: 3 6 8 11 23 28
5: 10 13 29 36 40 44
```
## For more details
Run 
```sh
vietlott.exe -h
```
> Output
```sh
vietlott.exe  [-t type] [-n quantity] [-f file] [-h]
         type: 1 is Mega, 2 is Power
         quantity: > 0
         file: file name contains the list of unexpected numbers
         -h: show this help

```

## Author

**Tai Le Thanh**
