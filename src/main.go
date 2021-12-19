package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	parseArgs()

	confFile, _ := filepath.Abs(CLI.ConfigFile)
	conf := readConfig(confFile)

	fmt.Printf("Watch file %q\n", conf.FileToWatch)
	tailf(conf)
}
