# Bitxpp

A cli tool for checking decimal, hexadecimal, binary notation.

## Install

```
go get github.com/Bo0km4n/bitxpp 
```

## Usage

### Example: Type Hex and check decimal, bit notation.

```
bitxpp -i a00f -m x
```

output is

```
$ bitxpp -i a00f -m x
(decimal):            40975
(hex)    :             a00f
(bit)    : 1010000000001111
```

### help

```
$ bitxpp -h
  -c	console mode
    	commands:
    		'/d => change decimal mode'
    		'/x => change hex mode'
    		'/b => change bit mode'
  -i string
    	input text
  -m string
    	input mode
    	'd: deciaml'
    	'x: hex'
    	'b: bit' (default "d")
```

## Mode

bitxpp has main 3 input modes.

1. `-m d` input decimal
2. `-m x` input hex
3. `-m b` input bit

### console mode
```
$ bitxpp -c
(decimal) >> 100
(decimal):     100
(hex)    :      64
(bit)    : 1100100
(decimal) >> /d
(decimal) >> /x
(hex) >> ffff
(decimal):            65535
(hex)    :             ffff
(bit)    : 1111111111111111
(hex) >> /cb
(bit) >> 01001010
(decimal):      74
(hex)    :      4a
(bit)    : 1001010
(bit) >>
```

type command `/b`, `/x`, `/d` changes mode.

## Option

- `-h` help
- `-i` input. please type number you want to check.
- `-m` mode


