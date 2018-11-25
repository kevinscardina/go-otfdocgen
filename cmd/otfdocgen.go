package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kevinscardina/go-otfdocgen-web/otfdocgenweb/templates"
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

var (
	srv  = flag.Bool("srv", false, "Add flag if you would like this to run as a server.")
	port = flag.String("port", "8080", "port to bind to.")
	addr = flag.String("addr", "", "address to bind to.")
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, tty.FgHiRed().Start(), "ERROR: ", tty.FgRed().Start(), err.Error(), tty.FgRed().End())
}

func main() {
	flag.Parse()
	if *srv {
		http.HandleFunc("/", handlerMain)
		http.HandleFunc("/upload", handlerUpload)
		log.Fatal(http.ListenAndServe(*addr+":"+*port, nil))
		os.Exit(0)
	}
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

	fmt.Fprintln(os.Stderr, tty.FgBlue().Start(), "\b\b   found: ",
		tty.FgBlue().End(), found)

	err = otfDocGen.Write(func(name string) {
		fmt.Fprintln(os.Stderr, tty.FgGreen().Start(),
			"\b-> writing:", tty.FgGreen().End(), name)
	})
	if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprint(w, fmt.Sprintf(templates.InputHtml, *port))
}

func handlerUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method == "POST" {
		tmpl := r.FormValue("submit")

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		inpath := handler.Filename
		outpath := randStringRunes(20) + "." + tmpl
		f, err := os.OpenFile(inpath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(f, file)
		f.Close()
		defer os.Remove(inpath)

		otfdocgen, err := otfdocgen.NewOTFDocGen(inpath, outpath, "", "icon", "", 0, 100000)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = otfdocgen.Read(func(name string) {}, func(count, start, end int) {}, func(r int) {}, func(err error) { fmt.Println(err) })
		if err != nil {
			fmt.Println(err)
			otfdocgen.Destroy()
			return
		}
		err = otfdocgen.Write(func(name string) {})
		if err != nil {
			fmt.Println(err)
			otfdocgen.Destroy()
			return
		}
		defer os.Remove(outpath)
		otfdocgen.Destroy()

		bytes, err := ioutil.ReadFile(outpath)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = w.Write(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
