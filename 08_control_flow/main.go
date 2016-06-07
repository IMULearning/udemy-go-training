package main

import (
	forLoop "github.com/davidiamyou/udemy-training/08_control_flow/for_loop"
	"github.com/davidiamyou/udemy-training/08_control_flow/rune"
	switchCase "github.com/davidiamyou/udemy-training/08_control_flow/switch_case"
	"fmt"
	"github.com/davidiamyou/udemy-training/08_control_flow/condition"
)

func main() {
	forLoop.RunLoopType1()
	forLoop.RunLoopType2()
	forLoop.RunLoopType3()

	rune.LoopRune()

	fmt.Println(switchCase.RunSwitchCase("Tim"))
	fmt.Println(switchCase.RunSwitchOnType(32))

	condition.RunCondition()
}
