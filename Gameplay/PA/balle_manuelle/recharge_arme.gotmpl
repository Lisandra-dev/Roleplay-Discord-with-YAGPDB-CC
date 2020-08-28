{{$img := "https://i.imgur.com/WNuPWCv.png"}}


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

{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{$pa := $groupe.Get (str $id)}}
{{if not $pa}}
	{{$groupe.Set (str $id) 4}}
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

{{/* Arme */}}
{{$arme := sdict}}
{{with (dbGet $id "arme")}}
	{{$arme = sdict .Value}}
{{end}}

{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{if eq $item "Pistolet2"}}
		{{$item = "Pistolet"}}
	{{else if eq $item "Fusil2"}}
		{{$item = "Fusil"}}
	{{end}}
	{{$item = print "[CHARGEUR] " $item}}
	{{$char := $inv.Get $item}}

{{/* Fusil */}}
{{if ge $pa 2}}

	{{if eq (index .CmdArgs 0) "fusil"}}
		{{if gt (toFloat ($arme.Get "fusil")) (toFloat 11)}}
			{{if gt (toInt $char) 0}}
				{{$arme.Del "fusil"}}
				{{$char = sub $char 1}}
				{{$inv.Set $item $char }}
				{{$desc = joinStr " " "Fusil rechargé.\n Il vous reste" $char "balleurs pleins." }}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$desc = "Plus de balleur dans l'inventaire."}}
			{{end}}
		{{else}}
			{{if ($arme.Get "fusil")}}
				{{$bul = ($arme.Get "fusil")}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le balleur."}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "fusil2"}}
		{{if gt (toFloat ($arme.Get "fusil2")) (toFloat 11)}}
			{{if gt (toInt $char) 0}}
				{{$arme.Del "fusil2"}}
				{{$char = sub $char 1}}
				{{$inv.Set $item $char }}
				{{$desc = joinStr " " "Fusil secondaire rechargé.\n Il vous reste" $char "balleurs pleins." }}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$desc = "Aucun balleur disponible."}}
			{{end}}
		{{else}}
			{{if ($arme.Get "fusil2")}}
				{{$bul = ($arme.Get "fusil2")}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le balleur."}}
		{{end}}

{{/* Pistolet */}}
	{{else if eq (index .CmdArgs 0) "pistolet"}}
		{{if gt (toFloat ($arme.Get "pistol")) (toFloat 7)}}
			{{if gt (toInt $char) 0}}
				{{$arme.Del "pistol"}}
				{{$char = sub $char 1}}
				{{$inv.Set $item $char }}
				{{$desc = joinStr " " "Pistolet rechargé.\n Il vous reste" $char "balleurs pleins." }}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$desc = "Aucun balleur disponible."}}
			{{end}}
		{{else}}
			{{if ($arme.Get "pistol")}}
				{{$bul = $arme.Get "pistol"}}
			{{else}}
				{{$bul = 8}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le balleur."}}
		{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet2"}}
		{{if gt (toFloat ($arme.Get "pistol2")) (toFloat 7)}}
			{{if gt (toInt $char) 0}}
				{{$arme.Del "pistol2"}}
				{{$char = sub $char 1}}
				{{$inv.Set $item $char }}
				{{$desc = joinStr " " "Pistolet secondaire rechargé.\n Il vous reste" $char "balleurs pleins." }}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$desc = "Aucun balleur disponible."}}
			{{end}}
		{{else}}
			{{if ($arme.Get "pistol2")}}
				{{$bul = ($arme.Get "pistol2")}}
			{{else}}
				{{$bul = 12}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le balleur."}}
		{{end}}

{{/* Canon */}}

	{{else if eq (index .CmdArgs 0) "canon"}}
		{{if gt (toFloat ($arme.Get "canon") (toFloat 19))}}
			{{if gt (toInt $char) 0}}
				{{$arme.Del "canon"}}
				{{$char = sub $char 1}}
				{{$inv.Set $item $char }}
				{{$desc = joinStr " " "Canon rechargé.\n Il vous reste" $char "balleurs pleins." }}
			{{else}}
				{{$desc = "Aucun balleur disponible."}}
			{{end}}
			{{if eq ($inv.Get $item) 0}}
				{{$inv.Del $item}}
			{{end}}
		{{else}}
			{{if ($arme.Get "canon")}}
				{{$bul = ($arme.Get "fusil2")}}
			{{else}}
				{{$bul = 20}}
			{{end}}
			{{$desc = joinStr " " "Il reste encore" $bul "balles dans le balleur."}}
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
	"author" (sdict "name" $user "icon_url" $img)
	"color" 0x6CAB8E
	"description" $desc}}
{{$idm := sendMessageRetID nil $embed}}
{{deleteMessage nil $idm 30}}
{{deleteTrigger 1}}
{{dbSet $id "arme" $arme}}
{{$userEco.Set "Inventory" $inv }}
{{dbSet $id "economy" $userEco}}
