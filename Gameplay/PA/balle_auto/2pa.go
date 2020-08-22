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

{{/* Recharge des armes*/}}
{{$rp1 := (toInt (dbGet $id "r_pistol1").Value)}}
{{$rp2 := (toInt (dbGet $id "r_pistol2").Value)}}
{{$rf1 := (toInt (dbGet $id "r_fusil1").Value)}}
{{$rf2 := (toInt (dbGet $id "r_fusil2").Value)}}
{{$rc := (toInt (dbGet $id "r_canon").Value)}}

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

{{/* PA count */}}
{{if eq $bool false}}
	{{$groupe.Set $idict 2}}
	{{$desc = joinStr " " "Il vous reste 2 PA"}}
	{{if $rp1}}
		{{if lt $rp1 2}}
			{{$x := dbIncr $id "r_pistol1" 2}}
		{{else if eq $rp1 2}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Pistolet de " $.User.Mention " rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{dbDel $id "r_pistol1"}}
		{{end}}
	{{end}}
	{{if $rp2}}
		{{if lt $rp2 2}}
			{{$x := dbIncr $id "r_pistol2" 2}}
		{{else if eq $rp2 2}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Deuxième pistolet de " $.User.Mention " rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{dbDel "r_pistol2"}}
		{{end}}
	{{end}}
	{{if $rf1}}
		{{if lt $rf1 4}}
			{{$x := dbIncr $id "r_fusil1" 2}}
		{{else if eq $rf1 4}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Fusil de " $.User.Mention " rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{dbDel $id "r_fusil1"}}
		{{end}}
	{{end}}
	{{if $rf2}}
		{{if lt $rf2 4}}
			{{$x := dbIncr $id "r_fusil2" 2}}
		{{else if eq $rf2 4}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Deuxième fusil de " $.User.Mention " rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{dbDel $id "r_fusil2"}}
		{{end}}
	{{end}}
	{{if $rc}}
		{{if lt $rc 8}}
			{{$x := dbIncr $id "r_canon" 2}}
		{{else if eq $rc 8}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Deuxième fusil de " $.User.Mention " rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{dbDel $id "r_canon"}}
		{{end}}
	{{end}}
	{{if $nameatq}}
		{{if lt $atq 8}}
			{{$x:= dbIncr $id $atq 2}}
		{{else if eq $atq 4}}
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
			{{$x := dbIncr $id $supp 2}}
		{{else if eq $supp 8}}
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
	{{if eq $j 1}}
		{{$j =1}}
		{{$desc = "Vous n'avez pas assez de PA"}}
	{{else}}
		{{$j = sub $j 2}}
		{{if gt $j 0}}
			{{$desc = joinStr " " "Il vous reste" $j "PA"}}
			{{if $rp1}}
				{{if lt $rp1 2}}
					{{$x := dbIncr $id "r_pistol1" 1}}
				{{else if eq $rp1 2}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Pistolet de " $.User.Mention " rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{dbDel $id "r_pistol1"}}
					{{dbDel $id "pistol"}}
				{{end}}
			{{end}}
			{{if $rp2}}
				{{if lt $rp2 2}}
					{{$x := dbIncr $id "r_pistol2" 2}}
				{{else if eq $rp2 2}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Deuxième pistolet de " $.User.Mention " rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{dbDel "r_pistol2"}}
					{{dbDel $id "pistol2"}}
				{{end}}
			{{end}}
			{{if $rf1}}
				{{if lt $rf1 4}}
					{{$x := dbIncr $id "r_fusil1" 2}}
				{{else if eq $rf1 4}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Fusil de " $.User.Mention " rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{dbDel $id "r_fusil1"}}
					{{dbDel $id "fusil"}}
				{{end}}
			{{end}}
			{{if $rf2}}
				{{if lt $rf2 4}}
					{{$x := dbIncr $id "r_fusil2" 2}}
				{{else if eq $rf2 4}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Deuxième fusil de " $.User.Mention " rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{dbDel $id "r_fusil2"}}
					{{dbDel $id "fusil2"}}
				{{end}}
			{{end}}
			{{if $rc}}
				{{if lt (toInt $rc) 8}}
					{{$x := dbIncr $id "r_canon" 2}}
				{{else if eq (toInt $rc) 8}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Canon de " $.User.Mention " rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{dbDel $id "r_canon"}}
					{{dbDel $id "canon"}}
				{{end}}
			{{end}}
			{{if $nameatq}}
				{{if lt $atq 8}}
					{{$x:= dbIncr $id $nameatq 2}}
				{{else if eq $atq 4}}
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
					{{$x := dbIncr $id $namesupp 2}}
				{{else if eq $supp 8}}
					{{dbDel $id "cdsupp"}}
					{{ $embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/9iRdtbM.png")
						"description" (joinStr "" "Compétence " $namesupp " de nouveau utilisable")
						"color" 0xDFAA58}}
					{{ $idM := sendMessageRetID $chan $embed }}
					{{dbDel $id $supp}}
				{{end}}
			{{end}}
		{{else if le $j 0}}
			{{$j = 0}}
			{{$desc = joinStr " " "Il vous reste 0 PA."}}
		{{end}}
		{{$groupe.Set $idict $j}}
	{{end}}
{{end}}
{{$embed := cembed
	"author" (sdict "name" $user "icon_url" "https://i.imgur.com/VvOhTON.png")
	"description" $desc
	"color" $col}}
	{{$idPA := sendMessageRetID nil $embed}}
	{{deleteMessage nil $idPA 30}}
{{dbSet .Server.ID "groupe" $groupe}}
