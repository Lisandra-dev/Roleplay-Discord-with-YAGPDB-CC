{{$recipe := sdict}}
{{with (dbGet 0 "recipe")}}
	{{$recipe = sdict .Value}}
{{end}}

{{$cmd := reFind `-(info|remove)` .Message.Content}}

{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{if (reFind "-bdg" .Message.Content)}}
		{{$item = print "[BDG] " $item}}
	{{end}}
	{{$balle := reFind `(?i)chargeur` .Message.Content}}
	{{if $balle}}
		{{$weap := reFind `(?i)(fusil|canon|pistolet)`}}
		{{$item = print "[CHARGEUR] " (title $weap)}}
	{{end}}
	{{if ($recipe.Get $item)}}
		{{$i := sdict ($recipe.Get $item)}}
		{{if eq $cmd "-info"}}
			{{print "**" $item "** : "}}
			{{range $k, $v := $i}}
				{{print $k " : " $v}}
			{{end}}
		{{else if eq $cmd "-remove"}}
			{{$recipe.Del $item}}
			Objet {{$item}} supprimé
			{{dbSet 0 "recipe" $recipe}}
		{{end}}
	{{else}}
		Objet inconnu
	{{end}}
{{end}}
{{deleteTrigger 1}}
