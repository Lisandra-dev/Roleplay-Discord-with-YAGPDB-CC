{$name := reFind `(\#\S*)` .Message.Content}}
{{$support:=reFindAllSubmatches `^\!(malus|soutien)` .Message.Content}}
{{$soin:= reFindAllSubmatches `^\!soin` .Message.Content}}
{{$poison:= reFindAllSubmatches `^\!poison` .Message.Content}}
{{$soinimg := "https://i.imgur.com/fWnxoZ4.png"}}
{{$supportimg:="https://i.imgur.com/fuHvIUn.png"}}
{{$poisonimg:="https://i.imgur.com/CmgNFYu.png"}}
{{$s := "+"}}
{{$ec := "Bonus de -"}}
{{$soincol:=0xb66dc1}}
{{$poisoncol:=0x370E6A}}
{{$supportcol:=0x2B7697}}
{{$id:= ""}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{if $name}}
	{{$user = $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{if .CmdArgs}}
	{{$d:=(toFloat (index .CmdArgs 0) ) }}
	{{if $support}}
		{{if eq $support "!malus"}}
			{{$s = "+"}}
			{{$ec = "Bonus de -"}}
			{{$w ="malus"}}
		{{else if eq $support "!soutien"}}
			{{$s = "-"}}
			{{$ec = "Malus de +"}}
			{{$w = "bonus"}}
		{{end}}
		{{if eq $d (toFloat 0)}}
			{{$embed:=cembed
				"description" "Ultra critique : Votre cible a un {{$w}} de {{$s}}4."
				"color" $supportcol
				"author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if or (eq $d (toFloat 1)) (eq $d (toFloat 2) ) }}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible a un {{$w}} de {{$s}}3."
				"color" $supportcol
				"author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if and (le $d (toFloat 5)) (ge $d (toFloat 3))}}
			{{$embed:=cembed
				"description" "Votre cible a un {{$w}} de {{$s}}2."
				"color" $supportcol
        "author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if  and (le $d (toFloat 8)) (ge $d (toFloat 6))}}
			{{$embed:=cembed
				"description" "Votre cible a un {{$w}} de {{$s}}1."
				"color" $supportcol
        "author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 9)}}
			{{$embed:=cembed
				"description" "Echec"
				"color" $supportcol
        "author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 10)}}
			{{$embed:=cembed
				"description" "Echec critique : Votre cible a un {{ec}}2."
				"color" $supportcol
        "author" (sdict "name" $user "icon_url" $supportimg)}}
			{{$id = sendMessageRetID nil $embed}}
    {{else}}
      {{$embed:=cembed
        "title" "ERREUR"
        "description" "**Usage** : `$support <valeur>`\n Votre valeur doit être entre 0 et 10."
        "color" 0xD83333}}
      {{$id = sendMessageRetID nil $embed}}
		{{end}}

	{{else if $soin}}
		{{if eq $d (toFloat 0)}}
			{{$embed:=cembed
				"description" "Ultra critique : Votre cible regagne 40 PV *(+5 si module/PSI)* et obtient un bonus de votre choix."
				"color" $soincol
				"author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 1)}}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible regagne 40 PV *(+5 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 9)}}
    {{$embed:=cembed
      "description" "Echec..."
      "color" $soincol
      "author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 10)}}
			{{$embed:=cembed
				"description" "Echec critique : Votre cible obtient un support."
				"color" $soincol
        "author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else}}
		{{$regen:= sub 9 $d}}
		{{$regen := mult $regen 5}}
		{{$r := cembed
			"author" (sdict "name" $user "icon_url" $soinimg)
				"description" (joinStr "" "**Réussite : ** *+" $regen " PV (+5 si capacité).*\n")
				"color" 0xEFA3EA }}
				{{$id = sendMessageRetID nil $r}}


		{{end}}

  {{else}}
		{{$val:= sub $d 9}}
		{{if eq $d (toFloat 0)}}
			{{$urc := cembed
				"author" (sdict "name" $user "icon_url" $poisonimg)
				"description" "**Ultra critique** : Votre cible a un empoisonnement de 10 PV sur 3 tours (-30 PV en tout).\n"
				"color" 0x4D399E}}
				{{sendMessage nil $urc}}
		{{else if eq $d (toFloat 1)}}
		{{$rc := cembed
			"author" (sdict "name" $user "icon_url" $poisonimg)
			"description" "**Réussite critique** : Votre cible a un empoisonnement de 40 PV sur 3 tours (-120 PV en tout).\n"
			"color" 0x4D399E }}
			{{sendMessage nil $rc}}
		{{else if eq $d (toFloat 9)}}
		{{$echec := cembed
			"author" (sdict "name" $user "icon_url" $poisonimg)
			"description" "**Echec de l'empoisonnement**...\n"
			"color" 0x4D399E }}
			{{sendMessage nil $echec}}
		{{else if eq $d (toFloat 10)}}
		{{$ec:= cembed
			"author" (sdict "name" $user "icon_url" $poisonimg)
			"description" "**Echec critique de l'empoisonnement :** Votre cible gagne une régénération de 15 PV sur 3 tours (+90 PV en tout).\n"
			"color" 0x4D399E }}
			{{sendMessage nil $ec}}
		{{else}}
			{{$p:= sub 10 $d}}
			{{$pa:= mult $p 3}}
			{{$r := cembed
				"author" (sdict "name" $user "icon_url" $poisonimg)
				"description" (joinStr "" "**Réussite** : Votre cible a un empoisonnement de " $p " PV sur 3 tours (-" $pa " PV en tout).\n")
				"color" 0x4D399E }}
				{{sendMessage nil $r}}
		{{end}}
  {{end}}
{{end}}
