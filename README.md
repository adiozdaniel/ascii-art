# ASCII-Graphics Tool

## Description

This program is designed to graphically represent an input string using different ASCII formats.

These formats are represented in banner files; standard.txt, shadow.txt and thinkertoy.txt.

This program can handle an input with printable ASCII charcters (numbers, letters, spaces, special characters) and `\n` (newline character).

To better showcase this, here is an example:

```Bash
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

## Implementation

To correctly graphically represent this input, we mapped the contents of the banner files and checked the input string against the map.

If a character in the input string is found in the map, this character will be printed.

**😃Fun fact: Each ASCII character in the banner files has a height of 8.**


## Installation

To clone and run this program, you'll need **Git** installed on your computer.

From the **command line**,

```Bash
# clone this repository
$ git clone https://learn.zone01kisumu.ke/git/josopondo/ascii-art-color

# go into the repository
$ cd ascii-art

#open a code editor like VS Code
$ code .
```

## Usage

The program supports three interfaces:

### 1. Command Line Interface (CLI)

Once the program has been installed and opened, on the terminal, run the program using an input string of choice, like this:

`$ go run . "It's a wonderful day!"`

With only one argument (input_text) the program is designed to select the standard.txt banner 
file as the default. Hence 

the graphical representation will be as per the format in standard.txt.

If you want to use a different format, introduce a third argument: ***A Flag***

The flags for the banner files are ***"-standard", "standard"*** for standard.txt, ***"-shadow", "shadow"*** for shadow.txt, and ***"-thinkertoy", "thinkertoy"*** for thinkertoy.txt. 

The flags will prompt the program to select the appropriate file and display the output in the correct format.

For example:

`$ go run . "Hello\nThere" "-thinkertoy"  //"-thinkertoy" as flag for thinkertoy.txt`

will have the following output:

```Bash
# output as per thinkertoy.txt format

o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                 
                 
                       
o-O-o o                
  |   |                
  |   O--o o-o o-o o-o 
  |   |  | |-' |   |-' 
  o   o  o o-o o   o-o 

```

Try with more examples and watch the magic happen!!😃


#### Handling Non-ASCII Characters

In the case of special non-ASCII characters like emojis, 
the program is designed to 

print the valid ASCII characters and let you know which invalid (non-ascii) characters were skipped.

For instance:

```Bash
$ go run . "Google😋🤯🫣"

# output

  _____                           _         
 / ____|                         | |        
| |  __    ___     ___     __ _  | |   ___  
| | |_ |  / _ \   / _ \   / _` | | |  / _ \ 
| |__| | | (_) | | (_) | | (_| | | | |  __/ 
 \_____|  \___/   \___/   \__, | |_|  \___| 
                           __/ |            
                          |___/             
😋 🤯 🫣 These characters were skipped!
```

**Note:**
Special characters can only appear once.

### 2. File Interface

The program transfers the ascii output to a given file as in the example below:

`go run . --output=sample.txt "Hello World!"`

### 3. Web Interface

The program displays a graphical Web Interface.

To start the web server:
- Navigate to /cmd: `cd cmd`
- Start the server by entering: `go run . -web`
The server will start as long as the first arguement is '-web' flag.

## Features

### 1. Color

The output can be displayed in different colors in any of the following formats:

#### Text-format

1. By adding a color flag and a refference string, i.e letters to be colored in the input string, like this:

    - `$ go run . --color=mint ho hello`

On the terminal, you should be able to see letters ***h*** and ***o*** in mint and the remaining letters in white.

2. By adding only the color flag, like this:

    - `$ go run . --color=mint hello`

 In this case, all the letters in ***hello*** will be colored in mint. Here, the string `hello` acts as the reference string.

#### RGB-format

 1. By adding an RGB color code, like this:

    - `$ go run . "--color=rgb(100, 150, 180)" hello`

 **Note:** The RGB color format requires that the color flag and it's RGB value be enclosed in quotation marks; as shown above. This is because brackets have a syntactical interpretation in bash. 

#### Hex-format

 1. By adding hexadecimal color codes, like this:

    - `$ go run . --color=#e3ee38 hello hello`

 2. This program also supports shorthand hexadecimal color codes:
    
    - `$ go run . --color=#ff0 hello hello`
    - `$ go run . --color=#333 hello hello`

#### HSL-format

 1. By adding HSL color codes, like this:
    - `$ go run . "--color=hsl(176, 95%, 50%)" hello hello`
    - `$ go run . "--color=HSL(176, 95%, 50%)" hello hello`

**Note:** The HSL color format requires that the color flag and it's HSL value be enclosed in quotation marks; as shown above. This is because brackets have a syntactical interpretation in bash.


 The ***Text-color-format*** supports a limited number of colors (22); while ***RGB-***, ***Hex-*** and ***HSL-color-formats*** have an unlimited number of colors.

 ***The color flag has to be written as one argument; without spaces or with double quotes enclosing the flag. This way, the program will give an accurate output.***

 Get more color combinations [here](https://htmlcolorcodes.com/) 

## Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/adiozdaniel>
            <img src=https://avatars.githubusercontent.com/u/42722945?v=4 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Adioz/>
            <br />
            <sub style="font-size:14px"><b>Adioz Eshitemi</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/andyosyndoh>
            <img src=https://lh3.googleusercontent.com/a/ACg8ocLUKAW3QwBqLDqDcmkFTC3wmCPq0dd25wVFn3CPEkCfhQQme9Lx=s288-c-no width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Andrew/>
            <br />
            <sub style="font-size:14px"><b>Andrew Osindo</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/josie-opondo>
            <img src=https://avatars.githubusercontent.com/u/77047643?v=4 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Josephine/>
            <br />
            <sub style="font-size:14px"><b>Josephine Opondo</b></sub>
        </a>
    </td>
</tr>
</table>
