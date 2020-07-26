{{$img := "https://i.imgur.com/DPuFF4i.png"}}



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
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}
{{$id = str $id}}

{{/* Recharge */}}
{{$rp1 := (dbGet $id "r_pistol1").Value}}
{{$rp2 := (dbGet $id "r_pistol2").Value}}
{{$rf1 := (dbGet $id "r_fusil1").Value}}
{{$rf2 := (dbGet $id "r_fusil2").Value}}
{{$rc := (dbGet $id "r_canon").Value}}

{{$mp1 := ""}}
{{$mp2 := ""}}
{{$mf1 := ""}}
{{$mf2 := ""}}
{{$mc := ""}}

{{/* PA count */}}
{{$desc := ""}}
{{$bool := false}}
{{range $i, $j := $groupe}}
	{{- if eq $id $i}}
		{{- $bool = true}}
	{{- end -}}
{{end}}
{{if eq $bool false}}
		{{$groupe.Set $id 3}}
		{{$desc = joinStr " " "Il vous reste 3 PA"}}

		{{if $rp1}}
			{{if lt $rp1 2}}
				{{$x := dbIncr $id $rp1 1}}
			{{else if eq $rp1 2}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
					"description" (joinStr "" "Pistolet de " $.User.Mention " rechargé.")
					"color" 0x6CAB8E}}
				{{sendMessage nil $embed}}
				{{dbDel $rp1}}
			{{end}}
		{{end}}
		{{if $rp2}}
			{{if lt $rp2 2}}
				{{$x := dbIncr $id $rp2 1}}
			{{else if eq $rp2 2}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
					"description" (joinStr "" "Deuxième pistolet de " $.User.Mention " rechargé.")
					"color" 0x6CAB8E}}
				{{sendMessage nil $embed}}
				{{dbDel $rp2}}
			{{end}}
		{{end}}
		{{if $rf1}}
			{{if lt $rf1 4}}
				{{$x := dbIncr $id $rf1 1}}
			{{else if eq $rf1 4}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
					"description" (joinStr "" "Fusil de " $.User.Mention " rechargé.")
					"color" 0x6CAB8E}}
				{{sendMessage nil $embed}}
				{{dbDel $rp2}}
			{{end}}
		{{end}}
		{{if $rf2}}
			{{if lt $rf2 4}}
				{{$x := dbIncr $id $rf2 1}}
			{{else if eq $rf2 4}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" "https://i.imgur.com/YeIsRmw.png")
					"description" (joinStr "" "Deuxième fusil de " $.User.Mention " rechargé.")
					"color" 0x6CAB8E}}
				{{sendMessage nil $embed}}
				{{dbDel $rf2}}
			{{end}}
		{{end}}

		{{if $rc}}
		{{end}}
{{else}}
	{{$j := $groupe.Get $id}}
	{{$j = sub $j 1}}
	{{if gt $j 0}}
		{{$desc = joinStr " " "Il vous reste" $j "PA"}}
	{{else if le $j 0}}
		{{$j = 0}}
		{{$desc = joinStr " " "Vous avez consommé tous vos PA, merci d'attendre le prochain tour"}}
	{{end}}
	{{$groupe.Set $id $j}}
{{end}}
{{$desc}}
{{dbSet .Server.ID "groupe" $groupe}}
