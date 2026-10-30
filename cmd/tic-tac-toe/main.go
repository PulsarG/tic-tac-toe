package main

import (
	"fmt"
	"tic-tack-toe/internal/business"
	"tic-tack-toe/internal/logger"

	"go.uber.org/fx"
)

func run(calc *business.Calc) {
	res := calc.Add(5, 3)
	fmt.Println("Res is: ", res)
}

func main() {
	fx.New(
		fx.Provide(
			logger.NewLogger,
			business.NewCalc,
		),
		fx.Invoke(run),
	).Run()
}
