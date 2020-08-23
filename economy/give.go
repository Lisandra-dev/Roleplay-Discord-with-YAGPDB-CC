{{/* Variable definition */}}
{{$targetEco := sdict}}
{{$cible := ""}}
{{$invtarget := sdict }}
{{$inv := ""}}
{{$idtarget := 0}}
{{$desc := "Erreur inconnue"}}
{{$arrow := "<:next:723131844643651655>" }}
{{/* Symbol */}}
{{$mon := ""}}


{{/* Give color */}}
{{$col := 16777215}}
{{$p := 0}}
{{$r := .Member.Roles}}
{{range .Guild.Roles}}
	{{if and (in $r .ID) (.Color) (lt $p .Position)}}
	{{$p = .Position}}
	{{$col = .Color}}
	{{end}}
{{end}}

{{/* DB */}}
{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{if $serverEco.Get "symbol"}}
	{{$mon = $serverEco.Get "symbol"}}
{{end}}


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
{{$target := ""}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{if .CmdArgs}}
	{{$idtarget := .User.ID }}
	{{if (reFind `(\&\S*)` (index .CmdArgs 0))}}
		{{$target = joinStr "" (split (index .CmdArgs 0) "&")}}
		{{with (userArg (index .CmdArgs 1))}}
			{{$idtarget = .}}
			{{$idtarget = $idtarget.ID}}
		{{end}}
		{{$idC := (toRune (lower $target))}}
		{{range $idC}}
			{{- $idtarget = add $idtarget .}}
		{{- end}}
		{{$cible = title $target}}
	{{else if not (reFind `(\&\S*)` (index .CmdArgs 0))}}
		{{with (userArg (index .CmdArgs 0))}}
			{{$target = .}}
			{{$idtarget = $target.ID}}
			{{$cible = (getMember $target).Nick}}			{{end}}
	{{else}}
		{{$desc = "ERREUR : Mauvaise cible."}}
	{{end}}

	{{with (dbGet $idtarget "economy")}}
		{{$targetEco = sdict .Value}}
	{{end}}

	{{if eq (index .CmdArgs 1) "-money"}}
		{{$value := toInt (index .CmdArgs 2)}}
		{{$bal := toInt ($userEco.Get "balance")}}
		{{if gt $value $bal}}
			{{$desc = "Tu ne peux pas donner autant !"}}
		{{else}}
			{{$newbal := add $value (toInt ($targetEco.Get "balance"))}}
			{{$oldbal := sub $bal $value }}
			{{$desc = joinStr " " $value $mon "\n\n" $arrow  $cible "a maintenant" $newbal $mon "sur son compte. \n" $arrow  $user "a maintenant" $oldbal $mon " sur son compte."}}
			{{$userEco.Set "balance" $oldbal}}
			{{$targetEco.Set "balance" $newbal}}
			{{dbSet $idtarget "economy" $targetEco}}
			{{dbSet $id "economy" $userEco}}
		{{end}}

	{{else if eq (index .CmdArgs 2) "-money"}}
		{{$value := toInt (index .CmdArgs 3)}}
		{{$bal := toInt ($userEco.Get "balance")}}
		{{if gt $value $bal}}
			{{$desc = "Tu ne peux pas donner autant !"}}
		{{else}}
			{{$newbal := add $value (toInt ($targetEco.Get "balance"))}}
			{{$oldbal := sub $bal $value }}
			{{$desc = joinStr " " $value $mon " \n\n" $arrow  $cible "a maintenant" $newbal $mon " sur son compte. \n" $arrow  $user "a maintenant" $oldbal $mon " sur son compte."}}
			{{$userEco.Set "balance" $oldbal}}
			{{$targetEco.Set "balance" $newbal}}
			{{dbSet $idtarget "economy" $targetEco}}
			{{dbSet $id "economy" $userEco}}
		{{end}}

	{{else if eq (index .CmdArgs 1) "-item" "-items"}}
		{{$amount :=1}}
		{{$item := title (index .CmdArgs 2)}}
		{{$chargeur := reFind `(?i)chargeur` $item}}
		{{$compo := reFind `(?i)(bc|lc|cb|sf|cu)` $item}}
		{{if $compo}}
			{{if eq $compo "bc" "BC" "Bc"}}
				{{$item = "[C] Biocomposant"}}
			{{else if eq $compo "lc" "LC" "Lc"}}
				{{$item = "[C] Liquide Cytomorphe"}}
			{{else if eq $compo "cb" "CB" "Cb"}}
				{{$item = "[C] Cellule Bionotropique"}}
			{{else if eq $compo "sf" "SF" "Sf"}}
				{{$item = "[C] Substrat Ferreux"}}
			{{else if eq $compo "cu" "CU" "Cu"}}
				{{$item = "[C] Composant Universel"}}
			{{end}}
		{{end}}

		{{if $chargeur}}
			{{$item = reFind `(?i)(fusil|pistolet|canon)` $item}}
			{{$item = print "[CHARGEUR] " $item}}
		{{end}}

		{{if reFind `-q` (index .CmdArgs 3)}}
			{{$amount = toInt (index .CmdArgs 4)}}
		{{end}}

		{{with ($userEco.Get "Inventory")}}
			{{$inv = sdict .}}
		{{end}}
		{{with ($targetEco.Get "Inventory")}}
			{{$invtarget = sdict .}}
		{{end}}

		{{if not ($inv.Get $item)}}
			{{$desc = "Objet introuvable dans l'inventaire"}}
		{{else if gt $amount (toInt ($inv.Get $item))}}
			{{$desc = "Quantité trop élevé par rapport à la quantité disponible"}}
		{{else}}
			{{if ($invtarget.Get $item)}}
				{{$invtarget.Set $item (add (toInt ($invtarget.Get $item)) $amount)}}
			{{else}}
				{{$invtarget.Set $item $amount}}
			{{end}}
			{{$inv.Set $item (sub (toInt ($inv.Get $item)) $amount)}}
			{{if eq (toInt ($inv.Get $item)) 0}}
				{{$inv.Del $item}}
			{{end}}
			{{$userEco.Set "Inventory" $inv}}
			{{$targetEco.Set "Inventory" $invtarget}}
			{{dbSet $id "economy" $userEco}}
			{{dbSet $idtarget "economy" $targetEco}}
			{{$desc = joinStr "" $amount " " $item}}
		{{end}}

	{{else if eq (index .CmdArgs 2) "-item" "-items"}}
		{{$amount :=1}}
		{{$item := title (index .CmdArgs 3)}}
		{{if reFind `-q` (index .CmdArgs 4)}}
			{{$amount = toInt (index .CmdArgs 5)}}
		{{end}}

		{{with ($userEco.Get "Inventory")}}
			{{$inv = sdict .}}
		{{end}}

		{{with ($targetEco.Get "Inventory")}}
			{{$invtarget = sdict .}}
		{{end}}

		{{if not ($inv.Get $item)}}
			{{$desc = "Objet introuvable dans l'inventaire"}}
		{{else if gt $amount (toInt ($inv.Get $item))}}
			{{$desc = "Quantité trop élevé par rapport à la quantité disponible"}}
		{{else}}
			{{if ($invtarget.Get $item)}}
				{{$invtarget.Set $item (add (toInt ($invtarget.Get $item)) $amount)}}
			{{else}}
				{{$invtarget.Set $item $amount}}
			{{end}}
			{{$inv.Set $item (sub (toInt ($inv.Get $item)) $amount)}}
			{{if eq (toInt ($inv.Get $item)) 0}}
				{{$inv.Del $item}}
			{{end}}
			{{$userEco.Set "Inventory" $inv}}
			{{$targetEco.Set "Inventory" $invtarget}}
			{{dbSet $id "economy" $userEco}}
			{{dbSet $idtarget "economy" $targetEco}}
			{{$desc = joinStr "" $amount " " $item}}
		{{end}}



	{{else}}
	{{$desc = "**Usage** :\n- `$give <cible> -money <valeur>`\n- `$give <cible> -item <nom> (-q <valeur)`\n > Note : `-q` permet d'indiquer une quantité. Cet argument est optionnel."}}
	{{end}}
{{end}}

{{$embed := cembed
"author" (sdict "name" (joinStr " " $user "donne") "icon_url" "https://i.imgur.com/DwoqSFH.png")
"description" $desc
"footer" (sdict "text" (joinStr "" "À" $cible) "icon_url" "https://i.imgur.com/WoypxHH.png")
"color" $col }}
{{sendMessage nil $embed}}
{{deleteTrigger 1}}
