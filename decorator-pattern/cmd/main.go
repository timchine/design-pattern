package main

import (
	"fmt"
	"github.com/echoloveyou/design-patterns/decorator-pattern"
)

func main() {
	//new一个麻辣锅底的火锅
	hotPot := decorator_pattern.NewSpicyHotSoupBottom("麻辣锅底", 100.00)
	celery := decorator_pattern.CeleryGarnish{
		Garnish: decorator_pattern.Garnish{
			Description: "芹菜",
			Price:       10.00,
		},
	}
	cabbage := decorator_pattern.CabbageGarnish{
		Garnish: decorator_pattern.Garnish{
			Description: "白菜",
			Price:       7.00,
		},
	}
	//加芹菜
	hotPot = celery.AddGarnishTo(hotPot)
	//加白菜
	hotPot = cabbage.AddGarnishTo(hotPot)
	fmt.Println(hotPot.Prices())
	fmt.Println(hotPot.GetDescription())

	//第一个火锅做好了。
	hotPot1 := decorator_pattern.NewSpicyHotSoupBottom("排骨锅底", 110.00)
	beef := decorator_pattern.BeefGarnish{
		Garnish: decorator_pattern.Garnish{
			Description: "牛肉",
			Price:       30.00,
		},
	}
	greens := decorator_pattern.GreensGarnish{
		Garnish: decorator_pattern.Garnish{
			Description: "青菜",
			Price: 5.00,
		},
	}
	hotPot1 = beef.AddGarnishTo(hotPot1)
	hotPot1 = greens.AddGarnishTo(hotPot1)
	fmt.Println(hotPot1.Prices())
	fmt.Println(hotPot1.GetDescription())
}
