package main

import (
"fmt"
)

func main() {
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4]
fmt.Println(b)
}

На экран будет выведено элементы слайса с 1 по 4 не включительно,
т.е 77 78 79