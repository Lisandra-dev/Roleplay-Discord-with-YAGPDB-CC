{{$recipe := sdict}}
{{with (dbGet 0 "recipe")}}
	{{$recipe = sdict .Value}}
{{end}}

{{$soin := sdict}}
{{with ($recipe.Get "Soin")}}
	{{$soin = sdict .}}
{{end}}

{{$bio := sdict}}
{{with ($recipe.Get "arme_bio")}}
	{{$bio = sdict .}}
{{end}}

{{$arme := sdict}}
{{with ($recipe.Get "arme_metal")}}
	{{$arme = sdict .}}
{{end}}

{{$module := sdict}}
{{with ($recipe.Get "module")}}
	{{$module = sdict .}}
{{end}}

{{$implant := sdict}}
{{with ($recipe.Get "implant")}}
	{{$implant = sdict .}}
{{end}}
{{$flag := reFind `-(soin|armebio|arme|module|implant)` .Message.Content}}

{{$bc := 0}}
{{$lc := 0}}
{{$sf := 0}}
{{$cb := 0}}
{{$cu := 0}}
{{$item := ""}}

{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{if ($recipe.Get $item)}}
		{{$i = sdict $soin.Get $item}}
		{{$bc = $i.Get "Biocomposant"}}
		{{$lc = $i.Get "Liquide Cytomorphe"}}
		{{$sf = $i.Get "Substrat Ferreux"}}
		{{$cb = $i.Get "Cellule Bionotropique"}}
		{{$cu = $i.Get "Composant universel"}}
		{{$compo := reFind `-(bc|lc|sf|cb|cu)` .Message.Content}}
		{{if not $compo}}
			Erreur : Composant non reconnu
		{{else}}
			{{if eq $compo "-bc"}}
				{{$bc = (index .CmdArgs 2)}}
			{{else if eq $compo "-lc"}}
				{{$lc = (index .CmdArgs 2)}}
			{{else if eq $compo "-sf"}}
				{{$sf = (index .CmdArgs 2)}}
			{{else if eq $compo "-cb"}}
				{{$cb = (index .CmdArgs 2)}}
			{{else if eq $compo "-cu"}}
				{{$cu = (index .CmdArgs 2)}}
			{{end}}
			{{$recipe.Set $item (sdict "Biocomposant" $bc "Substrat Ferreux" $sf "Liquide Cytomorphe" $lc "Cellule Bionotropique" $cb "Composant universel" $cu)}}
		{{end}}
	{{else}}
		Objet non reconnu.
	{{end}}
{{else}}
Pas d'arguments
{{end}}
{{dbSet 0 "recipe" $recipe}}
