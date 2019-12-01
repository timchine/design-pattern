package singleton_pattern

import "sync"

type ConnConfig struct {
	conner   *Conner
	address  string //连接的地址
	userName string //
	password string //
}

var ConnConf *ConnConfig

//最简单的单例模式，单线程没有问题，如果是多线程依旧会产生多个对象
func ConnLazy(address, userName, password string) *ConnConfig {
	if ConnConf == nil {
		ConnConf = &ConnConfig{
			address:  address,
			userName: userName,
			password: password,
		}
		ConnConf.conner = Open(*ConnConf)
	}
	return ConnConf
}

var mu *sync.Mutex
//针对多线程简易版的上锁
func ConnLazyLock(address, userName, password string) *ConnConfig {
	mu.Lock()
	defer mu.Unlock()
	if ConnConf == nil {
		ConnConf = &ConnConfig{
			address:  address,
			userName: userName,
			password: password,
		}
		ConnConf.conner = Open(*ConnConf)
	}
	return ConnConf
}

//双重锁
func ConnDoubleLock(address, userName, password string) *ConnConfig {
	if ConnConf == nil {
		mu.Lock()
		defer mu.Unlock()
		if ConnConf == nil {
			ConnConf = &ConnConfig{
				address:  address,
				userName: userName,
				password: password,
			}
			ConnConf.conner = Open(*ConnConf)
		}
	}
	return ConnConf
}

var once *sync.Once
//用go的sync.Once()实现 推荐使用
func ConnOnce(address, userName, password string) *ConnConfig {
	once.Do(func() {
		ConnConf = &ConnConfig{
			address:  address,
			userName: userName,
			password: password,
		}
		ConnConf.conner = Open(*ConnConf)
	})
	return ConnConf
}