# ASCII-Graphics Tool Version 0.4

## Description

ASCII Art Graphics Tool was crafted to turn plain old text into amazing, eye-catching ASCII art. Whether you’re into classic characters or creative patterns, this tool has you covered with various formats and fonts like standard, shadow, and thinkertoy.

These fonts and/or formats are represented in the following banner files:

- standard.txt
- shadow.txt
- thinkertoy.txt.

This program can handle an input with printable ASCII characters (numbers, letters, spaces, special characters) and `\n` (newline character).

Here is an example:
example input:
`go run . "Hello"`

example output:

![go run . "Hello"](/views/static/images/sample1.png)

## Implementation

To correctly graphically represent this input, we mapped the contents of the banner files and checked the input string against the map.

If a character in the input string is found in the map, this character will be printed.

**😃Fun fact: Each ASCII character in the banner files has a height of 8.**

## Installation

This application requires Go (golang) 1.18 or higher to run. You can get it [here](https://go.dev/doc/install)

To clone and run this program, you'll need **Git** installed on your computer.

From the **command line**,

```Bash
git clone https://learn.zone01kisumu.ke/git/adaniel/ascii-art-web.git
cd ascii-art-web
go mod init github.com/adiozdaniel/ascii-art
code .
```

## Usage

The program supports three interfaces:

### 1. Command Line Interface (CLI)

- Once the program has been installed, **navigate to the `cmd` directory**.

- Run the program using an input string of choice, like this:

`go run . "A wonderful day!"`

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

`go run . "Hello\nThere" "-thinkertoy"`

will have the following output:

![go run . "Hello\nThere" "-thinkertoy"](/views/static/images/sample2.png)

**Note:** Remember to use `""` for enclosing grouped input.

Try with more examples and watch the magic happen!!😃

#### Handling Non-ASCII Characters

In the case of special non-ASCII characters like emojis,
the program is designed to print the valid ASCII characters and let you know which invalid (non-ascii) characters were skipped.

For instance:

`go run . "Google😋🤯🫣"`

![go run . "Google😋🤯🫣"](/views/static/images/sample3.png)

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

1. By adding only the color flag and a valid color, like this:

   - `go run . --color=blue hello`

     ![go run . --color=blue hello](/views/static/images/sample5.png)

   In this case, all the letters in **_hello_** will be colored in blue. Here, the string `hello` acts as the reference string.

2. By adding a color flag and a refference string that is not a substring of the string to be colored, like in the example below, the program will look for the instance of the characters in the string to be colored and color them with the provided color, like this:

   - `go run . --color=blue ho hello`

     ![go run . --color=blue ho hello](/views/static/images/sample4.png)

     On the terminal, you should be able to see letters **_h_** and **_o_** in blue and the remaining letters in default terminal color(possibly, white).

3. By adding a color flag and a refference string that is a substring of the string to be colored,like in the example below, the program will look for the instance of the substring in the string to be colored and color them with the provided color, like this:

   - `go run . --color=#f0f "Will" "Will will come\nTo fetch the Will\nTo Will James"`

     ![go run . --color=blue ho hello](/views/static/images/sample6.png)

     On the terminal, you should be able to see the substring **_Will_** in yellow (or the provided colour) and the remaining letters in default terminal color(possibly, white).

#### RGB-format

1.  By adding an RGB color code, like this:

    - `go run . "--color=rgb(100, 150, 180)" hello`

#### Hex-format

1.  By adding hexadecimal color codes, like this:

    - `go run . --color=#e3ee38 hello hello`

2.  This program also supports shorthand hexadecimal color codes:

    - `go run . --color=#ff0 hello hello`
    - `go run . --color=#333 hello hello`

#### HSL-format

1.  By adding HSL color codes, like this:
    - `go run . "--color=hsl(176, 95%, 50%)" hello hello`
    - `go run . "--color=HSL(176, 95%, 50%)" hello hello`

**Note:** The HSL and RGB color formats require that the color flag and it's value be enclosed in quotation marks; as shown above. This is because brackets have a syntactical interpretation in bash.

The **_Text-color-format_** supports a limited number of colors (22); while **_RGB-_**, **_Hex-_** and **_HSL-color-formats_** have an unlimited number of colors.

Get more color combinations [here](https://htmlcolorcodes.com/)

**Note:** This feature is only available in the CLI mode.

### Alignment

To change the alignment of your output dynamically, you need to include the use of `--align=<type>` flag.
This can be:

- center
- left
- right
- justify

In alignment mode, the representation will adapt to the terminal size. If you were to reduce the terminal window, the graphical representation will adapt itself to fit into the new terminal size.

The input follows a specified format:

            [OPTION]       [YOUR INPUT]     [BANNER]

`go run . --align=center     "My Papa"       standard`

#### Current Version Updates (v 0.4)

- Supports color as an option and runs with or without [BANNER](#) specified
- Supports changing alignment options, input and color during runtime
- For the best experience, use the **full terminal size**

`go run . --color=#ff0 --align=center "Will" "Will will come\nHe has Will\n& Will"`

**Note:**

- This is an added feature for open source purposes.

            [       OPTION          ]      [         YOUR INPUT          ]          [BANNER]

`go run . --align=center --color=#ff0       "Papa" "My Papa\nIs Papa?\nNo Papa"    thinkertoy`

    or

`go run . --color=#ff0 --align=center       "Will" "Will will come\nHe has Will\n& Will"`

**_Output that surpasses the screen size will be hidden._**

- Try out this latest feature of manipulating output during runtime:
  - you can fire up the cli mode by using `go run . --align=center --`
  - color by entering a different color : "--color=<[your color choice](https://htmlcolorcodes.com/)>". You need to follow color rules to get expected output.
  - to specify color reference (the characters or substring to be coloured); use `--reff=<your ref>` for instance, `--reff=come`.
  - to reset the reff, use: `--reff=""`
    **Note** this _**limited**_ feature is only supported during runtime.
  - align by entering a different alignment choice `--align=<justify>` or `left` or `center` or `right`. other choices will be ignored.
  - lack of specifying flag option for instance `--align` will make the program to use it as input.
  - the program only displays characters that fit the screen size, this turns taking amazing screenshots into a beautiful experience.
  - to change font styling (what was previously done as `--banner=<file>`), you simply type your styling in the flag format:
    - `--standard` for standard font styling
    - `--thinkertoy` for thinkertoy font styling
    - `--shadow` for shadow font styling

To quit the program, type `exit` in lowercase only. 

## Disclaimer

The program currently supports three interfaces. You are thus adviced to **explicitly declare** your intended use, or else, you may encounter wrong output or usage errors.

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
