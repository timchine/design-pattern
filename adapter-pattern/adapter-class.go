package adapter_pattern

type ClassAdapt struct {
	Adaptee
}

func (c ClassAdapt) Request() {
	c.EspecialRequest()
}

