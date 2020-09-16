/**
 * @Time : 2020/9/16 4:34 PM
 * @Author : solacowa@gmail.com
 * @File : config
 * @Software: GoLand
 */

package foo

import (
	"strings"

	"github.com/Unknwon/goconfig"
)

const (
	SectionServer  = "server"
	SectionRedis   = "redis"
	SectionMysql   = "mysql"
	SectionService = "service"
	SectionCors    = "cors"
)

type Config struct {
	*goconfig.ConfigFile
	env string
}

func NewConfig(path string) (*Config, error) {
	// 处理配置文件

	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigFile: cfg}, nil
}

func (c *Config) Section(key string) string {
	return c.env + "." + key
}

func (c *Config) SetEnv(env string) {
	c.env = env
}
func (c *Config) GetEnv() string {
	return c.env
}

func (c *Config) GetString(section, key string) string {
	var val string
	val, _ = c.GetValue(section, key)
	return val
}

func (c *Config) GetStrings(section, key string) []string {
	val := c.GetString(section, key)
	return strings.Split(val, ",")
}

func (c *Config) GetInt(section, key string) int {
	val, _ := c.Int(section, key)
	return val
}

func (c *Config) GetBool(section, key string) bool {
	val, _ := c.Bool(section, key)
	return val
}
