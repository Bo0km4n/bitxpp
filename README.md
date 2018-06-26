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


## Mode

bitxpp has 3 input modes.

1. `-m d` input decimal
2. `-m x` input hex
3. `-m b` input bit

## Option

- `-h` help
- `-i` input. please type number you want to check.
- `-m` mode


