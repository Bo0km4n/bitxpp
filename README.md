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
$ bitxpp -i a00f -m h
(decimal):            40975
(hex)    :             a00f
(bit)    : 1010000000001111
```

### help

```
$ bitxpp -h
  -c	console mode
    	commands:
    		'/cd => change decimal mode'
    		'/ch => change hex mode'
    		'/cb => change bit mode'
  -i string
    	input text
  -m string
    	input mode
    	'd: deciaml'
    	'h: hex'
    	'b: bit' (default "d")
```

## Mode

bitxpp has 3 input modes.

1. `-m d` input decimal
2. `-m x` input hex
3. `-m b` input bit

## Option

- `-h` help
- `-i` input. please type number you want to check.
- `-m` mode


