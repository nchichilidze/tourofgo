package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	/* declare variables */
	var i,j int
	
	/* initialize slice*/
	picture := make([][]uint8,dy)
	for i = 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)
	}
	
	/* populate slice */
	for i = 0; i < dy; i++ {
		for j = 0; j < dx; j++ { 
			value := i*i + 2*i*j + j*j
			picture[i][j] = uint8(value)
		}
	}
	
	/* return result */
	return picture
}

func main() {
	pic.Show(Pic)
}
