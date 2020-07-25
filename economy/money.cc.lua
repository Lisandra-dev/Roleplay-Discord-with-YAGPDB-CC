{{/* Databases */}}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}
{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}
{{$symbol := ""}}
{{if $serverEco.Get "symbol"}}
	{{$symbol = $serverEco.Get "symbol"}}
{{end}}

{{/* Balance */}}
{{$bal := (toInt ($userEco.Get "balance"))}}
{{sendMessage nil (cembed "author" (sdict "name" (title $user) "icon_url" "https://i.imgur.com/ATSj8fe.png") "description" (print (str $bal) " " $symbol ) "color" 0x8CBAEF)}}
{{deleteTrigger 1}}
