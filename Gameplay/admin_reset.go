{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := ""}}
{{if .CmdArgs}}
	{{if not $name}}
	{{$id = (userArg (index .CmdArgs 0)).ID}}
	{{$user = (getMember $id).Nick}}
	{{else if $name}}
		{{$user = $name}}
		{{$idperso := (toRune (lower $name))}}
		{{$id = ""}}
		{{range $idperso}}
			{{- $id = (print $id .)}}
		{{- end}}
		{{$id = (toInt $id)}}
	{{end}}
{{end}}
{{$user = title $user}}
{{$idict := str $id}}

{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{dbDel $id "fusil"}}
{{dbDel $id "fusil2"}}
{{dbDel $id "pistol"}}
{{dbDel $id "pistol2"}}
{{dbDel $id "canon"}}
{{$atq := (dbGet $id "cdatq").Value}}
{{$supp := (dbGet $id "cdsupp").Value}}
{{dbDel $id $atq}}
{{dbDel $id $supp}}
{{dbDel $id "cdatq"}}
{{dbDel $id "cdsupp"}}
{{dbDel $id "r_pistol1"}}
{{dbDel $id "r_pistol2"}}
{{dbDel $id "r_fusil1"}}
{{dbDel $id "r_fusil2"}}
{{dbDel $id "r_canon"}}

{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{$bool := "false"}}
{{range $i, $j := $groupe}}
	{{- if eq $idict $i}}
		{{- $bool = "true"}}
	{{- end -}}
{{end}}

{{if eq $bool "true"}}
	{{$groupe.Del $idict}}
{{end}}
{{dbSet .Server.ID "groupe" $groupe}}

Toutes les variables d'armes et de PA de {{$user}} ont été supprimé de la DB !
