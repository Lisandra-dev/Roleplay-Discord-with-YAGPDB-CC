{{/* To adapt, change the text for the composant.
	Name used in :
	-> Inventory dict : [C] items
	-> Compo dict : diminutif without majuscule
	-> recipe : Complete name with majuscule
	- You can add more component with adding more loop
	- You can change the option with changing the ReFind. */}}


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

{{if .CmdArgs}}
	{{if ge (len .CmdArgs) 1}}
		{{$item = title (index .CmdArgs 0)}}
		{{if ge (len .CmdArgs) 2}}
			{{$compo := reFindAll `-(bc|sf|lc|cb)` .Message.Content}}
			{{if eq "-bc" (index $compo 0)}}
				{{$bc = index .CmdArgs 2}}
			{{else if eq "-sf" (index $compo 0)}}
				{{$sf = index .CmdArgs 2}}
			{{else if eq "-lc" (index $compo 0)}}
				{{$lc = index .CmdArgs 2}}
			{{else if eq "-cb" (index $compo 0)}}
				{{$cb = index .CmdArgs 2}}
			{{end}}
			{{if and (ge (len .CmdArgs) 4) (ge (len $compo) 2)}}
				{{if eq "-bc" (index $compo 1)}}
					{{$bc = index .CmdArgs 4}}
				{{else if eq "-sf" (index $compo 1)}}
					{{$sf = index .CmdArgs 4}}
				{{else if eq "-lc" (index $compo 1)}}
					{{$lc = index .CmdArgs 4}}
				{{else if eq "-cb" (index $compo 1)}}
					{{$cb = index .CmdArgs 4}}
				{{end}}
				{{if and (ge (len .CmdArgs) 6) (ge (len $compo) 3) }}
					{{if eq "-bc" (index $compo 2)}}
						{{$bc = index .CmdArgs 6}}
					{{else if eq "-sf" (index $compo 2)}}
						{{$sf = index .CmdArgs 6}}
					{{else if eq "-lc" (index $compo 2)}}
						{{$lc = index .CmdArgs 6}}
					{{else if eq "-cb" (index $compo 2)}}
						{{$cb = index .CmdArgs 6}}
					{{end}}
					{{if and (ge (len .CmdArgs) 8) (ge (len $compo) 4)}}
						{{if eq "-bc" (index $compo 3)}}
							{{$bc = index .CmdArgs 8}}
						{{else if eq "-sf" (index $compo 3)}}
							{{$sf = index .CmdArgs 8}}
						{{else if eq "-lc" (index $compo 3)}}
							{{$lc = index .CmdArgs 8}}
						{{else if eq "-cb" (index $compo 3)}}
							{{$cb = index .CmdArgs 8}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{else}}
			Merci d'indiquer au moins un composant.
		{{end}}
	{{else}}
		Merci d'indiquer un nom
	{{end}}
	{{if (reFind "-bdg" .Message.Content)}}
		{{$item = print "[BDG] " $item}}
	{{end}}
	{{$balle := reFind `(?i)chargeur` .Message.Content}}
	{{if $balle}}
		{{$weap := reFind `(?i)(pistolet|fusil|canon)` $item}}
		{{$item = print "[CHARGEUR] " (title $weap)}}
	{{end}}
	{{$recipe.Set $item (sdict "Biocomposant" $bc "Substrat Ferreux" $sf "Liquide Cytomorphe" $lc "Cellule Bionotropique" $cb)}}
	{{$item}}
	{{range $k, $v := ($recipe.Get $item)}}
		{{print $k " : " $v}}
	{{end}}
{{else}}
**Usage** : `$recipe "nom" -(bc|lc|sf|cb) (-(bc|lc|sf|cb)) (-(bc|lc|sf|cb)) (-(bc|lc|sf|cb)) (-bdg)`
{{end}}
{{dbSet 0 "recipe" $recipe}}
{{deleteTrigger 1}}
