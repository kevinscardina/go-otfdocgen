package templates

const InputHtml = `
<html>
	<head>
		<title>Upload file</title>
	</head>
	<body>
		<form enctype="multipart/form-data" action="http://localhost:%s/upload" method="post">
			<input type="file" name="uploadfile" />
			<input type="hidden" name="token" value="{{.}}"/>
			<input type="submit" value="swift" name="submit" />
			<input type="submit" value="html" name="submit" />
		</form>
	</body>
</html>
`
