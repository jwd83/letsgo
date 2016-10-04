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

Diving into arrays
*/

import "fmt"

func main() {

    /*
      __ _ _ __ _ __ __ _ _   _ ___
     / _` | '__| '__/ _` | | | / __|
    | (_| | |  | | | (_| | |_| \__ \
     \__,_|_|  |_|  \__,_|\__, |___/
                           __/ |
                          |___/

    Let's take a look at how some simple arrays function
    */

    // let's create an array containing 5 integers and print it to the screen

    var x [5] rune  // rune's an alternate for int32, just a reminder for
                    // good practice

    fmt.Println(x)  // This should yield [0 0 0 0 0] as Go defaults all values
                    // to 0 when creating an array of integers or floats


    // arrays start at index 0 so valid addresses for this array are 0.1.2.3.4
    x[0] = 1
    x[1] = 2
    x[2] = 3
    x[3] = 4
    x[4] = 5        // This puts the number 5 in the fifth spot in this array
    fmt.Println(x)  // This will now yield [1 2 3 4 5]

    // x[5] = 5     // This would fail to compile as this exists outside the
                    // bounds of x. The block comment below WOULD compile but
                    // would cause a "panic: runtime error: index out of range"
                    // once running.

    /*
    for i := 0; i <= 5; i++ {
        fmt.Println(x[i])
    }
    */

    // To find the length of an array
    fmt.Println("The length of x is", len(x))

    // To access a specific element within an array...
    fmt.Println("The 3rd element of x is", x[2])

    // The last element of x is...
    fmt.Println("The last element of x is ", x[len(x) - 1])


    /*
         _                _   _                     _
        | |              | | | |                   | |
     ___| |__   ___  _ __| |_| |__   __ _ _ __   __| |
    / __| '_ \ / _ \| '__| __| '_ \ / _` | '_ \ / _` |
    \__ \ | | | (_) | |  | |_| | | | (_| | | | | (_| |
    |___/_| |_|\___/|_|   \__|_| |_|\__,_|_| |_|\__,_|

    As we saw with other variable types Go has a quick way to define them. Let's
    make a list of test scores and write a loop to average them.

    */

    test_scores1 := [5]float64 { 100, 75, 80, 75, 88 }
    total := 0.0


    for i := 0; i < len(test_scores1); i++ {
        total += test_scores1[i]
    }

    /*
    Note len will return an int and needs to be converted to a float64 before
    we can divide it with another float64. Trying to simply divide by len
    will return "invalid operation: total / len(test_scores1) (mismatched
    types float64 and int)" from the compiler
    */

    fmt.Println("Average of test_scores1: ", total / float64(len(test_scores1)))

    /*

         _ _
        | (_)
     ___| |_  ___ ___  ___
    / __| | |/ __/ _ \/ __|
    \__ \ | | (_|  __/\__ \
    |___/_|_|\___\___||___/

    We can also omit the size entirely giving us what is called a slice. The
    declaration for test_scores1 could be rewritten as a slice as follows...

    test_scores1 := []float64 { 100, 75, 80, 75, 88 }


    Additionally we can define an array or slice over multiple lines as such...
    */

    test_scores2 := []float64{
        44,
        55,
        80,
        100,
        100,            // Go requires a comma after the final element when
                        // using this format. Don't worry it's not adding
                        // an empty element.
    }


    // Perform our average calculation again
    total = 0
    for i := 0; i < len(test_scores2); i++ {
        total += test_scores2[i]
    }
    fmt.Println("Average of test_scores2: ", total / float64(len(test_scores2)))

    /*
     _   _            _ _  __                                       _ _
    | | | |          ( | )/ _|                                     ( | )
    | |_| |__   ___   V V| |_ ___  _ __   _ __ __ _ _ __   __ _  ___V V
    | __| '_ \ / _ \     |  _/ _ \| '__| | '__/ _` | '_ \ / _` |/ _ \
    | |_| | | |  __/     | || (_) | |    | | | (_| | | | | (_| |  __/
     \__|_| |_|\___|     |_| \___/|_|    |_|  \__,_|_| |_|\__, |\___|
                                                           __/ |
                                                          |___/
     _
    | |
    | | ___   ___  _ __
    | |/ _ \ / _ \| '_ \
    | | (_) | (_) | |_) |
    |_|\___/ \___/| .__/
                  | |
                  |_|

    The for range loop allows you to loop over  a slice or array and provides
    you the index position as well as it's value in two separate variables. If
    you only need one of these two variables in your loop you can omit the
    inclusion of it to your loop by providing an underscore "_" in place of that
    respective place. Go will not compile if you have unused variables so it's
    important to use _ here to avoid this. Here are a few examples...
    */

    // Using both the index postion and the value
    total = 0
    test_scores3 := []float64{ 99.0, 95.0, 95.0, 98.0 , 33.0}
    for index, value := range test_scores3 {
        fmt.Println("Student", index, "scored", value)
        total += value
    }
    fmt.Println("Average of test_scores3:", total / float64(len(test_scores3)))

    // Using only the value and omitting the index with an underscore _
    total = 0
    test_scores4 := []float64{0,1,2,3,4,5,6,7,8,9,10}
    for _,value := range test_scores4 {
        fmt.Println("Adding score of ", value)
        total += value
    }
    fmt.Println("Average of test_scores4:", total / float64(len(test_scores4)))

    // Using only the index and omitting the value
    total = 0
    test_scores5 := []float64{10,11,12,13,14,15,16,17,18,19,20}
    for index, _ := range test_scores5 {
        fmt.Println ("The index is ", index)
    }
}
