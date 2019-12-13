package main

import adapter_pattern "github.com/echoloveyou/design-patterns/adapter-pattern"

func main()  {
	adapt := adapter_pattern.Adapt{}
	adapt.SetAdaptee(&adapter_pattern.Adaptee{})
	adapt.Request()

	classAdapt := adapter_pattern.ClassAdapt{}
	classAdapt.Request()
}
