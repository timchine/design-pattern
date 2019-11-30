package observer_pattern

import "fmt"

//预测显示
type ForecastDisplay struct {
	Observable  Observable
	Temperature float64 //温度
	Humidity    float64 //湿度
}

func NewForecastDisplay(observable Observable) Observer {
	forecastDisplay := &ForecastDisplay{Observable: observable}
	forecastDisplay.Observable.RegisterObserver(forecastDisplay)
	return forecastDisplay
}

func (forecastDisplay *ForecastDisplay) Display() {
	fmt.Printf("预测显示 => 温度: %f摄氏度;湿度为:%f\n", forecastDisplay.Temperature, forecastDisplay.Humidity)
}

func (forecastDisplay *ForecastDisplay) Update(observable Observable) {
	weatherData := observable.(*WeatherData)
	forecastDisplay.Humidity = weatherData.Humidity
	forecastDisplay.Temperature = weatherData.Temperature
	forecastDisplay.Display()
}
