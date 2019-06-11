package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		levelKey string
		level    string
		schedKey string
	)
	flag.StringVar(&levelKey, "level_key", "severity", "GODEBUG level key")
	flag.StringVar(&level, "level", "debug", "GODEBUG log level")
	flag.StringVar(&schedKey, "sched_key", "sched", "GODEBUG schedtrace key")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if strings.HasPrefix(s, "SCHED") {
			s = `{"` + levelKey + `":"` + level + `","` + schedKey +
				`":"` + s[:len(s)-1] + "\"}\n"
		}
		w.WriteString(s)
		w.Flush()
	}
}
