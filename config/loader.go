package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tonet-me/tonet-core/logger"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"log"
	"log/slog"
	"strings"
)

const (
	defaultPrefix       = "TONET_"
	defaultDelimiter    = "."
	defaultSeparator    = "__"
	defaultYamlFilePath = "config.yml"
)

var c Config

type Option struct {
	Prefix       string
	Delimiter    string
	Separator    string
	YamlFilePath string
	CallbackEnv  func(string) string
}

// our environment variables must prefix with `EB_`
// for nested env should use `__` aka: TONET__MONGO_CLIENT__HOST.
func defaultCallbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, defaultPrefix))

	return strings.ReplaceAll(base, defaultSeparator, defaultDelimiter)
}

func init() {
	const op = richerror.OP("config.init")
	k := koanf.New(defaultDelimiter)

	// load default configuration from Default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	// load configuration from yaml file
	if err := k.Load(file.Provider(defaultYamlFilePath), yaml.Parser()); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	// load from environment variable
	if err := k.Load(env.Provider(defaultPrefix, defaultDelimiter, defaultCallbackEnv), nil); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}
}

func C() Config {
	return c
}

func New(opt Option) Config {
	const op = richerror.OP("config.New")

	k := koanf.New(opt.Separator)

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	if err := k.Load(file.Provider(opt.YamlFilePath), yaml.Parser()); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	if err := k.Load(env.Provider(opt.Prefix, opt.Delimiter, opt.CallbackEnv), nil); err != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))
	}

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	return c
}
