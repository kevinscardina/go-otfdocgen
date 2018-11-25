package templates

const InputHtml = `
<html>
	<head>
		<title>OTFDocGen Web Interface</title>
	</head>
	<body>

<pre>
This is a simple Web Interface to the otfdocgen commandline tool.

Choose the OTF file you would like to create a "doc" for and click on the type of doc to create.
</pre>
<hr />
		<form enctype="multipart/form-data" action="http://localhost:%s/upload" method="post">
			<input type="file" name="uploadfile" />
			<input type="submit" value="md" name="submit" />
			<input type="submit" value="html" name="submit" />
			<input type="submit" value="swift" name="submit" />
		</form>
	</body>
</html>
`
