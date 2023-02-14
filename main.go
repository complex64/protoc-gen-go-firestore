package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/complex64/protoc-gen-go-firestore/v2/internal/gen"
	"github.com/complex64/protoc-gen-go-firestore/v2/internal/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	docURL = "https://github.com/complex64/protoc-gen-go-firestore"
)

var (
	showVersion = flag.Bool("version", false, "")
	showHelp    = flag.Bool("help", false, "")
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	flag.Parse()
	if *showVersion {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if *showHelp {
		fmt.Fprintf(os.Stdout, "Please see %s for usage information.\n", docURL)
		os.Exit(0)
	}

	var (
		flags flag.FlagSet
	)
	protogen.Options{
		// Calls `flags.Set(param, value)` for each from `--go_firestore_out=<param1>=<value1>,...`.
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		var (
			packages = gen.NewPackages(plugin)
		)
		for _, f := range plugin.Files {
			if f.Generate {
				if err := gen.GenerateFile(flags, plugin, f, packages); err != nil {
					return err
				}
			}
		}
		if err := packages.Gen(); err != nil {
			return err
		}
		return nil
	})
}
