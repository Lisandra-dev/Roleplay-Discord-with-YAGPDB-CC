{{if .CmdArgs}}
	{{$y := (toFloat (index .CmdArgs 0))}}
	{{if eq $y (toFloat 0)}}
		Merci de rentrer un nombre.
	{{else}}
		{{$x := mult $y 1.4 }}
		{{$x := (toInt $x)}}
		{{$x}}
	{{end}}
{{else}}
	**Usage** : `$cc valeur`
{{end}}
