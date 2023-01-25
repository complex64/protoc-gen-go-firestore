package gen

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFile(
	fs flag.FlagSet,
	plugin *protogen.Plugin,
	proto *protogen.File,
	packages *Packages,
) error {
	log.Trace().Str("file", proto.Desc.Path()).Msg("GenerateFile()")
	defer log.Trace().Msg("-----------------------------------------------------------------------")

	file, err := NewFile(plugin, proto)
	if err != nil {
		return err
	}
	packages.Collect(file)
	file.Gen()
	return nil
}

func Comment(format string, args ...interface{}) protogen.Comments {
	return protogen.Comments(fmt.Sprintf(format, args...))
}

func appendDeprecationNotice(prefix protogen.Comments, deprecated bool) protogen.Comments {
	if !deprecated {
		return prefix
	}
	if prefix != "" {
		prefix += "\n"
	}
	return prefix + " Deprecated: Do not use.\n"
}
