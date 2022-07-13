// Package intset provides a specialized set for integers or runes
package intset

// Config defines the configuration for creating a new set
type Config struct {
	bucketSize   int
	bucketGrowBy int
}

// New Config creates a new config with usable defaults
func NewConfig() *Config {
	return &Config{bucketSize: 4, bucketGrowBy: 1}
}

// BucketSize sets the initial bucket size
func (c *Config) BucketSize(size uint32) *Config {
	c.bucketSize = int(size)
	return c
}

// BucketGrowBy sets the amount a bucket will grow by when full
func (c *Config) BucketGrowBy(growBy uint32) *Config {
	c.bucketGrowBy = int(growBy)
	return c
}

// Default is a default Config which favors probing performance
// at the cost of memory.
var Default = NewConfig()