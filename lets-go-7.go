package main

import "fmt"

func main() {
    fmt.Println(func1())
    fmt.Println(average_float64([]float64{1,2,3,4,5,6}))
    fmt.Println(average_float64([]float64{}))
    fmt.Println(average_float64_checked([]float64{}))

}

func func1() float64 {
    return func2()
}

func func2() float64 {
    return 1.1
}

func average_float64(the_floats []float64) float64 {
    total := 0.0
    if len(the_floats) > 0 {
        for _,value := range the_floats {
            total += value
        }
        return total / float64(len(the_floats))
    } else {
        return 0
    }
}


func average_float64_checked(the_floats []float64) [float64, bool] {
    total := 0.0
    if len(the_floats) > 0 {
        for _,value := range the_floats {
            total += value
        }
        return total / float64(len(the_floats)), true
    } else {
        return 0, false
    }
}