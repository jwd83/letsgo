package main
s/*

Go Language Examples. These files are intended to be read in the order of their
filenames.

   ooooo  oooo     oooo  ooooooooo
    888    88   88  88    888    88o
    888     88 888 88     888    888
    888      888 888      888    888
    888       8   8      o888ooo88
 8o888

Copyright (C) Jared De Blander

Useful resources
-----------------------------------------------------------------------
The Go Programming Language Specification : https://golang.org/ref/spec
ASCII Table : http://www.asciitable.com/
ASCII Art : http://patorjk.com/software/taag/#p=display&f=Doom&t=sample

*/

// Let's Begin!

// this is a comment

/*
this is a multiline
comment
*/

/*
now let's import a package library. fmt is a package for text manipulation with
functions similar to printf and scanf. Try running "godoc fmt" on your console.
*/
import "fmt"

// create our main function that will be run when the script is launched
func main() {
    // Simple Hello, World example
    fmt.Println("Hello, Word")

    /*
                       _
                      | |
 _ __  _   _ _ __ ___ | |__   ___ _ __ ___
| '_ \| | | | '_ ` _ \| '_ \ / _ \ '__/ __|
| | | | |_| | | | | | | |_) |  __/ |  \__ \
|_| |_|\__,_|_| |_| |_|_.__/ \___|_|  |___/

===============================================================================|
type | range (| alias OR * notes)                                              |
===============================================================================|
* * * Integer Types                                                            |
===============================================================================|
int8  | −128 to 127                                             |              |
int16 | −32,768 to 32,767                                       |              |
int32 | −2,147,483,648 to 2,147,483,647                         | rune         |
int64 | −9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |              |
-------------------------------------------------------------------------------|
uint8  | 0 to 255                                               | byte         |
uint16 | 0 to 65,535                                            |              |
uint32 | 0 to 4,294,967,295                                     |              |
uint64 | 0 to 18,446,744,073,709,551,615                        |              |
-------------------------------------------------------------------------------|
uint    | architecture specific * either 32 or 64 bits                         |
int     | architecture specific * same size as uint. most commonly used integer|
uintptr | architecture specific * an unsigned integer large enough to store the|
                                  uninterpreted bits of a pointer value        |
===============================================================================|
* * * Floating Point Types (these are inexact, float64 increases precision)    |
===============================================================================|
float32 | the set of all IEEE-754 32-bit floating-point numbers                |
float64 | the set of all IEEE-754 64-bit floating-point numbers                |
-------------------------------------------------------------------------------|
complex64 | the set of complex numbers with float32 real and imaginary parts   |
complex128 | the set of complex numbers with float64 real and imaginary parts  |
===============================================================================|

    */
    // doing some interger addition
    fmt.Println("1 + 1 = ", 1 + 1)

    // floating point addition by including .0 to a number
    fmt.Println("1.0 + 1.1 = ", 1.0 + 1.1)

    /*
     _        _
    | |      (_)
 ___| |_ _ __ _ _ __   __ _ ___
/ __| __| '__| | '_ \ / _` / __|
\__ \ |_| |  | | | | | (_| \__ \
|___/\__|_|  |_|_| |_|\__, |___/
                       __/ |
                      |___/

    A string type represents the set of string values. A string value is a
(possibly empty) sequence of bytes. Strings are immutable: once created, it is
impossible to change the contents of a string. The predeclared string type is
string.
    The length of a string s (its size in bytes) can be discovered using the
built-in function len. The length is a compile-time constant if the string is a
constant. A string's bytes can be accessed by integer indices 0 through
len(s)-1. It is illegal to take the address of such an element; if s[i] is the
i'th byte of a string, &s[i] is invalid.

        using strings. strings can use either `` or ""
        using " allows you to escape characters: "\n" = newline, "\t" = tab
        using ` allows you to create multiline strings: `
            still
            inside
            this
            string
        `
    */

    var long_string string = `Hello
W
o
r
l
d
    `
    var hey string = "Hey\t\tJoe"

    fmt.Println(long_string)
    fmt.Println(hey)
    hey = "where you goin' with that gun of yours?"
    fmt.Println(hey)

    fmt.Println("\n\n   OFFSET        012345678")
    fmt.Println("   RULER         123456789")
    fmt.Println(`Length (len) of "Hey Joe" = `, len("Hey Joe"))
    fmt.Println(`"Hey Joe"[1] = `, "Hey Joe"[1], " // per http://www.asciitable.com/ Decimal 101 = e")
    fmt.Println(`"Hey " + "Joe" = `, "Hey " + "Joe")

    /*
 _                 _
| |               | |
| |__   ___   ___ | | ___  __ _ _ __  ___
| '_ \ / _ \ / _ \| |/ _ \/ _` | '_ \/ __|
| |_) | (_) | (_) | |  __/ (_| | | | \__ \
|_.__/ \___/ \___/|_|\___|\__,_|_| |_|___/

    A boolean type represents the set of Boolean truth values denoted by the
predeclared constants true and false. The predeclared boolean type is bool.
    */

    var im_true bool = true;
    var im_false bool = false;

    fmt.Println("\n\nboolean logic examples")
    fmt.Println("------------------------------")
    fmt.Println("true && true =", true && true)
    fmt.Println("true && false =", true && false)
    fmt.Println("true || true =", true || true)
    fmt.Println("true || false =", true || false)
    fmt.Println("!true =", !true)
    fmt.Println(`

        var im_true bool = true;
        var im_false bool = false;

    `)
    fmt.Println("im_true && im_false = ", im_true && im_false)

    /*
     _                _   _                     _
    | |              | | | |                   | |
 ___| |__   ___  _ __| |_| |__   __ _ _ __   __| |
/ __| '_ \ / _ \| '__| __| '_ \ / _` | '_ \ / _` |
\__ \ | | | (_) | |  | |_| | | | (_| | | | | (_| |
|___/_| |_|\___/|_|   \__|_| |_|\__,_|_| |_|\__,_|

    There are shorter ways to declare variables in go. The following
declarations are equivalent:

    */

    var normal_one string = "I'm the same"  // normal variable declaration
    var shorter_one = "I'm the same"        // compiler infers the string type
    shortest_one := "I'm the same"          // using : replaces var in notifying
                                            // the complier of a declaration

    fmt.Println(`
        var normal_one string = "I'm the same"
        var shorter_one = "I'm the same"
        shortest_one := "I'm the same"
    `)

    fmt.Println("normal_one == shorter_one =", normal_one == shorter_one)
    fmt.Println("shorter_one == shortest_one =", shorter_one == shortest_one)
    fmt.Println("normal_one == shortest_one =", normal_one == shortest_one)

}
