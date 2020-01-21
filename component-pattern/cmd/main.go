package main

import (
	"fmt"
	component_pattern "github.com/echoloveyou/design-patterns/component-pattern"
)

func main() {
	menu := component_pattern.NewMenu("精选菜单", "本店精选")
	menu.Add(component_pattern.NewMenuItem("小炒肉", "就是香", 12))
	menu2 := component_pattern.NewMenu("素菜", "素菜")
	menu2.Add(component_pattern.NewMenuItem("土豆丝", "没得商量", 30))
	allMenu := component_pattern.NewMenu("总菜单", "全部")
	allMenu.Add(menu)
	allMenu.Add(menu2)
	fmt.Println(allMenu.String())
}
