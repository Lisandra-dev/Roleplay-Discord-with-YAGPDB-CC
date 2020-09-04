
{{/* To adapt, change the text for the composant.
	Name used in :
	-> Inventory dict : [C] items
	-> Compo dict : diminutif without majuscule
	-> recipe : Complete name with majuscule
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
	{{$item := title (index .CmdArgs 0)}}
	{{if (reFind "-bdg" .Message.Content)}}
		{{$item = print "[BDG] " $item}}
	{{end}}
	{{$balle := reFind `(?i)chargeur` .Message.Content}}
	{{if $balle}}
		{{$weap := reFind `(?i)(fusil|canon|pistolet)` $item}}
		{{$item = print "[CHARGEUR] " (title $weap)}}
	{{end}}
	{{if ($recipe.Get $item)}}
		{{$i := sdict ($recipe.Get $item)}}
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
			{{with ($recipe.Get $item)}}
				{{print .}}
			{{end}}
		{{end}}
	{{else}}
		Objet non reconnu.
	{{end}}
{{else}}
Pas d'arguments
{{end}}
{{dbSet 0 "recipe" $recipe}}
{{deleteTrigger 1}}
