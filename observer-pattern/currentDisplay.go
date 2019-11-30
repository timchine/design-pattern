package observer_pattern

import "fmt"

//现在显示
type CurrentDisplay struct {
	Observable  Observable
	Temperature float64 //温度
	Pressure    float64 //气压
}

func NewCurrentDisplay(observable Observable) CurrentDisplay {
	currentDisplay := CurrentDisplay{Observable: observable}
	currentDisplay.Observable.RegisterObserver(&currentDisplay)
	return currentDisplay
}

func (currentDisplay *CurrentDisplay) Display() {
	fmt.Printf("现在显示 => 温度: %f摄氏度;气压为:%f\n", currentDisplay.Temperature, currentDisplay.Pressure)
}

func (currentDisplay *CurrentDisplay) Update(observable Observable) {
	weatherData := observable.(*WeatherData)
	currentDisplay.Pressure = weatherData.Pressure
	currentDisplay.Temperature = weatherData.Temperature
	currentDisplay.Display()
}
