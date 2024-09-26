package config

import (
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var viperI *viper.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	viperI = viper.New()
	viperI.SetConfigType("env")
	viperI.AddConfigPath(".")
	viperI.SetEnvPrefix("appenv")
	viperI.AutomaticEnv()
	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEvn(env)
}
func loadConfig() {
	for name, fn := range ConfigFuncs {
		viperI.Set(name, fn)
	}
}
func loadEvn(env string) {
	var envPath = ".env"
	if len(env) > 0 {
		filePath := env + envPath
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}
	viperI.SetConfigName(envPath)
	if err := viperI.ReadInConfig(); err != nil {
		panic(err)
	}
	viperI.WatchConfig()
}

func Env(name string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(name, defaultValue[0])
	} else {
		return internalGet(name)
	}
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viperI.Get(path)
}
func Add(name string, fn ConfigFunc) {
	ConfigFuncs[name] = fn
}

func Get(name string, defaultValue ...interface{}) string {
	return GetString(name, defaultValue...)
}

func GetString(name string, defultValue ...interface{}) string {
	return cast.ToString(internalGet(name, defultValue...))
}

func GetInt(name string, defultValue ...interface{}) int {
	return cast.ToInt(internalGet(name, defultValue...))
}

func GetFloat64(name string, defultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(name, defultValue...))
}

func GetBool(name string, defultValue ...interface{}) bool {
	return cast.ToBool(internalGet(name, defultValue...))
}
