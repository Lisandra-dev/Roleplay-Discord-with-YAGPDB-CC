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
	{{if ge (len .CmdArgs) 1}}
		{{$item = title (index .CmdArgs 0)}}
		{{if ge (len .CmdArgs) 2}}
			{{$bc = index .CmdArgs 1}}
			{{if (ge (len .CmdArgs) 3)}}
				{{$sf = add $sf (index .CmdArgs 2)}}
				{{if (ge (len .CmdArgs) 4)}}
					{{$lc = add $lc (index .CmdArgs 3)}}
					{{if (ge (len .CmdArgs) 5)}}
						{{$cb = add $cb (index .CmdArgs 4)}}
						{{if (ge (len .CmdArgs) 6)}}
							{{$cu = add $cu (index .CmdArgs 5)}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		Merci d'indiquer un nom
	{{end}}
	{{$recipe.Set $item (sdict "Biocomposant" $bc "Substrat Ferreux" $sf "Liquide Cytomorphe" $lc "Cellule Bionotropique" $cb "Composant universel" $cu)}}
	{{$recipe}}
{{else}}
Les arguments ???
{{end}}
{{dbSet 0 "recipe" $recipe}}
