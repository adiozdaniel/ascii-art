# ASCII-Graphics Tool

## Description

This program is designed to graphically represent an input string using different ASCII formats.

These formats are represented in the following banner files:

- standard.txt
- shadow.txt
- thinkertoy.txt.

This program can handle an input with printable ASCII characters (numbers, letters, spaces, special characters) and `\n` (newline character).

Here is an example:
example input:
`$ go run . "Hello"`

example output:

![go run . "Hello"](/static/sample1.png)

## Implementation

To correctly graphically represent this input, we mapped the contents of the banner files and checked the input string against the map.

If a character in the input string is found in the map, this character will be printed.

**😃Fun fact: Each ASCII character in the banner files has a height of 8.**

## Installation

This application requires Go (golang) 1.18 or higher to run. You can get it [here](https://go.dev/doc/install)

To clone and run this program, you'll need **Git** installed on your computer.

From the **command line**,

```Bash
# clone this repository
$ git clone https://learn.zone01kisumu.ke/git/aosindo/ascii-art-output

# go into the repository
$ cd ascii-art

# open a code editor like VS Code
$ code .

# you may be required to rebuild the Go module
# thus, delete the existing module, if any and run
$ go mod init github.com/adiozdaniel/ascii-art
```

## Usage

The program supports three interfaces:

### 1. Command Line Interface (CLI)

- Once the program has been installed, navigate to the `cmd` directory.

- Run the program using an input string of choice, like this:

`$ go run . "A wonderful day!"`

With only one argument the program is designed to select the 'standard.txt' banner
file as the default. Hence the graphical representation will be as per the format in 'standard.txt'.

If you want to use a different format, introduce a second argument; a flag.

The flags for the banner files are:

- **_"standard"_** or **_"-standard"_** or **_"--standard"_** for standard.txt
- **_"shadow"_** or **_"-shadow"_** or **_"--shadow"_** for shadow.txt
- **_"thinkertoy"_** or **_"-thinkertoy"_** or **_"--thinkertoy"_** for thinkertoy.txt.

The flags will prompt the program to select the appropriate file and display the output in the correct format.

For example:

- To use thinkertoy:

`$ go run . "Hello\nThere" "-thinkertoy"`

will have the following output:

![go run . "Hello\nThere" "-thinkertoy"](/static/sample2.png)

Try with more examples and watch the magic happen!!😃

#### Handling Non-ASCII Characters

In the case of special non-ASCII characters like emojis,
the program is designed to print the valid ASCII characters and let you know which invalid (non-ascii) characters were skipped.

For instance:

`$ go run . "Google😋🤯🫣"`

![go run . "Google😋🤯🫣"](/static/sample3.png)

**Note:**
These characters will only appear once in the warning output.

### 2. File Interface

The program writes the ascii output to a given file, when run like this:

`go run . --output=sample.txt "Hello World!"`

Take a good look at **_--output=sample.txt_**. Here we have to use the flag **_--output=_** and specify the **_text_** file we are writing to, in the exact same format as in this example.

We currently only support writing text files. Other formats are yet to be added... keep following.

### 3. Web Interface

The program displays a graphical Web Interface.

To start the web server:

- Navigate to /cmd: `cd cmd`
- Start the server by entering: `go run . -web`

The server will start as long as the first argument is '-web' flag.

## Features

### Color

The output can be displayed in different colors in any of the following formats:

#### Text-format

1. By adding only the color flag, like this:

   - `$ go run . --color=blue hello`
     ![go run . --color=blue hello](/static/sample5.png)

   In this case, all the letters in **_hello_** will be colored in blue. Here, the string `hello` acts as the reference string.

2. By adding a color flag and a refference string that is not a substring of the next word,like in the example below, the program will look for the instance of the characters in the input string and color them with the provided color, like this:

   - `$ go run . --color=blue ho hello`
     ![go run . --color=blue ho hello](/static/sample4.png)

     On the terminal, you should be able to see letters **_h_** and **_o_** in blue and the remaining letters in default terminal color(possibly, white).

3. By adding a color flag and a refference string that is a substring of the next word,like in the example below, the program will look for the instance of the substring in the input string and color them with the provided color, like this:

   - `$ go run . --color=#f0f "Will" "Will will come\nTo fetch the Will\nTo Will James"`
     ![go run . --color=blue ho hello](/static/sample6.png)

     On the terminal, you should be able to see the substring **_Will_** in yellow (or the provided colour) and the remaining letters in default terminal color(possibly, white).

#### RGB-format

1.  By adding an RGB color code, like this:

    - `$ go run . "--color=rgb(100, 150, 180)" hello`

**Note:** The RGB color format requires that the color flag and it's RGB value be enclosed in quotation marks; as shown above. This is because brackets have a syntactical interpretation in bash.

#### Hex-format

1.  By adding hexadecimal color codes, like this:

    - `$ go run . --color=#e3ee38 hello hello`

2.  This program also supports shorthand hexadecimal color codes:

    - `$ go run . --color=#ff0 hello hello`
    - `$ go run . --color=#333 hello hello`

#### HSL-format

1.  By adding HSL color codes, like this:
    - `$ go run . "--color=hsl(176, 95%, 50%)" hello hello`
    - `$ go run . "--color=HSL(176, 95%, 50%)" hello hello`

**Note:** The HSL color format requires that the color flag and it's HSL value be enclosed in quotation marks; as shown above. This is because brackets have a syntactical interpretation in bash.

The **_Text-color-format_** supports a limited number of colors (22); while **_RGB-_**, **_Hex-_** and **_HSL-color-formats_** have an unlimited number of colors.

**_The color flag has to be written as one argument; without spaces and with double quotes enclosing the flag. This way, the program will give an accurate output._**

Get more color combinations [here](https://htmlcolorcodes.com/)

**Note:** This feature is only available in the CLI mode.

## Disclaimer

The program currently supports three interfaces. You are thus adviced to **explicitly declare** your intended use, or else, you may encounter wrong output or wrong usage errors.

For instance:

`go run .` will throw a full fledged usage error like:

```Bash
For color:
EX: go run . --color=<color> <substring to be colored> "something" standard
For output:
EX: go run . --output=<fileName.txt> something standard
For justify:
Example: go run . --align=right something standard
For web:
go run . -web
```

But explicitly telling the program the intended use, such as `go run . --output=sample.txt something`, will prompt the program to use the file interface.

### Alignment

To change the alignment of your output dynamically, you need to include the use of `--align=<type>` flag.
This can be:

- center
- left
- right
- justify

Your representation adopts to the terminal size. If you reduce the terminal window the graphical representation adapts to the terminal size.

The input MUST follow a specified format:

            [OPTION]       [YOUR INPUT]     [BANNER]

`go run . --align=center     "My Papa"       standard`

The updated version 0.3 supports color as an option:

**Disclaimer**
Note, this is an added feature for the web that would still work for justified.

            [       OPTION          ]      [         YOUR INPUT          ]          [BANNER]

`go run . --align=center --color=#ff0       "Papa" "My Papa\nIs Papa?\nNo Papa"    thinkertoy`

**Note:** Only text that fits the terminal size should be tested.

To quit the program, type `exit` in lowercase only. This is case sensitive

## Enjoy!

## Authors

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
