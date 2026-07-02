package config 

type Config struct {
	entries map[string]string
}

func New(entries map[string]string) *Config {
	return &Config{entries: entries}
}

func (c *Config) Get(key string) string {
	return c.entries[key]
}

func (c *Config) Set(key, value string) {
	c.entries[key] = value
}

func (c *Config) GetAll() map[string]string {
	return c.entries
}
