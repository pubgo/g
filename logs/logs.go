package logs

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
)

// Debug
// log
func Debug(d ...interface{}) {
	for i := range d {
		fmt.Printf("%#v", d[i])
	}
}

// P
// log
func P(s string, d ...interface{}) {
	fmt.Print(s)
	for _, i := range d {
		if i == nil || _isNone(i) {
			continue
		}

		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dt))
	}
}

func Match(e string) bool {
	return zerolog.DebugLevel.String() == e ||
		zerolog.ErrorLevel.String() == e ||
		zerolog.WarnLevel.String() == e ||
		zerolog.FatalLevel.String() == e ||
		zerolog.InfoLevel.String() == e ||
		zerolog.PanicLevel.String() == e
}

// DebugLog init log
func DebugLog(key, value string) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str(key, value).Caller().Timestamp().Logger()
}
