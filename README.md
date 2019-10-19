# Translate a name written in English to Klingon

Translate a name written in English to Klingon and find out its species

## Test task requirements

    1. It should be a public Git repository with all commit history included;
    2. It should be runnable on a Unix bash. The name in English will be passed as the
       first parameter. It might contain spaces in case of composed names.
    3. Consider each Klingon cognate letter avalid correspondence to an English letter. For example, ​D is a valid correspondence of ​d on so on.
        You might notice that some letters are missing which means they are not translatable for this test purposes, then ignore the whole input.
    4. The output should contain:
        a. The translated name in Klingon written using the correspondent
        hexadecimal numbers according to the given table. Format:
            i. Each hexadecimal number should be separated from each other
               using a single space;
            ii. If the translated name has spaces, use ​0x0020 for representing
                each space character;
        b. The species of the given Star Trek character name using the API;
        c. The translated name and the species name separated by a new line;

## Cli application requirements

    Go >= 1.11.4
    GNU Make >= 3.81 (optionnal)

## Build cli executable file

```bash
$ git clone git@github.com:RadikChernyshov/klingon.git
$ cd klingon
$ make
```
or

```bash
$ git clone git@github.com:RadikChernyshov/klingon.git
$ cd klingon
$ go get -d
$ go build -o klingon cmd/main.go
```


## Cli application usage

```bash
./klingon --help
NAME:
   Translate a name written in English to Klingon - cli application

USAGE:
   klingon name

VERSION:
   1.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Cli application output examples

Valid character name and character species found.

```bash
./klingon Joe Tormolen
0xF8D8 0xF8DD 0xF8D4 0x0020 0xF8E3 0xF8DD 0xF8E1 0xF8DA 0xF8DD 0xF8D9 0xF8D4 0xF8DB
Human
```

Valid character name but character species not found.

```bash
./klingon Eddie Newsom
0xF8D4 0xF8D3 0xF8D3 0xF8D7 0xF8D4 0x0020 0xF8DB 0xF8D4 0xF8E7 0xF8E2 0xF8DD 0xF8DA
character `Eddie Newsom` species not found
```

Invalid character name. (not translatable for this test purposes)

```bash
./klingon Radik
K character is not translatable
```

