{{$desc := ""}}
{{$col := 16777215}}
{{$p := 0}}
{{$r := .Member.Roles}}
{{range .Guild.Roles}}
	{{if and (in $r .ID) (.Color) (lt $p .Position)}}
	{{$p = .Position}}
	{{$col = .Color}}
	{{end}}
{{end}}

{{$chan := "735938256038002818"}}

{{/* Groupe dictionnaire */}}

{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Get joueur */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id:= .User.ID}}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}
{{$idict := str $id}}

{{/* Compétence */}}
{{$nameatq := (dbGet $id "cdatq").Value}}
{{$namesupp := (dbGet $id "cdsupp").Value}}
{{$atq := (toInt (dbGet $id $nameatq).Value)}}
{{$supp := (toInt (dbGet $id $namesupp).Value)}}

{{/* PA ID search */}}
{{$bool := false}}
{{range $i, $j := $groupe}}
	{{- if eq $idict $i}}
		{{- $bool = true}}
	{{- end -}}
{{end}}

{{if eq $bool false}}
	{{$groupe.Set $idict 3}}
	{{$desc = joinStr " " "3 PA RESTANTS."}}
	{{if $nameatq}}
		{{if lt $atq 8}}
			{{$x:= dbIncr $id $atq 1}}
		{{else if ge $atq 4}}
			{{dbDel $id "cdatq"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/zNofnyh.png")
				"description" (joinStr "" "Compétence " $nameatq " de nouveau utilisable")
				"color" 0xDFAA58}}
			{{ $idM := sendMessageRetID $chan $embed }}
			{{dbDel $id $atq}}
		{{end}}
	{{end}}

	{{if $namesupp}}
		{{if lt $supp 8}}
			{{$x := dbIncr $id $supp 1}}
		{{else if ge $supp 8}}
			{{dbDel $id "cdsupp"}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/9iRdtbM.png")
				"description" (joinStr "" "Compétence " $namesupp " de nouveau utilisable")
				"color" 0xDFAA58}}
			{{ $idM := sendMessageRetID $chan $embed }}
			{{dbDel $id $supp}}
		{{end}}
	{{end}}

	{{else}}
		{{$j := $groupe.Get $idict}}
		{{$j = sub $j 1}}
		
			{{if ge $j 0}}
				{{$desc = joinStr " " $j "PA RESTANT(S)."}}
				{{if $nameatq}}
				{{if lt $atq 8}}
					{{$x:= dbIncr $id $nameatq 1}}
				{{else if ge $atq 4}}
					{{dbDel $id "cdatq"}}
					{{ $embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/zNofnyh.png")
						"description" (joinStr "" "Compétence " $nameatq " de nouveau utilisable")
						"color" 0xDFAA58}}
					{{ $idM := sendMessageRetID $chan $embed }}
					{{dbDel $id $atq}}
				{{end}}
			{{end}}
			{{if $namesupp}}
				{{if lt $supp 8}}
					{{$x := dbIncr $id $namesupp 1}}
				{{else if ge $supp 8}}
					{{dbDel $id "cdsupp"}}
					{{ $embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/9iRdtbM.png")
						"description" (joinStr "" "Compétence " $namesupp " de nouveau utilisable")
						"color" 0xDFAA58}}
					{{ $idM := sendMessageRetID $chan $embed }}
					{{dbDel $id $supp}}
				{{end}}
			{{end}}
			{{else if lt $j 0}}
			{{$j = 0}}
			{{$desc = joinStr " " "PA INSUFFISANTS POUR RÉALISER L'ACTION."}}
		{{end}}
		{{$groupe.Set $idict $j}}
	{{end}}

{{$embed := cembed
	"author" (sdict "name" $user "icon_url" "https://i.imgur.com/VvOhTON.png")
	"description" $desc
	"color" $col}}
	{{$idPA := sendMessageRetID nil $embed}}
	{{deleteMessage nil $idPA 30}}
{{dbSet .Server.ID "groupe" $groupe}}
