# go-otfdocgen

## Overview

otfdocgen is a generator that will open a OTF (OpenType Font) and search for the individual glyphs, mapping there UTF unicode values to the glyph.  It uses this mapping and a go template file to generate the document. 

## Installation

otfdocgen uses [Go](https://golang.org/doc/install).

Once Go is install, you can clone otfdocgen  with 

```
go get github.com/kevinscardina/go-otfdocgen/...
```

*this might need* `sudo` *depending on your system*

Once cloned, go to the `$GOPATH/github.com/kevinscardina/go-otfdocgen` directory and run the `install.sh` script.

```
./install.sh
```

*this might need* `sudo` *depending on your system*

This will create the executable `otfdocgen` in a `build` directory and copy it into your `/usr/local/bin` directory which should already be in your path.

[video](./otfdocgen.mp4)

## Usage

Once installed you can run the `otfdocgen -help` to get the list of commands.  

```
Usage of otfdocgen:
  -addr string
      address to bind to.
  -end int
      Glyph index to end searching for glyphs at (default 32767)
  -in string
      OTF file to parse (default stdin)
  -out string
      File to put output, include an extension to use that template for generation (default stdout)
  -port string
      port to bind to. (default "8080")
  -prefix string
      Suffix to add to rune name when glyph name does not exist (default u0x) (default "u0x")
  -srv
      Add flag if you would like this to run as a server.
  -start int
      Glyph index to start searching for glyphs at (default -32768) (default -32768)
  -suffix string
      Suffix to add to rune name when glyph name does not exist (default )
  -tmpl string
      Template file for output (default MD)
```
A swift, html, and md, template are included in the application.  If you would like to use a different template, you can use the `-tmpl <filename>` option to use your own template file.  

### Examples

To run as a server so you can go to `addr:port` try:

```
otfdocgen -svr -addr "" -port "8080"
```

then open your browser and put in URL `localhost:8080`.

To generate a Markdown document from a font try:

```
otfdocgen -in OTFIcons.otf -out OTFIcons.md
```

Which will output markdown of the font with inline images to the OTFIcons.md file.  [OTFIcons can be found here](https://www.fontspace.com/adobe/otf-icons).

To generate a html file try:

```
otfdocgen -in OTFIcons.otf -out OTFIcons.html
```

To generate a swift enum class with all the unicode values and handy functions try:

```
otfdocgen -in OTFIcons.otf -out OTFIcons.swift
```

This will generate a OTFIcons.swift file like below, and example of using this file can be found [here](https://github.com/kevinscardina/swift-otficon):

```swift

//
// generated by otfdocgen, a go-lang OTF icon font generator
//

import UIKit
    
/**
 Different Glyphs for runes in OTFIcons 

 [OTFDocGen](https://github.com/kevinscardina/go-otfdocgen)
 */
enum OTFIcons: String, CustomStringConvertible, CaseIterable {
	/// The OTFIcons as a UIFont
	static func font(size:CGFloat) -> UIFont? {
		return UIFont(name: "OTFIcons", size: size)
	}

	 case u0x0020 = "\u{0020}"
	 case u0x0030 = "\u{0030}"
	 case u0x0031 = "\u{0031}"
	 case u0x0032 = "\u{0032}"
	 case u0x0033 = "\u{0033}"
	 case u0x0034 = "\u{0034}"
	 case u0x0035 = "\u{0035}"
	 case u0x0036 = "\u{0036}"
	 case u0x0041 = "\u{0041}"
	 case u0x0042 = "\u{0042}"
	 case u0x0043 = "\u{0043}"
	 case u0x0044 = "\u{0044}"
	 case u0x0045 = "\u{0045}"
	 case u0x0046 = "\u{0046}"
	 case u0x0047 = "\u{0047}"
	 case u0x0048 = "\u{0048}"
	 case u0x0049 = "\u{0049}"
	 case u0x004A = "\u{004A}"
	 case u0x004B = "\u{004B}"
	 case u0x004C = "\u{004C}"
	 case u0x004D = "\u{004D}"
	 case u0x004E = "\u{004E}"
	 case u0x0050 = "\u{0050}"
	 case u0x0051 = "\u{0051}"
	 case u0x0052 = "\u{0052}"
	 case u0x0053 = "\u{0053}"
	 case u0x0055 = "\u{0055}"
	 case u0x0056 = "\u{0056}"
	 case u0x0057 = "\u{0057}"
	 case u0x0061 = "\u{0061}"
	 case u0x0062 = "\u{0062}"
	 case u0x0063 = "\u{0063}"
	 case u0x0064 = "\u{0064}"
	 case u0x0065 = "\u{0065}"
	 case u0x0066 = "\u{0066}"
	 case u0x0067 = "\u{0067}"
	 case u0x0068 = "\u{0068}"
	 case u0x0069 = "\u{0069}"
	 case u0x006A = "\u{006A}"
	 case u0x006B = "\u{006B}"
	 case u0x006C = "\u{006C}"
	 case u0x006D = "\u{006D}"
	 case u0x006E = "\u{006E}"
	 case u0x006F = "\u{006F}"
	 case u0x0070 = "\u{0070}"
	 case u0x007A = "\u{007A}"
	

	/// Gets a String of the icon
	var string: String? {
		return self.rawValue
	}
    
	/// Convenience
	var description: String {
		return string ?? "?"
	}
}
```
