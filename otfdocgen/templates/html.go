package templates

const (
	HTML = `
<html>
	<head>
		<title>Glyphs in {{.Name}}</title>
		<style>
			body {
					background-color: white;
					font-family: verdana;
					font-size: 12pt;
			}
		</style>
	</head>
	<body>
		<table width="400pt">
			<tr><th>Name</th><th>Hex String</th><th>Image</th></tr>
			{{range .Glyphs}}{{println "<tr><td>" .Name "</td><td>" .HexString "</td><td><img src=\"data:image/png;base64," .ImageBase64 "\" /></td></tr>"}}{{end}}
		</table>
	</body>
</html>
`
)
