{{/* img*/}}
{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}
{{$arme := "https://i.imgur.com/WNuPWCv.png"}}

{{/* Get joueur */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr " " (split $name "#")}}
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

{{/* Arme*/}}
{{$recharge := sdict}}
{{with (dbGet $id "recharge")}}
	{{$recharge = sdict .Value}}
{{end}}

{{$rp1 := toInt ($recharge.Get "r_pistol1")}}
{{$rp2 := toInt ($recharge.Get "r_pistol2")}}
{{$rf1 := toInt ($recharge.Get "r_fusil1")}}
{{$rf2 := toInt ($recharge.Get "r_fusil2")}}
{{$rc := toInt ($recharge.Get "r_canon")}}

{{$weapon := sdict}}
{{with (dbGet $id "arme")}}
	{{$weapon = sdict .Value}}
{{end}}


{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "fusil"}}
		{{if not $rf1}}
			{{if gt (toFloat ($weapon.Get "fusil") (toFloat 11))}}
				{{$recharge.Set "r_fusil1" 1}}
				{{ $embed := cembed
					 "author" (sdict "name" $user "icon_url" $arme)
					 "color" 0x6CAB8E
				 "description" "Début du rechargement du fusil..."}}
				{{$idM := sendMessageRetID nil $embed}}
				{{deleteMessage nil $idM 30}}
			{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" (joinStr " " "Il y a encore" ($weapon.Get "fusil") "balles dans le fusil.")}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
				"description" "Fusil en cours de recharge..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "fusil2"}}
		{{if not $rf2}}
			{{if gt (toFloat ($weapon.Get "fusil2") (toFloat 11))}}
			{{$recharge.Set "r_fusil1" 1}}
				{{ $embed := cembed
					 "author" (sdict "name" $user "icon_url" $arme)
					 "color" 0x6CAB8E
				 "description" "Début du rechargement du fusil secondaire..."}}
				{{$idM := sendMessageRetID nil $embed}}
				{{deleteMessage nil $idM 30}}
			{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
				"description" (joinStr " " "Il y a encore" ($weapon.Get "fusil2") "balles dans le fusil secondaire.")}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Fusil secondaire en cours de recharge..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet"}}
		{{if not $rp1}}
			{{if gt (toFloat ($weapon.Get "pistol")) (toFloat 7)}}
				{{$recharge.Set "r_pistol1" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" "Début du rechargement du pistolet..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" (joinStr " " "Il y a encore" ($weapon.Get "pistol") "balles dans le pistolet.")}}
		  	{{ $idM := sendMessageRetID nil $embed }}
		    {{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Pistolet en cours de recharge..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{else if eq (index .CmdArgs 0) "pistolet2"}}
		{{if not $rp2}}
			{{if gt (toFloat ($weapon.Get "pistol2")) (toFloat 7)}}
				{{$recharge.Set "r_pistol2" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
				"description" "Début du rechargement du pistolet secondaire..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" (joinStr " " "Il y a encore" ($weapon.Get "pistol2") "balles dans le pistolet secondaire.")}}
				{{ $idM := sendMessageRetID nil $embed }}
		    {{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Pistolet secondaire en cours de recharge..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "canon"}}
		{{if not $rc}}
			{{if gt (toFloat ($weapon.Get "canon")) (toFloat 19)}}
				{{$recharge.Set "r_canon" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" "Début du rechargement du canon..."}}
  	    {{ $idM := sendMessageRetID nil $embed }}
        {{deleteMessage nil $idM 30}}
			{{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" (joinStr " " "Il y a encore" ($weapon.Get "canon") "balles dans le canon.")}}
					{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Canon en cours de recharge..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{else}}
		**Usage** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`
	{{end}}
{{else}}
**Usage** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`
{{end}}
{{dbSet $id "recharge" $recharge}}
