package chess

import "fmt"

type Colour uint8

const (
	black Colour = iota
	white
)

func Chess() {
	b := StartPosition()
	fmt.Println(b)
}
