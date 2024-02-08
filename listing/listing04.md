package main

func main() {
ch := make(chan int)
go func() {
for i := 0; i < 10; i++ {
ch <- i
}
}()
	for n := range ch {
		println(n)
	}
}
В выводе будут числа от 0 до 9, и потом случается дедлок. Это происходит потому что горутина пишет в канал чисел и после окончания записи канал не закрывается. Главная горутина читает числа из канала и когда чисел нет, она становится в ожидание, но чисел больше не будет и происходит дедлок.