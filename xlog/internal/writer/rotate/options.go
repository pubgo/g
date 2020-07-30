package rotate

import (
	"os"
	"time"
)

type Option interface {
	apply(*Config)
}

type optionFunc func(*Config)

func (f optionFunc) apply(cfg *Config) {
	f(cfg)
}

func WithLogDir(dir string) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Dir = dir
	})
}

func WithLogSubDir(dir string) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Sub = dir
	})
}

func WithFilename(fn string) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Filename = fn
	})
}

func WithFileMode(Perm os.FileMode) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Perm = Perm
	})
}

func WithAge(Age time.Duration) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Age = Age
	})
}

func WithDuration(Duration time.Duration) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Duration = Duration
	})
}

// sets the number of files should be kept before it gets purged from the file system.
func WithCount(Count uint) Option {
	return optionFunc(func(cfg *Config) {
		// options Age and Count cannot be both set
		cfg.Count = Count
		cfg.Age = 0
	})
}

func WithPattern(Pattern string) Option {
	return optionFunc(func(cfg *Config) {
		cfg.Pattern = Pattern
	})
}

func WithLocation(loc *time.Location) Option {
	return optionFunc(func(cfg *Config) {
		if loc != nil {
			cfg.Loc = loc
		}
	})
}
