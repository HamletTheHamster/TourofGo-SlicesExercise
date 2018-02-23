package main

import (
    "../tour/pic"
    "math"
)

func Pic(dx, dy int) [][]uint8 {
    // Allocate 2D array
    picture := make([][]uint8, dy)
    for j := 0; j < dy; j++ {
        picture[j] = make([]uint8, dx)
    }
    
    // Make picture
    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            //picture[y][x] = uint8((x+y)/2)
            //picture[y][x] = uint8(x*y)
            //picture[y][x] = uint8(x^y)
            
            //picture[y][x] = uint8(math.Sin(float64(y)) + math.Sin(float64(x)))
            //picture[y][x] = uint8(x*x + y*y)
            picture[y][x] = uint8(math.Pi*float64(x*y^(x+y)))
            }
        }
    return picture
}

func main() {
    pic.Show(Pic)
}
