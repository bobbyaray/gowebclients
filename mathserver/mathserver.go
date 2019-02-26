package main

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello World")
		ctx.View("index.html")
	})

	app.Get("/fibonacci/{size:int}", func(ctx iris.Context) {
		ctx.ViewData("function", "Fibonacci")
		size, _ := ctx.Params().GetInt("size")
		ctx.ViewData("params", fmt.Sprintf("Size: %d", size))
		ctx.ViewData("results", fibonacci(size))
		ctx.View("results.html")
	})

	app.Get("/squared/{num:int}", func(ctx iris.Context) {
		ctx.ViewData("function", "Squared")
		num, _ := ctx.Params().GetInt("num")
		ctx.ViewData("params", fmt.Sprintf("Number: %d", num))
		ctx.ViewData("results", strconv.Itoa(num*num))
		ctx.View("results.html")
	})

	app.Get("/factorial/{basenum:int}", func(ctx iris.Context) {
		ctx.ViewData("function", "Factorial")
		num, _ := ctx.Params().GetInt("basenum")
		ctx.ViewData("params", fmt.Sprintf("Base Number: %d", num))
		ctx.ViewData("results", factorial(num))
		ctx.View("results.html")
	})

	app.Run(iris.Addr(":9080"))
}

func factorial(baseNum int) int64 {
	var total int64
	for x := baseNum; x > 0; x-- {
		if total == 0 {
			total = int64(x)
		} else {
			total = total * int64(x)
		}
	}

	return total
}

func fibonacci(count int) string {
	var seq string
	var x, y int = 0, 1
	for z := 0; z < count; z++ {
		sum := x + y
		if len(seq) > 0 {
			seq = seq + "," + strconv.Itoa(x)
		} else {
			seq = strconv.Itoa(x)
		}

		x = y
		y = sum
	}

	seq = seq + "," + strconv.Itoa(x)
	return seq
}
