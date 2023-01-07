package config

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type Reader interface {
	Names() []string
	Bytes(name string) ([]byte, error)
}

type BinaryReader struct {
	names func() []string
	bytes func(string) ([]byte, error)
}

func NewBinaryReader(names func() []string, bytes func(string) ([]byte, error)) *BinaryReader {
	return &BinaryReader{
		names: names,
		bytes: bytes,
	}
}

func (b *BinaryReader) Names() []string {
	return b.names()
}

func (b *BinaryReader) Bytes(name string) ([]byte, error) {
	return b.bytes(name)
}

type Config struct {
	*viper.Viper
	runtimeEnv string
}

var c *Config = New()

func New() *Config {
	c := &Config{
		Viper: viper.New(),
	}
	return c
}

func SetRuntimeEnv(s string) *Config {
	return c.SetRuntimeEnv(s)
}

func (c *Config) SetRuntimeEnv(s string) *Config {
	c.runtimeEnv = s
	return c
}

func Read(b Reader) *Config {
	return c.Read(b)
}

func (c *Config) Read(b Reader) *Config {
	v := viper.New()
	for _, name := range b.Names() {
		extWithPoint := filepath.Ext(name)
		if extWithPoint == "" {
			continue
		}
		ext := extWithPoint[1:]
		if !c.isExtValid(ext) {
			continue
		}
		v.SetConfigType(ext)
		data, err := b.Bytes(name)
		if err != nil {
			panic(err)
		}
		reader := bytes.NewReader(data)
		if err := v.MergeConfig(reader); err != nil {
			panic(err)
		}
	}

	// merge runtime config
	if c.runtimeEnv != "" {
		runtimeViper := v.Sub(fmt.Sprintf("<%s>", c.runtimeEnv))
		if runtimeViper != nil {
			if err := v.MergeConfigMap(runtimeViper.AllSettings()); err != nil {
				panic(err)
			}
		}
	}

	settings := v.AllSettings()
	for k := range settings {
		if k[0] == '<' && k[len(k)-1] == '>' {
			delete(settings, k)
		}
	}
	if err := c.MergeConfigMap(settings); err != nil {
		panic(err)
	}

	return c
}

func (c *Config) isExtValid(ext string) bool {
	for _, e := range viper.SupportedExts {
		if ext == e {
			return true
		}
	}
	return false
}

func ReadBinary(names func() []string, bytes func(string) ([]byte, error)) *Config {
	return c.ReadBinary(names, bytes)
}

func (c *Config) ReadBinary(names func() []string, bytes func(string) ([]byte, error)) *Config {
	c.Read(NewBinaryReader(names, bytes))
	return c
}
