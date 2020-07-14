{$name := reFind `(\#\S*)` .Message.Content}}
{{$malus:=reFindAllSubmatches `^\;malus` .Message.Content}}
{{$soin:= reFindAllSubmatches `^\;soin` .Message.Content}}
{{$poison:= reFindAllSubmatches `^\;poison` .Message.Content}}
{{$soinimg := "https://www.sphanalytics.com/wp-content/uploads/2018/01/Health-Icon334.png"}}
{{$malusimg:="https://www.pixenli.com/image/xsms14QB"}}
{{$poisonimg:="https://www.pixenli.com/image/eJlppGzy"}}
{{$soincol:=0xb66dc1}}
{{$poisoncol:=0x370E6A}}
{{$maluscol:=0x2B7697}}
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
	{{if $malus}}
		{{if eq $d (toFloat 0)}}
			{{$embed:=cembed
				"description" "Ultra critique : Votre cible a un malus de +4 à tous ses dé."
				"color" $maluscol
				"author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if or (eq $d (toFloat 1)) (eq $d (toFloat 2) ) }}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible a un malus de +3 à son dé."
				"color" $maluscol
				"author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if and (le $d (toFloat 5)) (ge $d (toFloat 3))}}
			{{$embed:=cembed
				"description" "Votre cible a un malus de +2 à son dé."
				"color" $maluscol
        "author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if  and (le $d (toFloat 8)) (ge $d (toFloat 6))}}
			{{$embed:=cembed
				"description" "Votre cible a un malus de +1 à son dé."
				"color" $maluscol
        "author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 9)}}
			{{$embed:=cembed
				"description" "Echec"
				"color" $maluscol
        "author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 10)}}
			{{$embed:=cembed
				"description" "Echec critique : Votre cible a un bonus de -2 à son dé."
				"color" $maluscol
        "author" (sdict "name" $user "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
    {{else}}
      {{$embed:=cembed
        "title" "ERREUR"
        "description" "**Usage** : `$malus <valeur>`\n Votre valeur doit être entre 0 et 10."
        "color" 0xD83333}}
      {{$id = sendMessageRetID nil $embed}}
		{{end}}

	{{else if $soin}}
		{{if eq $d (toFloat 0)}}
			{{$embed:=cembed
				"description" "Ultra critique : Votre cible regagne 40 PV *(+1 si module/PSI)* et obtient un bonus de votre choix."
				"color" $soincol
				"author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 1)}}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible regagne 40 PV *(+1 si module/PSI).*"
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
				"description" "Echec critique : Votre cible obtient un malus."
				"color" $soincol
        "author" (sdict "name" $user "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else}}
		{{$regen:= sub 9 $d}}
		{{$regen := mult $regen 5}}
		{{$r := cembed
			"author" (sdict "name" $user "icon_url" $soinimg)
				"description" (joinStr "" "**Réussite : ** *+" $regen " PV (+1 si capacité).*\n")
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
{{else}}
  {{$embed:=cembed
  	"title" "ERREUR"
  	"description" "**Usage** : `$(soin|poison|malus) <valeur>`"
  	"color" 0xD83333}}
  {{$id = sendMessageRetID nil $embed}}
{{end}}

{{deleteTrigger 1}}
