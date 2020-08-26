{{$recipe := sdict}}
{{with (dbGet 0 "recipe")}}
	{{$recipe = sdict .Value}}
{{end}}

{{$bc := 0}}
{{$lc := 0}}
{{$sf := 0}}
{{$cb := 0}}
{{$cu := 0}}
{{$item := ""}}
{{$i := sdict}}

{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{if ($recipe.Get $item)}}
	{{$i = sdict $soin.Get $item}}
	{{with $i}}
		{{.}}
	{{end}}
{{end}}
