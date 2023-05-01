package config

import "github.com/spf13/viper"

func Get(key string) interface{} { return viper.Get(key) }

func GetString(key string) string { return viper.GetString(key) }

func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

func GetInt(key string) int { return viper.GetInt(key) }

func GetInt32(key string) int32 { return viper.GetInt32(key) }

func GetInt64(key string) int64 { return viper.GetInt64(key) }

func GetIntSlice(key string) []int { return viper.GetIntSlice(key) }

func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

func GetBool(key string) bool { return viper.GetBool(key) }

func Unmarshal(rawVal interface{}) error { return viper.Unmarshal(rawVal) }
