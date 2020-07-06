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

{{if .CmdArgs}}
	{{$d:=(toFloat (index .CmdArgs 0) ) }}
	{{if $malus}}
		{{if eq $d (toFloat 0)}}
			{{$embed:=cembed
				"description" "Ultra critique : Votre cible a un malus de +4 à tous ses dé."
				"color" $maluscol
				"author" (sdict "name" "Malus" "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if or (eq $d (toFloat 1)) (eq $d (toFloat 2) ) }}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible a un malus de +3 à son dé."
				"color" $maluscol
				"author" (sdict "name" "Malus" "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if and (le $d (toFloat 5)) (ge $d (toFloat 3))}}
			{{$embed:=cembed
				"description" "Votre cible a un malus de +2 à son dé."
				"color" $maluscol
        "author" (sdict "name" "Malus" "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if  and (le $d (toFloat 8)) (ge $d (toFloat 6))}}
			{{$embed:=cembed
				"description" "Votre cible a un malus de +1 à son dé."
				"color" $maluscol
        "author" (sdict "name" "Malus" "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 9)}}
			{{$embed:=cembed
				"description" "Echec"
				"color" $maluscol
        "author" (sdict "name" "Malus" "icon_url" $malusimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 10)}}
			{{$embed:=cembed
				"description" "Echec critique : Votre cible a un bonus de -2 à son dé."
				"color" $maluscol
        "author" (sdict "name" "Malus" "icon_url" $malusimg)}}
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
				"description" "Ultra critique : Votre cible regagne 8 PV *(+1 si module/PSI)* et obtient un bonus de votre choix."
				"color" $soincol
				"author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 1)}}
			{{$embed:=cembed
				"description" "Réussite critique : Votre cible regagne 8 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 2)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 7 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 3)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 6 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 4)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 5 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 5)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 4 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 6)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 3 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 7)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 2 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 8)}}
			{{$embed:=cembed
				"description" "Votre cible regagne 1 PV *(+1 si module/PSI).*"
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 9)}}
    {{$embed:=cembed
      "description" "Echec critique : Votre cible obtient un malus."
      "color" $soincol
      "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else if eq $d (toFloat 10)}}
			{{$embed:=cembed
				"description" "Echec critique : Votre cible obtient un malus."
				"color" $soincol
        "author" (sdict "name" "Soin" "icon_url" $soinimg)}}
			{{$id = sendMessageRetID nil $embed}}
		{{else}}
			{{$embed:=cembed
				"title" "ERREUR"
				"description" "**Usage** : `$soin <valeur>`\n Votre valeur doit être entre 0 et 10."
				"color" 0xD83333}}
      {{$id = sendMessageRetID nil $embed}}
		{{end}}

  {{else}}
		{{if eq $d (toFloat 10)}}
      {{$embed:=cembed
        "description" "Ultra critique : Votre cible a un empoisonnement de 10 PV sur 3 tours (-30 PV en tout)."
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}
    {{else if or (eq $d (toFloat 1)) (eq $d (toFloat 2))}}
      {{$embed:=cembed
      "description" "Réussite critique : Votre cible a un empoisonnement de 8 PV sur 3 tours *(-24 PV en tout).*"
      "color" $poisoncol
      "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}
    {{else if or (eq $d (toFloat 3)) (eq $d (toFloat 4))}}
      {{$embed:=cembed
        "description" "Réussite critique : Votre cible a un empoisonnement de 6 PV sur 3 tours *(-18 PV en tout).*"
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
        {{$id = sendMessageRetID nil $embed}}
    {{else if or (eq $d (toFloat 5)) (eq $d (toFloat 6))}}
      {{$embed:=cembed
        "description" "Réussite critique : Votre cible a un empoisonnement de 4 PV sur 3 tours *(-12 PV en tout).*"
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}
    {{else if or (eq $d (toFloat 7)) (eq $d (toFloat 8))}}
      {{$embed:=cembed
        "description" "Réussite critique : Votre cible a un empoisonnement de 2 PV sur 3 tours *(-6 PV en tout).*"
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}
    {{else if eq $d (toFloat 9)}}
      {{$embed:=cembed
        "description" "Echec critique"
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}
    {{else if eq $d (toFloat 10)}}
      {{$embed:=cembed
        "description" "Echec critique : Votre cible gagne une régénération de 3 PV sur 3 tours *(+9 PV en tout).*"
        "color" $poisoncol
        "author" (sdict "name" "Poison" "icon_url" $poisonimg)}}
      {{$id = sendMessageRetID nil $embed}}}}
    {{else}}
      {{$embed:=cembed
        "title" "ERREUR"
        "description" "**Usage** : `$poison <valeur>`\n Votre valeur doit être entre 0 et 10."
        "color" 0xD83333}}
      {{$id = sendMessageRetID nil $embed}}
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
{{deleteMessage nil $id 30}}
