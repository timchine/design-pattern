package strategy_pattern

type Duck struct {
	FlyBehavior
	QuackBehavior
}

func (duck Duck)PerformFly ()  {
	duck.Fly()
}

func (duck Duck)PerformQuack() {
	duck.Quack()
}

func (duck *Duck) SetQuackBehavior(behavior QuackBehavior) {
	duck.QuackBehavior = behavior
}

func (duck *Duck) SetFlyBehavior(behavior FlyBehavior)  {
	duck.FlyBehavior = behavior
}

type ModeDuck struct {
	Duck
}

func NewModeDuck() ModeDuck {
	duck := Duck{
		FlyBehavior: new(FlyNoWay),
		QuackBehavior: new(MuteQuack),
	}
	return ModeDuck{duck}
}

func Behavior() {
	modeDuck := NewModeDuck()
	modeDuck.Quack()
	modeDuck.Fly()
	modeDuck.SetFlyBehavior(new(FlyWithWings))
	modeDuck.Fly()
}