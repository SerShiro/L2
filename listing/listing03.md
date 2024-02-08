package main

import (
"fmt"
"os"
)

func Foo() error {
var err *os.PathError = nil
return err
}

func main() {
err := Foo()
fmt.Println(err)
fmt.Println(err == nil)
}

В выводе будет nil false, в первом случае выводится nil, это значение переменной, которое мы положили.
При сравнении с nil будет false, потому что у нас значение переменной nil, а тип данных *os.PathError, поэтому будет false