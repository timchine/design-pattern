package main

import observer_pattern "github.com/echoloveyou/design-patterns/observer-pattern"

func main()  {
	weatherData := new(observer_pattern.WeatherData)
	//new三个观察者
	observer_pattern.NewCurrentDisplay(weatherData)
	observer_pattern.NewStatisticsDisplay(weatherData)
	observer_pattern.NewForecastDisplay(weatherData)

	weatherData.SetWeatherData(12.0, 123.1, 13)
	weatherData.SetWeatherData(13.0, 12.1, 15)
}