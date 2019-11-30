package observer_pattern

type Observer interface {
	Update(observable Observable)
}

type Display interface {
	Display()
}

type Observable interface {
	RegisterObserver(observer Observer) //观察者注册接口
	RemoveObserver(observer Observer) //删除观察者方法
	NotifyObservers() //发送消息给所有的观察者
}

type WeatherData struct {
	Observers   []Observer
	Temperature float64 //温度
	Humidity    float64 //湿度
	Pressure    float64 //气压
}

func (weatherData *WeatherData) RegisterObserver(observer Observer) {
	weatherData.Observers = append(weatherData.Observers, observer)
}

func (weatherData *WeatherData) RemoveObserver(observer Observer) {
	for k, v := range weatherData.Observers {
		if v == observer {
			weatherData.Observers = append(weatherData.Observers[:k], weatherData.Observers[k+1:]...)
			break
		}
	}
}

func (weatherData *WeatherData) NotifyObservers() {
	for _, observer := range weatherData.Observers {
		observer.Update(weatherData)
	}
}

func (weatherData *WeatherData) SetWeatherData(temperature, humidity, pressure float64) {
	weatherData.Pressure = pressure
	weatherData.Temperature = temperature
	weatherData.Humidity = humidity
	weatherData.WeatherChanged()
}

func (weatherData *WeatherData)WeatherChanged()  {
	weatherData.NotifyObservers()
}