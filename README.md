# go-cowsay-clone

A Go-based clone of the classic ['cowsay' command-line tool](https://en.wikipedia.org/wiki/Cowsay#:~:text=cowsay%20is%20a%20program%20that,It%20is%20written%20in%20Perl.)

This was built based of this [tutorial](https://flaviocopes.com/go-tutorial-cowsay/)!

This progeram takes input piped in from another command and displays it on your terminal in an ASCII speech balloon, with an animal of your choice

## Features
- Written from scratch in Go
- Speech balloon with proper Unicode character alignment
- Various ASCII figures to choose from (currently only supports cow and stegosaurus)
- Tab normalisation for cleaner output
- Support piped input

## Installation instructions

### Prerequisites
- Install [Go](https://go.dev/doc/install)

### Clone and Build
```bash
git clone git@github.com:kailermai/go-cowsay-clone.git
cd go-cowsay-clone
go build
```

## Usage
Instructions below are for windows. If you are on Linux or Mac, replace `./go-cowsay-clone.exe` with `./go-cowsay-clone`

### Basic usage
`"Hello!" | .\go-cowsay-clone.exe `

### Use a specific figure
`"Rawr!" | .\go-cowsay-clone.exe -f stegosaurus`

### Help
`.\go-cowsay-clone.exe -h`

## Example Outputs
```bash
"Rawr!" | .\go-cowsay-clone.exe -f stegosaurus

 _______
< Rawr! >
 _______
         \                      .       .
          \                    / `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    `._.'               |         }  \/~~~/
       `----.          /       }     |        /    \__/
             `-.      |       /      |       /      `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      `..-'
                      |     /        |     /     ~-.     `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>



```

```bash
"Rawr!" | .\go-cowsay-clone.exe -h

Usage of D:\Projects\go-cowsay-clone\go-cowsay-clone.exe:
  -f cow
        the animal name. Valid names are cow and `stegosaurus` (default "cow")
```