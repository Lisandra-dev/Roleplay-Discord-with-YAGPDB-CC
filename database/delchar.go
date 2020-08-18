{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .User.Username }}
{{$id := .User.ID}}
{{if .CmdArgs}}
	{{$id = toInt (index .CmdArgs 0)}}
	{{$user = (userArg $id)}}
{{end}}
{{if $name}}
	{{$user = $name}}
	{{$id = .User.ID}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = (add $id .)}}
	{{- end}}
{{end}}

{{$stats := sdict}}
{{with (dbGet $id "stats")}}
	{{$stats = sdict .Value}}
{{end}}

{{if not $stats}}
	Le personnage n'était pas dans la base de données !
{{else}}
	{{dbDel $id "stats"}}
		{{$embed := cembed
		"description" (joinStr "" (title $user) " aka " $id " a bien été supprimé de la base de données")}}
	{{sendMessage nil $embed}}
{{end}}
