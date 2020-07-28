{{/* Compétence img*/}}
{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}

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
{{$idict := str $id}}

{{/* Recharge des armes*/}}
{{$rp1 := dbGet $id "r_pistol1"}}
{{$rp2 := dbGet $id "r_pistol2"}}
{{$rf1 := dbGet $id "r_fusil1"}}
{{$rf2 := dbGet $id "r_fusil2"}}
{{$rc := dbGet $id "r_canon"}}
{{$arme := "https://i.imgur.com/WNuPWCv.png"}}

{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "fusil"}}
		{{if not $rf1}}
			{{if gt (toFloat (dbGet $id "fusil").Value (toFloat 11))}}
				{{dbSet $id "r_fusil1" 1}}
				{{ $embed := cembed
					 "author" (sdict "name" $user "icon_url" $arme)
					 "color" 0x6CAB8E
				 "description" "Début du rechargement du fusil...\n Please wait..."}}
				{{$idM := sendMessageRetID nil $embed}}
				{{deleteMessage nil $idM 30}}
			{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Fusil encore chargé...\n Please wait..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
				"description" "Fusil en cours de recharge... \n Please wait..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{else if eq (index .CmdArgs 0) "fusil2"}}
		{{if not $rf2}}
			{{if gt (toFloat (dbGet $id "fusil2").Value (toFloat 11))}}
				{{dbSet $id "r_fusil2" 1}}
				{{ $embed := cembed
					 "author" (sdict "name" $user "icon_url" $arme)
					 "color" 0x6CAB8E
				 "description" "Début du rechargement du deuxième fusil...\n Please wait..."}}
				{{$idM := sendMessageRetID nil $embed}}
				{{deleteMessage nil $idM 30}}
			{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Deuxième fusil encore chargé...\n Please wait..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Deuxième fusil en cours de recharge... \n Please wait..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet"}}
		{{if not $rp1}}
			{{if gt (toFloat (dbGet $id "pistol").Value) (toFloat 7)}}
				{{dbSet $id "r_pistol1" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" "Début du rechargement du pistolet...\n Please wait..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
				"description" "Pistolet encore chargé... \n Please wait..."}}
		  	{{ $idM := sendMessageRetID nil $embed }}
		    {{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Pistolet en cours de recharge... \n Please wait..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet2"}}
		{{if not $rp2}}
			{{if gt (toFloat (dbGet $id "pistol2").Value) (toFloat 7)}}
				{{dbSet $id "r_pistol2" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
				"description" "Début du rechargement du deuxième pistolet...\n Please wait..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
				"description" "Deuxième pistolet encore chargé...\n Please wait... "}}
		  	{{ $idM := sendMessageRetID nil $embed }}
		    {{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Deuxième pistolet en cours de recharge... \n Please wait..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{else if eq (index .CmdArgs 0) "canon"}}
		{{if not $rc}}
			{{if gt (toFloat (dbGet $id "canon").Value) (toFloat 19)}}
				{{dbSet $id "r_canon" 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" "Début du rechargement du canon...\n Please wait..."}}
  	    {{ $idM := sendMessageRetID nil $embed }}
        {{deleteMessage nil $idM 30}}
			{{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $arme)
					"color" 0x6CAB8E
					"description" "Canon encore chargé...\n Please wait..."}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $arme)
				"color" 0x6CAB8E
			"description" "Canon en cours de recharge...\n Please wait..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "attaque"}}
		{{if not (index .CmdArgs 1)}}
			**ERREUR** : N'oubliez pas d'indiquer le nom de l'attaque
		{{else}}
			{{$name := (index .CmdArgs 1)}}
			{{$arg := dbGet $id $name}}
			{{if not $arg}}
				{{dbSet $id "cdatq" $name}}
				{{dbSet $id $name 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Début de la recharge de la competence " $name "...\n Please wait...")
					"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID 735938256038002818 $embed }}
			{{else}}
				{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "La compétence " $name " est toujours en cooldown... \n Please wait...")
				"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{end}}
	{{else if (index .CmdArgs 0) "support"}}
		{{if not (index .CmdArgs 1)}}
			**ERREUR** : N'oubliez pas d'indiquer le nom de l'attaque
		{{else}}
			{{$name := (index .CmdArgs 1)}}
			{{$arg := dbGet $id $name}}
			{{if not $arg}}
				{{dbSet $id $name 1}}
				{{dbSet $id "cdsupp" $name}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Début de la recharge de la competence " $name "...\n Please wait...")
					"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID 735938256038002818 $embed }}
			{{else}}
				{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "La compétence " $name " est toujours en cooldown... \n Please wait...")
				"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{end}}
	{{else}}
		**USAGE** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon|attaque|support)`
	{{end}}
{{else}}
	**USAGE** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon|attaque|support)`
{{end}}
