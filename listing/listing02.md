package main

import (
"fmt"
)

func test() (x int) {
defer func() {
x++
}()
x = 1
return
}

func anotherTest() int {
var x int
defer func() {
x++
}()
x = 1
return x
}

func main() {
fmt.Println(test())
fmt.Println(anotherTest())
}
В выводе получим 2 1, потому что в первом случае мы увеличиваем х на 1, а потом возвращаем значение.
Во втором случае мы сначала возвращаем значение, а потом увеличиваем его на 1, но вернулась уже единица.