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

**Commande pour remettre : ** `$setchar -all -ID {{$id}} {{(dbGet $id "force").Value}} {{(dbGet $id "i_force").Value}} {{(dbGet $id "endurance").Value}} {{(dbGet $id "i_endu").Value}} {{(dbGet $id "agi").Value}} {{(dbGet $id "i_agi").Value}} {{(dbGet $id "preci").Value}} {{(dbGet $id "i_preci").Value}} {{(dbGet $id "intelligence").Value}} {{(dbGet $id "i_intel").Value}} {{(dbGet $id "karma").Value}}`

{{dbDel $id "force"}}
{{dbDel $id "i_force"}}
{{dbDel $id "agi"}}
{{dbDel $id "i_agi"}}
{{dbDel $id "preci"}}
{{dbDel $id "i_preci"}}
{{dbDel $id "intelligence"}}
{{dbDel $id "i_intel" }}
{{dbDel $id "karma"}}
