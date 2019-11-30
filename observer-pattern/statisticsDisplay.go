package observer_pattern

import "fmt"

type StatisticsDisplay struct {
	Observable  Observable
	Humidity    float64 //湿度
	Temperature float64 //温度
	Pressure    float64 //气压
}

func NewStatisticsDisplay(observable Observable) Observer {
	statisticsDisplay := &StatisticsDisplay{
		Observable: observable,
	}
	statisticsDisplay.Observable.RegisterObserver(statisticsDisplay)
	return statisticsDisplay
}

func (statisticsDisplay *StatisticsDisplay) Update(observable Observable) {
	weatherData := observable.(*WeatherData)
	statisticsDisplay.Humidity = weatherData.Humidity
	statisticsDisplay.Temperature = weatherData.Temperature
	statisticsDisplay.Pressure = weatherData.Pressure
	statisticsDisplay.Display()
}

func (statisticsDisplay *StatisticsDisplay) Display() {
	fmt.Printf("统计显示 => 温度:%f;湿度:%f;气压:%f;\n", statisticsDisplay.Temperature, statisticsDisplay.Humidity, statisticsDisplay.Pressure)
}
