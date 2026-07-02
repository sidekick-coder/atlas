package config 

type Config struct {
	Entries map[string]string
}

func New(entries map[string]string) *Config {
	return &Config{
		Entries: entries,
	}
}

func (c *Config) Get(key string) (string, bool) {
	value, exists := c.Entries[key]
	return value, exists
}

func (c *Config) Set(key, value string) {
	c.Entries[key] = value
}
