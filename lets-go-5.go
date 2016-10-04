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

More fun with slices
*/

import "fmt"

func main() {
    /*

     _ _                _       _ _               _ _
    ( | )              | |     ( | )             | (_)
     V V_ __ ___   __ _| | _____V V    __ _   ___| |_  ___ ___
       | '_ ` _ \ / _` | |/ / _ \     / _` | / __| | |/ __/ _ \
       | | | | | | (_| |   <  __/    | (_| | \__ \ | | (_|  __/
       |_| |_| |_|\__,_|_|\_\___|     \__,_| |___/_|_|\___\___|


    Let's make a slice that has a size of 3. The make keyword in go sets up a
    slice that has a size of 3 in this example and is popuated with all zero
    values.
    */
    slice_one := make([]float64, 3)
    fmt.Println(slice_one)              // [0 0 0] 3 positions specified

    /*
    We can also specify a capacity for a slice using the make keyword. This
    value must be greater than or equal to the size.
    */

    slice_one = make([]float64, 5, 10)
    fmt.Println(slice_one)              // [0 0 0 0 0] only 5 values populated

    // Lets redefine that slice as 6 elements instead
    slice_one = []float64{1,2,3,4,5,6}
    fmt.Println(slice_one)              // [1 2 3 4 5 6]

    // Let's change the last element of the array
    slice_one[5] = 99
    fmt.Println(slice_one)




    /*
     _ _                                 _ _ _               _ _
    ( | )                               | ( | )             | (_)
     V V __ _ _ __  _ __   ___ _ __   __| |V V    __ _   ___| |_  ___ ___
        / _` | '_ \| '_ \ / _ \ '_ \ / _` |      / _` | / __| | |/ __/ _ \
       | (_| | |_) | |_) |  __/ | | | (_| |     | (_| | \__ \ | | (_|  __/
        \__,_| .__/| .__/ \___|_| |_|\__,_|      \__,_| |___/_|_|\___\___|
             | |   | |
             |_|   |_|

    We can append values to a slice. If the slice has room left to increase it's
    size up to the capacity it simply appends these items in the next available
    spots. If there is not capacity left to insert items into the slice a new
    slice is created, the exisinting values are copied into it and the and then
    the new values are copied into it afterwards. Note this incurs a performance
    penalty once a new slice is created. Lets take a look at the code below...

    */

    // First lets create a slice with a capacity of 10 and that has no values
    // yet assigned.
    slice_two := make([]float64, 0, 10)
    fmt.Println(slice_two)

    // Next lets append items one at a time and display the address of the
    // first element in the slice.
    for i := 0; i < 15; i++ {
        slice_two = append(slice_two, float64(i))
        /*
        Using %p with Printf allows us to print the address. Note it's important
        to look at the address of the first element. If we were to pass
        &slice_two instead we would not see an address change. This is beceause
        we are looking at the slice and not it's underlying array that contains
        the data.

        */
        fmt.Printf("First element @ %p", &slice_two[0])
        fmt.Println(", Slice =", slice_two)
    }

    /*

    You will see output similar to...

    First element @ 0x0xc4200120f0, Slice = [0]
    First element @ 0x0xc4200120f0, Slice = [0 1]
    First element @ 0x0xc4200120f0, Slice = [0 1 2]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4 5]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4 5 6]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4 5 6 7]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4 5 6 7 8]
    First element @ 0x0xc4200120f0, Slice = [0 1 2 3 4 5 6 7 8 9]
    First element @ 0x0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10]
    First element @ 0x0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11]
    First element @ 0x0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11 12]
    First element @ 0x0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
    First element @ 0x0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]

    You can see starting at the line where we add the integer 10 to the slice
    the address underlying array has changed as it no longer had enough memory
    reserved at that address.

    TODO: It seems Go simply doubles the size of the slice once it has exceeded
    it's capacity. Adjusting the above loop to run to 150 for instance....

    First element @ 0xc4200120f0, Slice = [0]
    First element @ 0xc4200120f0, Slice = [0 1]
    First element @ 0xc4200120f0, Slice = [0 1 2]
    First element @ 0xc4200120f0, Slice = [0 1 2 3]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4 5]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4 5 6]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4 5 6 7]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4 5 6 7 8]
    First element @ 0xc4200120f0, Slice = [0 1 2 3 4 5 6 7 8 9]
    First element @ 0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10] CHANGED
    First element @ 0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11]
    First element @ 0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11 12]
    First element @ 0xc4200a0000, Slice = [0 1 2 3 4 5 6 7 8 9 10 11 12 13] ...
    First element @ 0xc4200a0000, Slice = [0 1 2 ... 18 19]
    First element @ 0xc4200a2000, Slice = [0 1 2 ... 18 19 20] CHANGED...
    First element @ 0xc4200a2000, Slice = [0 1 2 ... 37 38 39]
    First element @ 0xc4200a6000, Slice = [0 1 2 ... 37 38 39 40] CHANGED...
    First element @ 0xc4200a6000, Slice = [0 1 2 ... 37 38 39 40 41]
    First element @ 0xc4200a6000, Slice = [0 1 2 ... 37 38 39 40 41 42]...
    First element @ 0xc4200a6000, Slice = [0 1 2 ... 76 77 78 79]
    First element @ 0xc4200b0000, Slice = [0 1 2 ... 76 77 78 79 80] CHANGED...

    */

    /*
     _                     _     _       _
    | |                _  | |   (_)     | |
    | | _____      __ (_) | |__  _  __ _| |__
    | |/ _ \ \ /\ / /     | '_ \| |/ _` | '_ \
    | | (_) \ V  V /   _  | | | | | (_| | | | |
    |_|\___/ \_/\_/   (_) |_| |_|_|\__, |_| |_|
                                    __/ |
                                   |___/

    Another way to  create  slices  is  to  use the [low : high] expression. low
    is the starting address offset in the slice and high is the address where to
    end it (without including itself.) For instance....

    */
    simple_array := [5]float64{1,2,3,4,5}
    slice_three := simple_array[0:5]
    fmt.Println("When slice_three = simple_array[0:5] slice_three is", slice_three)
    // When slice_three = simple_array[0:5] slice_three is [1 2 3 4 5]

    slice_three = simple_array[2:3]
    fmt.Println("When slice_three = simple_array[0:5] slice_three is", slice_three)
    // When slice_three = simple_array[0:5] slice_three is [3]

    /*
     _ _                     _ _               _ _
    ( | )                   ( | )             | (_)
     V V ___ ___  _ __  _   _V V    __ _   ___| |_  ___ ___
        / __/ _ \| '_ \| | | |     / _` | / __| | |/ __/ _ \
       | (_| (_) | |_) | |_| |    | (_| | \__ \ | | (_|  __/
        \___\___/| .__/ \__, |     \__,_| |___/_|_|\___\___|
                 | |     __/ |
                 |_|    |___/

    The copy keyword takes two arguments - destination "dst" and source "src".
    If the two slices are different sizes only the copy will only proceed until
    it has reached the end of either slice.

    */

    slice_four := []int {1,2,3,4,5}
    slice_five := make([]int, 3)

    copy(slice_five, slice_four)    //copy what is in four into five
    fmt.Println(slice_five)         // [1 2 3]

    copy(slice_four, slice_five)    //copy what is in five into four
    fmt.Println(slice_five)         // [1 2 3] - Note slice_four was resized.
                                    // It's later elements have been lost

}
