{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .User.ID }}
{{$id := .User.ID}}
{{if .CmdArgs}}
	{{$id = toInt (index .CmdArgs 0)}}
	{{$user = toInt (index .CmdArgs 0)}}
{{end}}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
{{end}}

{{$stats := sdict}}
{{with (dbGet $id "stats")}}
	{{$stats = sdict .Value}}
{{end}}

{{$force:= $stats.Get "force"}}
{{$endurance := $stats.Get "endurance"}}
{{$agi:=$stats.Get "agi"}}
{{$preci:=$stats.Get "preci"}}
{{$intel:=$stats.Get "intelligence"}}
{{$karma:=$stats.Get "karma"}}

{{$iforce:=$stats.Get "i_force"}}
{{$iendu:=$stats.Get "i_endu"}}
{{$iagi:=$stats.Get "i_agi"}}
{{$ipreci:=$stats.Get "i_preci"}}
{{$iintel:=$stats.Get "i_intel"}}

**Commande pour remettre : ** `$setchar -all -ID {{$id}} {{$force}} {{$iforce}} {{$endurance}} {{$iendu}} {{$agi}} {{$iagi}} {{$preci}} {{$ipreci}} {{$intel}} {{$iintel}} {{$karma}}`

{{dbDel $id "stats"}}
