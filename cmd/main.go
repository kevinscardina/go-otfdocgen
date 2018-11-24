package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/kevinscardina/go-otfdocgen/otfdocgen"
	"github.com/popmedic/go-color/colorize/tty"
)

const (
	defaultNamePrefix  = "u0x"
	defaultNameSuffix  = ""
	defaultSearchStart = math.MinInt16
	defaultEndStart    = math.MaxInt16
)

const (
	progressOut = "|/-\\"
)

var (
	infilepath  = flag.String("in", "", "OTF file to parse (default stdin)")
	outfilepath = flag.String("out", "", "File to put output, include an extension to use that template for generation (default stdout)")
	tmplpath    = flag.String("tmpl", "", "Template file for output (default MD)")
	searchStart = flag.Int("start", defaultSearchStart, "Glyph index to start searching for glyphs at (default "+strconv.Itoa(defaultSearchStart)+")")
	searchEnd   = flag.Int("end", defaultEndStart, "Glyph index to end searching for glyphs at")
	namePrefix  = flag.String("prefix", defaultNamePrefix,
		"Suffix to add to rune name when glyph name does not exist (default "+defaultNamePrefix+")")
	nameSuffix = flag.String("suffix", defaultNameSuffix,
		"Suffix to add to rune name when glyph name does not exist (default "+defaultNameSuffix+")")
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, tty.FgHiRed().Start(), "ERROR: ", tty.FgRed().Start(), err.Error(), tty.FgRed().End())
}

func main() {
	flag.Parse()

	otfDocGen, err := otfdocgen.NewOTFDocGen(*infilepath, *outfilepath,
		*tmplpath, *namePrefix, *nameSuffix, *searchStart, *searchEnd)
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	defer otfDocGen.Destroy()

	found, err := otfDocGen.Read(func(name string) {
		fmt.Fprintln(os.Stderr, tty.FgGreen().Start(), "\b-> reading:",
			tty.FgHiGreen().End(), name)
	},
		func(count, start, end int) {
			fmt.Fprintln(os.Stderr, tty.FgMagenta().Start(), " glyphs: ",
				count, "\n   range: ", start, "-", end, tty.FgMagenta().End())
		},
		func(r int) {
			fmt.Fprint(os.Stderr, "\b", string(progressOut[int(math.Abs(float64(r%len(progressOut))))]))
		}, printError)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, tty.FgGreen().Start(), "  found: ",
		tty.FgHiGreen().End(), found)

	err = otfDocGen.Write(func(name string) {
		fmt.Fprintln(os.Stderr, tty.FgGreen().Start(),
			"\b-> writing:", tty.FgGreen().End(), name)
	})
	if err != nil {
		printError(err)
		os.Exit(1)
	}
}
