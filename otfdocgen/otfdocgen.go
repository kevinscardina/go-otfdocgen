package otfdocgen

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kevinscardina/go-otfdocgen/otfdocgen/templates"
	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/vector"
)

type OTFDocGen struct {
	in         *os.File
	out        *os.File
	tmpl       *template.Template
	start      int
	end        int
	namePrefix string
	nameSuffix string
	Name       string
	TypeName   string
	Glyphs     []struct {
		Name        string
		HexString   string
		ImageBase64 string
	}
}

func NewOTFDocGen(inpath, outpath, tmplpath, namePrefix, nameSuffix string,
	start, end int) (*OTFDocGen, error) {

	in := os.Stdin
	name := "Stdin"
	typeName := name
	var err error
	if len(inpath) > 0 {
		in, err = os.Open(inpath)
		if err != nil {
			return nil, fmt.Errorf("unable to open file %v - (%v)", inpath, err)
		}
		basename := filepath.Base(inpath)
		name = strings.TrimSuffix(basename, filepath.Ext(basename))
		if i := strings.LastIndex(name, "-"); i > 0 {
			typeName = name[0:i]
		} else {
			typeName = name
		}
	}

	out := os.Stdout
	if len(outpath) > 0 {
		_, err = os.Stat(outpath)
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("output file %s already exists", outpath)
		}
		out, err = os.Create(outpath)
		if err != nil {
			return nil, fmt.Errorf("unable to create file %v - (%v)", outpath, err)
		}
	}

	toParse := templates.MD
	if len(tmplpath) > 0 {
		tmplFile, err := os.Open(tmplpath)
		if err != nil {
			return nil, err
		}
		defer tmplFile.Close()
		tmplBytes, err := ioutil.ReadAll(tmplFile)
		if err != nil {
			return nil, err
		}
		toParse = string(tmplBytes)
	} else if strings.HasSuffix(outpath, ".html") {
		toParse = templates.HTML
	} else if strings.HasSuffix(outpath, ".swift") {
		toParse = templates.Swift
	}

	tmpl, err := template.New("tmpl").Parse(toParse)
	if err != nil {
		return nil, err
	}

	if start > end {
		return nil, fmt.Errorf("-start must be less then or equal to -end (%v > %v)", start, end)
	}

	return &OTFDocGen{
		in:         in,
		out:        out,
		tmpl:       tmpl,
		start:      start,
		end:        end,
		namePrefix: namePrefix,
		nameSuffix: nameSuffix,
		Name:       name,
		TypeName:   typeName,
		Glyphs: []struct {
			Name        string
			HexString   string
			ImageBase64 string
		}{},
	}, nil
}

func (otf *OTFDocGen) Destroy() {
	otf.in.Close()
	otf.out.Close()
}

func (otf *OTFDocGen) Read(readingFunc func(name string),
	statsFunc func(count, start, end int),
	progressFunc func(r int), errorFunc func(err error)) (int, error) {

	readingFunc(otf.in.Name())
	fnt, err := sfnt.ParseReaderAt(otf.in)
	if err != nil {
		return 0, err
	}

	numGlyphs := fnt.NumGlyphs()
	statsFunc(numGlyphs, otf.start, otf.end)
	buffer := &sfnt.Buffer{}

	numGlyphsFound := 0
	for r := otf.start; r < otf.end && numGlyphsFound < numGlyphs; r++ {
		if r%1001 == 0 {
			progressFunc(r)
		}

		index, err := fnt.GlyphIndex(buffer, rune(r))
		if err != nil {
			return 0, fmt.Errorf("unable to parse glyph index %v - (%v)", r, err)
		}
		if index != 0 {
			name, err := fnt.GlyphName(buffer, index)
			if err != nil {
				return 0, err
			}
			if len(name) == 0 {
				name = fmt.Sprintf("%s%0.4X%s", otf.namePrefix, r, otf.nameSuffix)
			}

			ppem := fixed.I(24)
			segments, err := fnt.LoadGlyph(buffer, index, ppem, nil)
			if err != nil {
				errorFunc(err)
				continue
			}

			fontBounds, err := fnt.Bounds(buffer, ppem, font.HintingFull)
			if err != nil {
				errorFunc(err)
				continue
			}
			w := fontBounds.Max.X.Ceil()
			h := ppem.Ceil() + 4
			raster := vector.NewRasterizer(w, h)
			raster.DrawOp = draw.Src
			for _, seg := range segments {
				// The divisions by 64 below is because the seg.Args values have type
				// fixed.Int26_6, a 26.6 fixed point number, and 1<<6 == 64.
				switch seg.Op {
				case sfnt.SegmentOpMoveTo:
					raster.MoveTo(
						float32(seg.Args[0].X)/64,
						float32(h-4)+float32(seg.Args[0].Y)/64,
					)
				case sfnt.SegmentOpLineTo:
					raster.LineTo(
						float32(seg.Args[0].X)/64,
						float32(h-4)+float32(seg.Args[0].Y)/64,
					)
				case sfnt.SegmentOpQuadTo:
					raster.QuadTo(
						float32(seg.Args[0].X)/64,
						float32(h-4)+float32(seg.Args[0].Y)/64,
						float32(seg.Args[1].X)/64,
						float32(h-4)+float32(seg.Args[1].Y)/64,
					)
				case sfnt.SegmentOpCubeTo:
					raster.CubeTo(
						float32(seg.Args[0].X)/64,
						float32(h-4)+float32(seg.Args[0].Y)/64,
						float32(seg.Args[1].X)/64,
						float32(h-4)+float32(seg.Args[1].Y)/64,
						float32(seg.Args[2].X)/64,
						float32(h-4)+float32(seg.Args[2].Y)/64,
					)
				}
			}

			dstMask := image.NewAlpha(raster.Bounds())
			raster.Draw(dstMask, dstMask.Bounds(), image.Opaque, image.ZP)

			dst := image.NewRGBA(image.Rect(0, 0, w+4, h+4))
			draw.Draw(dst, dst.Bounds(), image.White, image.ZP, draw.Over)

			draw.DrawMask(dst, dst.Bounds(), image.Black, image.ZP, dstMask, image.Point{X: -2, Y: -2}, draw.Over)

			buffer := bytes.NewBuffer([]byte{})
			if err = png.Encode(buffer, dst); err != nil {
				return 0, err
			}

			otf.Glyphs = append(otf.Glyphs, struct {
				Name        string
				HexString   string
				ImageBase64 string
			}{
				Name:        name,
				HexString:   fmt.Sprintf("%0.4X", r),
				ImageBase64: base64.StdEncoding.EncodeToString(buffer.Bytes()),
			})
			numGlyphsFound++
		}
	}
	return numGlyphsFound, nil
}

func (otf *OTFDocGen) Write(writingFunc func(name string)) error {

	writingFunc(otf.out.Name())
	return otf.tmpl.Execute(otf.out, otf)
}
