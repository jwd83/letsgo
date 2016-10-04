package main

/*

Go Language Examples. These files are intended to be read in the order of their
filenames.

   ooooo  oooo     oooo  ooooooooo
    888    88   88  88    888    88o
    888     88 888 88     888    888
    888      888 888      888    888
    888       8   8      o888ooo88
 8o888

Copyright (C) Jared De Blander
 _       _                 _            _   _
(_)     | |               | |          | | (_)
 _ _ __ | |_ _ __ ___   __| |_   _  ___| |_ _  ___  _ __
| | '_ \| __| '__/ _ \ / _` | | | |/ __| __| |/ _ \| '_ \
| | | | | |_| | | (_) | (_| | |_| | (__| |_| | (_) | | | |
|_|_| |_|\__|_|  \___/ \__,_|\__,_|\___|\__|_|\___/|_| |_|


Picking up from where we left off going to be doing more with variables, scope
and functions

*/

// let's start off by grabbing our favorite output lib
import "fmt"

/*
 ___  ___ ___  _ __   ___
/ __|/ __/ _ \| '_ \ / _ \
\__ \ (_| (_) | |_) |  __/
|___/\___\___/| .__/ \___|
              | |
              |_|

Let's create a variable that's scope is outside of the main function. To do so
we can create one here like so...
*/

var hendrix string = "Hey Joe," // VALID

/*
However while we cannot do the following shorthand...

hendrix := "Hey, Joe"       // INVALID!!!

...we can use

var hendrix = "Hey Joe"     // ALSO VALID!!!
*/

const the_next_line = "Hey Joe, I said where you goin' with that gun in your hand, oh"

func main() {
    /*
                     _ _   _                 _       _     _
                    | | | (_)               (_)     | |   | |        ___
     _ __ ___  _   _| | |_ ___   ____ _ _ __ _  __ _| |__ | | ___   ( _ )
    | '_ ` _ \| | | | | __| \ \ / / _` | '__| |/ _` | '_ \| |/ _ \  / _ \/\
    | | | | | | |_| | | |_| |\ V / (_| | |  | | (_| | |_) | |  __/ | (_>  <
    |_| |_| |_|\__,_|_|\__|_| \_/ \__,_|_|  |_|\__,_|_.__/|_|\___|  \___/\/


                     _ _   _                     _              _
                    | | | (_)                   | |            | |
     _ __ ___  _   _| | |_ _  ___ ___  _ __  ___| |_ __ _ _ __ | |_
    | '_ ` _ \| | | | | __| |/ __/ _ \| '_ \/ __| __/ _` | '_ \| __|
    | | | | | | |_| | | |_| | (_| (_) | | | \__ \ || (_| | | | | |_
    |_| |_| |_|\__,_|_|\__|_|\___\___/|_| |_|___/\__\__,_|_| |_|\__|


         _       __ _       _ _   _
        | |     / _(_)     (_) | (_)
      __| | ___| |_ _ _ __  _| |_ _  ___  _ __  ___
     / _` |/ _ \  _| | '_ \| | __| |/ _ \| '_ \/ __|
    | (_| |  __/ | | | | | | | |_| | (_) | | | \__ \
     \__,_|\___|_| |_|_| |_|_|\__|_|\___/|_| |_|___/

    we can also define blocks of variables quickly using the following styles...
    */
    var (
        a = 4
        b = 5
        c = 9
        user_input float64
    )

    // as well as constants
    const (
        the_line_after_that = "I'm goin' down to shoot my old lady"
        pi = 3.14159265

    )

    // calling jimisays which will print Hey Joe and update the value of the
    // hendrix variable
    jimisays()              // displays: Hey Joe,
    fmt.Println(hendrix)    // displays: where you goin' with that gun of yours?

    // if you uncomment the next line the complie will fail as we cannot reassign
    // a constant
    //the_next_line = "THIS WILL FAIL TO COMPILE AS the_next_line IS A const"

    // lets print our out of scope constant containing the next line
    fmt.Println(the_next_line)
    fmt.Println(the_line_after_that)

    c += a + b
    fmt.Println("after c += a + b, c =", c)

    /*
     _               _                              _                   _
    | |             (_)                            (_)                 | |
    | |__   __ _ ___ _  ___   _   _ ___  ___ _ __   _ _ __  _ __  _   _| |_
    | '_ \ / _` / __| |/ __| | | | / __|/ _ \ '__| | | '_ \| '_ \| | | | __|
    | |_) | (_| \__ \ | (__  | |_| \__ \  __/ |    | | | | | |_) | |_| | |_
    |_.__/ \__,_|___/_|\___|  \__,_|___/\___|_|    |_|_| |_| .__/ \__,_|\__|
                                                           | |
                                                           |_|
    let's get a number from the user and multiply it by the pi constant we
    defined in the multivarabile section
    */

    fmt.Print("Enter a number: ")
    fmt.Scanf("%f", &user_input)
    user_output := user_input * pi
    fmt.Println("Number entered multiplied by pi =", user_output)

}

func jimisays() {

    fmt.Println(hendrix)
    hendrix = "where you goin' with that gun of yours?"
}
