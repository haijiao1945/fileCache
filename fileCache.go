package fileCache

type Cache struct {
	driver *Driver
}

func (c Cache) DB() Cache {
	c.driver = new(Driver)
	return c
}

func (c Cache) SetConfig(configPath string) {
	new(Config).Init(configPath)
}

func (c Cache) Get(key string) (string, error) {
	return c.driver.Read(key)
}

func (c Cache) Set(key string, value string) (bool, error) {
	return c.driver.Write(key, value, 0)
}

func (c Cache) SetEx(key string, value string, expire int) (bool, error) {
	return c.driver.Write(key, value, expire)
}
