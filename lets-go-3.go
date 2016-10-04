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


Expanding our knowledge into control structures and loops
*/

import "fmt"

// Going to use the time pacakge breifly this... time!
import "time"

func main() {

    /*
     _   _            _ _  __          _ _   _
    | | | |          ( | )/ _|        ( | ) | |
    | |_| |__   ___   V V| |_ ___  _ __V V  | | ___   ___  _ __
    | __| '_ \ / _ \     |  _/ _ \| '__|    | |/ _ \ / _ \| '_ \
    | |_| | | |  __/     | || (_) | |       | | (_) | (_) | |_) |
     \__|_| |_|\___|     |_| \___/|_|       |_|\___/ \___/| .__/
                                                          | |
                                                          |_|
    */

    x := 1 // let's begin by creating an integer x and assign it a value of 1

    for x <= 10 {                       // let's run this loop while x is
                                        // less than or equal to (<=) 10
        fmt.Println("x is now:", x)     // display the current value of i
        x += 1  // let's simply increment x each time through the loop
    }

    fmt.Println("After the loop x is now equal to ", x)

    /*
    Many other programming languages have various types of loops such as while,
    do, until, foreach, etc; however, Go simplifies this to just the for loop.
    Another way to write the above....
    */

    for y := 1; y <= 10; y++ {  // this includes the declaration (y : = 1),
                                // test y <= 10 and incrementer y++ that runs
                                // each time the loop completes all in one line.
                                // Also of note When declared in this manner y
                                // only has scope inside the loop and is
                                // inaccessible once it completes.
        fmt.Println("y is now: ", y)
    }

    // Uncommenting the line below will cause this script to fail to compile...
    // fmt.Println("After the loop y is now equal to ", y)

    // Want to create an infinite loop? Simple use for by itself. Uncomment
    // the loop below to see a simple example

    /*
    for {
        fmt.Println("Please ctrl-break me! ")
    }
    */

    /*
     _   _            _ _ _  __ _ _       _        _                            _
    | | | |          ( | |_)/ _( | )     | |      | |                          | |
    | |_| |__   ___   V V _| |_ V V   ___| |_ __ _| |_ ___ _ __ ___   ___ _ __ | |_
    | __| '_ \ / _ \     | |  _|     / __| __/ _` | __/ _ \ '_ ` _ \ / _ \ '_ \| __|
    | |_| | | |  __/     | | |       \__ \ || (_| | ||  __/ | | | | |  __/ | | | |_
     \__|_| |_|\___|     |_|_|       |___/\__\__,_|\__\___|_| |_| |_|\___|_| |_|\__|

    The timeless & venerable if statement... Let's make a for loop that lists
    all even numbers from 1 to 10
    */

    // let's make a loop where z goes from 1 to 10
    for z := 1; z <= 10; z++ {
        // let's check if we divide by 2 we get 0 as a remainder indicating an
        // even number. remember the % operator returns the remainder of an
        // integer divide of said operators. 14 % 4 would return 2 and 4 goes
        // into fourteen 3 times and leaves a remainder of 2.
        if z % 2 == 0 {
            // if it is
            fmt.Println(z, "is an even number")
        } else {
            fmt.Println(z, "is an odd number")
        }
    }

    /*
                       _                _   _                     _
                      | |              | | | |                   | |
     _ __   ___    ___| |__   ___  _ __| |_| |__   __ _ _ __   __| |
    | '_ \ / _ \  / __| '_ \ / _ \| '__| __| '_ \ / _` | '_ \ / _` |
    | | | | (_) | \__ \ | | | (_) | |  | |_| | | | (_| | | | | (_| |
    |_| |_|\___/  |___/_| |_|\___/|_|   \__|_| |_|\__,_|_| |_|\__,_|


    If you are coming from other languages you might assume go supports
    shorthand for one liner loops and if statements such as follows....

    */

    /*
    // unblock comment me to see a compile failure!
    for w := 1; w <= 10; w++
        if w % 2 == 0
            fmt.Println(w, "is an even number")
        else
            fmt.Println(w, "is an odd number")

    */

    /*

    ...however go will fail to compile and complain about missing { braces

     _ _             _ _       _    _ _            _ _                  _ _
    ( | )           (_) |     | |  ( | )   ___    ( | )                ( | )
     V V_____      ___| |_ ___| |__ V V   ( _ )    V V ___ __ _ ___  ___V V
       / __\ \ /\ / / | __/ __| '_ \      / _ \/\     / __/ _` / __|/ _ \
       \__ \\ V  V /| | || (__| | | |    | (_>  <    | (_| (_| \__ \  __/
       |___/ \_/\_/ |_|\__\___|_| |_|     \___/\/     \___\__,_|___/\___|

    Go supports the the switch, case and default statements found in other
    langauges. Go does not however use the break statement found in many
    languages and instead simply executes the code between case statements.
    Go also supports boolean switches where each case is a condition.

    */

    // Let's create a loop that goes through and prints out the english names
    // for some numbers
    for w := 0; w <= 13; w++ {
        // Let's begin each iteration of our for loop by displaying this
        // simple text to show the number whose name we will display...
        fmt.Print("The name of ",w ," is ")

        // Next we will use a switch statement to print out the name of said
        // number. If a matching number is not found the switch will isntead
        // execute the code found in the special default case.
        switch w {
            case 0: fmt.Println("Zero")
            case 1: fmt.Println("One")
            case 2: fmt.Println("Two")
            case 3: fmt.Println("Three")
            // case 4: fmt.Println("Four")  //comment this one out to show it will
                                            // trigger Unknown Number
            case 5: fmt.Println("Five")
            case 6: fmt.Println("Six")
            case 7: fmt.Println("Seven")
            case 8: fmt.Println("Eight")
            case 9: fmt.Println("Nine")
            case 10: fmt.Println("Ten")

            // if none of the conditions above are met print out a message
            // letting the user know the number is undefined.
            default: fmt.Println("Undefined")
        }
    }

    /*

    We can also perform a switch without providing an input and perform a
    boolean test on each case.

    */

    // Create a time object equal to the current time.
    current_time := time.Now()

    // Switch into one of 6 hour ranges
    switch {
        case current_time.Hour() < 6: fmt.Println ("It's before 6 am")
        case current_time.Hour() < 12: fmt.Println ("It's between 6am and noon")
        case current_time.Hour() < 18: fmt.Println ("It's between noon and 6 pm")
        default: fmt.Println ("It's after 6pm")
    }

    // We can also easily test multiple cases at the same time using a comma
    v := 4
    switch v {
        case 1,2,3: fmt.Println("v is in the first case")
        case 4,5,6: fmt.Println("v is in the second case")
        default: fmt.Println("v was not found in either group")
    }
}
