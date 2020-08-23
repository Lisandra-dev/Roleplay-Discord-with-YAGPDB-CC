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

{{/* Arme */}}
{{$arme := sdict}}
{{with (dbGet $id "arme")}}
	{{$arme = sdict .Value}}
{{end}}

{{/* recharge des armes*/}}
{{$recharge := sdict}}
{{with (dbGet $id "recharge")}}
	{{$recharge = sdict .Value}}
{{end}}

{{$rp1 := toInt ($recharge.Get "r_pistol1")}}
{{$rp2 := toInt ($recharge.Get "r_pistol2")}}
{{$rf1 := toInt ($recharge.Get "r_fusil1")}}
{{$rf2 := toInt ($recharge.Get "r_fusil2")}}
{{$rc := toInt ($recharge.Get "r_canon")}}

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
	{{$groupe.Set $idict 3}}
	{{$desc = joinStr " " "3 PA RESTANTS."}}
	{{if $rp1}}
		{{if lt $rp1 2}}
			{{$recharge.Set "r_pistol1" (add $rp1 1)}}
		{{else if ge $rp1 2}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Pistolet rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{$recharge.Del "r_pistol1"}}
		{{end}}
	{{end}}
	{{if $rp2}}
		{{if lt $rp2 2}}
			{{$recharge.Set "r_pistol2" (add $rp2 1)}}
		{{else if ge $rp2 2}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Pistolet secondaire rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{$recharge.Del "r_pistol2"}}
		{{end}}
	{{end}}
	{{if $rf1}}
		{{if lt $rf1 4}}
			{{$recharge.Set "r_fusil1" (add $rf1 1)}}
		{{else if ge $rf1 4}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Fusil rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{$recharge.Del "r_fusil1"}}
		{{end}}
	{{end}}
	{{if $rf2}}
		{{if ge $rf2 4}}
			{{$recharge.Set "r_fusil2" (add $rf2 1)}}
		{{else if eq $rf2 4}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Fusil secondaire rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{$recharge.Del "r_fusil2"}}
		{{end}}
	{{end}}
	{{if $rc}}
		{{if lt $rc 8}}
			{{$recharge.Set "r_canon" (add $rc 1)}}
		{{else if ge $rc 8}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
				"description" (joinStr "" "Canon rechargé.")
				"color" 0x6CAB8E}}
			{{sendMessage nil $embed}}
			{{$recharge.Del "r_canon"}}
		{{end}}
	{{end}}
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
	{{if eq $j 1}}
		{{$groupe.Set $idict 0}}
		{{$desc = "DERNIER PA UTILISÉ"}}
	{{else}}
		{{$j = sub $j 1}}
		{{if gt $j 0}}
			{{$desc = joinStr " " $j "PA RESTANTS"}}
			{{if $rp1}}
				{{if lt $rp1 2}}
					{{$recharge.Set "r_pistol1" (add $rp1 1)}}
				{{else if ge $rp1 2}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Pistolet rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{$recharge.Del "r_pistol1"}}
					{{$arme.Del "pistol"}}
				{{end}}
			{{end}}
			{{if $rp2}}
				{{if lt $rp2 2}}
				{{$recharge.Set "r_pistol2" (add $rp2 1)}}
				{{else if ge $rp2 2}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Pistolet secondaire rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{$recharge.Del "r_pistol2"}}
					{{$arme.Del "pistol"}}
				{{end}}
			{{end}}
			{{if $rf1}}
				{{if ge $rf1 4}}
				{{$recharge.Set "r_fusil1" (add $rf1 1)}}
				{{else if eq $rf1 4}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Fusil rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{$recharge.Del "r_fusil1"}}
					{{$arme.Del "fusil"}}
				{{end}}
			{{end}}
			{{if $rf2}}
				{{if ge $rf2 4}}
				{{$recharge.Set "r_fusil2" (add $rf2 1)}}
				{{else if eq $rf2 4}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" "Fusil secondaire rechargé."
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{$recharge.Del "r_fusil2"}}
					{{$arme.Del "fusil2"}}
				{{end}}
			{{end}}
			{{if $rc}}
				{{if ge (toInt $rc) 8}}
				{{$recharge.Set "r_canon" (add $rc 1)}}
				{{else if eq (toInt $rc) 8}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
						"description" (joinStr "" "Canon rechargé.")
						"color" 0x6CAB8E}}
					{{sendMessage nil $embed}}
					{{$recharge.Del "r_canon"}}
					{{$arme.Del "canon"}}
				{{end}}
			{{end}}
			{{if $nameatq}}
				{{if ge $atq 8}}
					{{$x:= dbIncr $id $nameatq 1}}
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
				{{if ge $supp 8}}
					{{$x := dbIncr $id $namesupp 1}}
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
			{{$desc = joinStr " " "PA INSUFFISANTS POUR RÉALISER L'ACTION."}}
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
{{dbSet $id "recharge" $recharge}}
{{dbSet $id "arme" $arme}}
