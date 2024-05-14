# ASCII-Graphics Tool

This program is designed to graphically represent an input string using different ASCII formats.

These formats are represented in banner files; standard.txt, shadow.txt and thinkertoy.txt.

To better showcase this, here is an example:

```
# example input
$ go run . "Hello"

# example output using "standard.txt" format
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/

```

To correctly graphically represent this input, we mapped the contents of the banner files and checked the input string against the map.

If a character in the input string is found in the map, this character will be printed.

**😃Fun fact: Each ASCII character in the banner files has a height of 8.**

This program can handle an input with printable ASCII charcters (numbers, letters, spaces, special characters) and `\n` (newline character).

## Installation

To clone and run this program, you'll need **Git** installed on your computer.

From the **command line**,

```
# clone this repository
$ git clone https://learn.zone01kisumu.ke/git/adaniel/ascii-art.git

# go into the repository
$ cd ascii-art

#open a code editor like VS Code
$ code .
```

## How To Use

Once the program has been installed and opened, on the terminal, run the program using an input string of choice, like this:

```
$ go run . "It's a wonderful day!!"
```

and watch the magic happen!!😃

### Special Characters Handling

You can use the special characters like emojis without spaces between them.
You can also use them to form a special pattern of your choice.

for instance:

`go run . "Google😋🤯🫣"`

**Note**
Special characters can only appear once.

## Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/adaniel>
            <img src=https://learn.zone01kisumu.ke/git/avatars/4f09df65142ebdeb582566255916d89c?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Aaron/>
            <br />
            <sub style="font-size:14px"><b>Adioz Eshitemi</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aosindo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/fa966ef34b0ccdfe772414745aeee49f?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Emmanuel/>
            <br />
            <sub style="font-size:14px"><b>Andrew Osindo</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/josopondo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/150a3fdbc5cc89dd642dd783c474b61c?=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Abraham/>
            <br />
            <sub style="font-size:14px"><b>Josephine Opondo</b></sub>
        </a>
    </td>
</tr>
</table>

### Todos

- unit testing
- map non-ascii
- if the file is corrupted
- if the file is missing
- data security
- tabs
- backspace
- flags
- default settings
