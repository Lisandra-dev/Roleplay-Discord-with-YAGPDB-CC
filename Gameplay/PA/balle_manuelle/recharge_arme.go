{{$arme := "https://i.imgur.com/WNuPWCv.png"}}


{{/* Joueur */}}

{{/* Databases */}}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
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

{{/* get PA */}}
{{$pa := $groupe.Get (str $id)}}
{{if not $pa}}
	{{$groupe.Set (str $id) 4}}
	{{$pa = 4}}
{{end}}


{{/* Inventaire */}}
{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$inv := sdict}}
{{with ($userEco.Get "Inventory")}}
	{{$inv = sdict .}}
{{end}}

{{/* Msg */}}
{{$desc := ""}}
{{$bul := 0}}

{{if .CmdArgs}}
{{/* chargeur : CHANGER LA LIGNE ICI SI NOM DIFFERENTS
	Il est possible de créer une fusion des deux avec : (joinStr "" "Chargeur : " (title (index .CmdArgs 0)))
	LE NOM DOIT ETRE LE MEME DANS LA BOUTIQUE/INVENTAIRE DE L'UTILISATEUR !!!!!!!!!!*/}}
	{{$char := $inv.Get (title (index .CmdArgs 0))}}

{{/* Fusil */}}
{{if ge $pa 2}}

	{{if eq (index .CmdArgs 0) "fusil"}}
		{{if gt (toFloat (dbGet $id "fusil").Value) (toFloat 11)}}
			{{if gt (toInt $char) 0}}
				{{dbDel $id "fusil"}}
				{{$char = sub $char 1}}
				{{$inv.Set (index .CmdArgs 0) $char }}
				{{$desc = joinStr " " "Fusil rechargé.\n Il vous reste" $char "chargeurs pleins." }}
			{{else}}
				{{$desc = "Plus de chargeur dans l'inventaire."}}
			{{end}}
		{{else}}
			{{if (dbGet $id "fusil")}}
				{{$bul = (dbGet $id "fusil").Value}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le chargeur."}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "fusil2"}}
		{{if gt (toFloat (dbGet $id "fusil2").Value (toFloat 11))}}
			{{if gt (toInt $char) 0}}
				{{dbDel $id "fusil"}}
				{{$char = sub $char 1}}
				{{$inv.Set (index .CmdArgs 0) $char }}
				{{$desc = joinStr " " "Fusil secondaire rechargé.\n Il vous reste" $char "chargeurs pleins." }}
			{{else}}
				{{$desc = "Aucun chargeur disponible."}}
			{{end}}
		{{else}}
			{{if (dbGet $id "fusil2")}}
				{{$bul = (dbGet $id "fusil2").Value}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le chargeur."}}
		{{end}}

{{/* Pistolet */}}
	{{else if eq (index .CmdArgs 0) "pistolet"}}
		{{if gt (toFloat (dbGet $id "pistol").Value (toFloat 7))}}
			{{if gt (toInt $char) 0}}
				{{dbDel $id "pistol"}}
				{{$char = sub $char 1}}
				{{$inv.Set (index .CmdArgs 0) $char }}
				{{$desc = joinStr " " "Pistolet rechargé.\n Il vous reste" $char "chargeurs pleins." }}
			{{else}}
				{{$desc = "Aucun chargeur disponible."}}
			{{end}}
		{{else}}
			{{if (dbGet $id "pistol")}}
				{{$bul = (dbGet $id "pistol").Value}}
			{{else}}
				{{$bul = 8}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le chargeur."}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet2"}}
		{{if gt (toFloat (dbGet $id "pistol2").Value (toFloat 7))}}
			{{if gt (toInt $char) 0}}
				{{dbDel $id "pistol2"}}
				{{$char = sub $char 1}}
				{{$inv.Set (index .CmdArgs 0) $char }}
				{{$desc = joinStr " " "Pistolet secondaire rechargé.\n Il vous reste" $char "chargeurs pleins." }}
			{{else}}
				{{$desc = "Aucun chargeur disponible."}}
			{{end}}
		{{else}}
			{{if (dbGet $id "pistol2")}}
				{{$bul = (dbGet $id "pistol2").Value}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le chargeur."}}
		{{end}}

{{/* Canon */}}

	{{else if eq (index .CmdArgs 0) "canon"}}
		{{if gt (toFloat (dbGet $id "canon").Value (toFloat 19))}}
			{{if gt (toInt $char) 0}}
				{{dbDel $id "pistol2"}}
				{{$char = sub $char 1}}
				{{$inv.Set (index .CmdArgs 0) $char }}
				{{$desc = joinStr " " "Canon rechargé.\n Il vous reste" $char "chargeurs pleins." }}
			{{else}}
				{{$desc = "Aucun chargeur disponible."}}
			{{end}}
		{{else}}
			{{if (dbGet $id "canon")}}
				{{$bul = (dbGet $id "fusil2").Value}}
			{{else}}
				{{$bul = 20}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le chargeur."}}
		{{end}}


	{{else}}
		{{$desc := "**Usage** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`"}}
	{{end}}

{{else}}
	{{$desc := "Vous n'avez pas les PA pour réaliser cette action."}}
{{end}}
{{else}}
	{{$desc := "**Usage** : `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`"}}
{{end}}

{{$embed := cembed
	"author" (sdict "name" $user "icon_url" $arme)
	"color" 0x6CAB8E
	"description" $desc}}
{{$idm := sendMessageRetID nil $embed}}
{{deleteMessage nil $idm 30}}
{{deleteTrigger 1}}
