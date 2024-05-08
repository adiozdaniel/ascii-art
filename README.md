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

This program can handle an input with printable ASCII charcters (numbers, letters, spaces, special characters) and ``` \n ``` (newline character).

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
With only two arguments (program_name and input_text) the program is designed to select the standard.txt banner file as the default. Hence 

the graphical representation will be as per the format in standard.txt.

If you want to use a different format, introduce a third argument: ***A Flag***.

The flags for this program are only ***s*** and ***shadow*** for shadow.txt, and ***t*** and ***thinkertoy*** for thinkertoy.txt. The flags

will prompt the program to select the appropriate file and display the output in the correct format.

For example:

```
$ go run . "Hello\nThere" "t"
```

will have the following output:

```










```

Try with more examples and watch the magic happen!!😃

## Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/adaniel>
            <img src=https://learn.zone01kisumu.ke/git/avatars/4f09df65142ebdeb582566255916d89c?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Adioz/>
            <br />
            <sub style="font-size:14px"><b>Adioz Eshitemi</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/aosindo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/248ebb81529cc2dcee5e20f60e2e4d24?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Andrew/>
            <br />
            <sub style="font-size:14px"><b>Andrew Osindo</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://learn.zone01kisumu.ke/git/josopondo>
            <img src=https://learn.zone01kisumu.ke/git/avatars/150a3fdbc5cc89dd642dd783c474b61c?=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Josephine/>
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