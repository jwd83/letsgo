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

*/
import "fmt"

func main() {
    /*
     _      _   _       _ _                     _ _
    | |    | | ( )     ( | )                   ( | )
    | | ___| |_|/ ___   V V_ __ ___   __ _ _ __ V V
    | |/ _ \ __| / __|    | '_ ` _ \ / _` | '_ \
    | |  __/ |_  \__ \    | | | | | | (_| | |_) |
    |_|\___|\__| |___/    |_| |_| |_|\__,_| .__/
                                          | |
                                          |_|

    A map in Go is an unordered collection of key value pairs. In other
    languages these are often called associative arrays, has tables or
    dictionaries. Maps are defined by both their key type AND their value type.
    For instance you could have a map that is indexed by a string whose value
    is an integer and vice versa. You can use len to count the number of items
    in a map. A simple declaration for a map is...

    */

    my_map := make(map[string]int)

    // You can then add items to the map...
    my_map["John"] = 27

    // You can then access the stored value by using it's key
    fmt.Println("John =", my_map["John"])

    // Display the count
    fmt.Println("len of my_map:", len(my_map))

    // Display the contents of the map
    fmt.Println("content of my_map:", my_map)

    // Let's create another entry
    my_map["Mike"] = 44

    // Display the count and contents again
    fmt.Println("len of my_map:", len(my_map))
    fmt.Println("content of my_map:", my_map)

    // Now lets remove an item
    delete(my_map, "John")

    // Display the count and contents again
    fmt.Println("len of my_map:", len(my_map))
    fmt.Println("content of my_map:", my_map)

    /*

         _                           _                          _     _  ___
        | |                         | |                        (_)   | ||__ \
      __| | ___   ___  ___    __ _  | | _____ _   _    _____  ___ ___| |_  ) |
     / _` |/ _ \ / _ \/ __|  / _` | | |/ / _ \ | | |  / _ \ \/ / / __| __|/ /
    | (_| | (_) |  __/\__ \ | (_| | |   <  __/ |_| | |  __/>  <| \__ \ |_|_|
     \__,_|\___/ \___||___/  \__,_| |_|\_\___|\__, |  \___/_/\_\_|___/\__(_)
                                               __/ |
                                              |___/

    John has now been removed from the map; however, if we try to access John
    we don't get an error, it simply returns a zero value, that is literally a
    zero value for numbers and an emptry string for strings.
    */

    fmt.Println("John =", my_map["John"])

    /*
    To check whether or not a key exists Go returns a boolean as an optional
    second variable returned when accesssing members of a map that is true/false
    depending on whether or not the key exists. When accessing a key, provide
    a placeholder for the boolean as well as follows...
    */
    num, ok := my_map["John"]
    fmt.Println("num =", num, "ok =", ok)

    // If we attempt to access Mike we will see that ok returns true.
    num, ok = my_map["Mike"]
    fmt.Println("num =", num, "ok =", ok)

    /*
    As we have seen before we can use _ to omit the return of the value and
    only return whether or not the key was found.
    */
    _, ok = my_map["Mike"]
    fmt.Println("ok =", ok)

    /*
    To bring this full circle we can also setup an if statement to test whether
    or not a key exists simply by...
    */

    if _, ok = my_map["John"]; ok {
        fmt.Println("John was found")
    } else {
        fmt.Println("John was not found")
    }

    /*
    We can also define a variable to be used within the if statement's scope.
    Also of note it appears we can mix declared and undeclared variables in a
    := statement if it's the first time being used for either
    */
    if age, ok := my_map["Mike"]; ok {
        fmt.Println("Mike was found, age =", age)
    } else {
        fmt.Println("Mike was not found")
    }

    /*
    Age is undeclared again outside of the if statement. Sicne it is out of
    scope we can recreate it here...
    */
    age := 20

    /*
    Further inside the scope of the if statement it is using it's own copy of
    age if we use the := operator to define variables used for the key lookup
    */

    fmt.Println("age outside of scope =", age)

    if age, found := my_map["Mike"]; found {
        fmt.Println("Mike was found again! age =", age)
    } else {
        fmt.Println("Mike was not found again!")
    }

    fmt.Println("age outside of scope again =", age)
    /*
         _                _   _                     _
        | |              | | | |                   | |
     ___| |__   ___  _ __| |_| |__   __ _ _ __   __| |  _ __ ___   __ _ _ __  ___
    / __| '_ \ / _ \| '__| __| '_ \ / _` | '_ \ / _` | | '_ ` _ \ / _` | '_ \/ __|
    \__ \ | | | (_) | |  | |_| | | | (_| | | | | (_| | | | | | | | (_| | |_) \__ \
    |___/_| |_|\___/|_|   \__|_| |_|\__,_|_| |_|\__,_| |_| |_| |_|\__,_| .__/|___/
                                                                       | |
                                                                       |_|
    It's also possible to quickly define maps. Again we will need the trailing
    command if we want to use a multliline definition.
    */

    quick_map := map[string]int{
        "John": 40,
        "Mike": 42,
        "Mark": 53,
        "Paul": 55,
        "Jimbob": 63,
    }
    fmt.Println(quick_map)

    /*
                                          _   _
                                         | | (_)
     _ __ ___   __ _ _ __   ___ ___ _ __ | |_ _  ___  _ __
    | '_ ` _ \ / _` | '_ \ / __/ _ \ '_ \| __| |/ _ \| '_ \
    | | | | | | (_| | |_) | (_|  __/ |_) | |_| | (_) | | | |
    |_| |_| |_|\__,_| .__/ \___\___| .__/ \__|_|\___/|_| |_|
                    | |            | |
                    |_|            |_|

    In addition to simple key value pairs maps can also map to... maps! Here's
    a shorthand example of a map to a map definition...
    */

    students := map[string]map[string]int{
        "Jim": map[string]int{
            "Age": 21,
            "Height(cm)": 185,
            "Weight(kb)": 65,
        },
        "Dwight": map[string]int{
            "Age": 23,
            "Height(cm)": 175,
            "Weight(kb)": 72,
        },
        "Pam": map[string]int{
            "Age": 22,
            "Height(cm)": 167,
            "Weight(kb)": 49,
        },
        "Kevin": map[string]int{
            "Age": 22,
            "Height(cm)": 167,
            "Weight(kb)": 120,
        },
    }
    fmt.Println(students)

    if details, ok := students["Pam"]; ok {
        fmt.Println("Found Pam. She is",details["Age"],"years old.")
        fmt.Println("Pam's full details", details)
    }

    fmt.Println("Number of students:", len(students))
}