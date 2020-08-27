{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}

{{$armes := sdict}}
{{with (dbGet 0 "armelist")}}
	{{$armes = sdict .Value}}
{{end}}

{{$module := sdict}}
{{with (dbGet .Server.ID "module")}}
	{{$module = sdict .Value}}
{{end}}

{{$implant := sdict }}
{{with (dbGet .Server.ID "implant")}}
	{{$implant = sdict .Value}}
{{end}}

{{$chargeur := sdict}}
{{with (dbGet 0 "chargeur_Multi")}}
	{{$chargeur = sdict .Value}}
{{end}}

{{$soin := sdict}}
{{with (dbGet 0 "soin")}}
	{{$soin = sdict .Value}}
{{end}}

{{/* Inventaire personnel */}}

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
	{{dbSetExpire $id "rerollName" $name 3600}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}

	{{/* Fonction */}}

{{$flag := reFind `\-(?i)(armes?|soin(s?)|implant(s?)|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur)` .Message.Content}}
{{$bdg := reFind `-bdg` .Message.Content}}
{{$log := ""}}

{{if .CmdArgs}}
	{{if or (eq $flag "-arme") (eq $flag "-armes")}}
		{{$item := title (index .CmdArgs 1)}}
		{{$type := title (reFind `(?i)(poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content)}}
		{{if $armes.Get $item}}
			{{$armes.Set $item (sub ($armes.Get $item) 1)}}
			{{if le ($armes.Get $item) 0}}
				{{$armes.Del $item}}
			{{end}}
			{{dbSet 0 "armelist" $armes}}
			{{$log = joinStr " " $user "a retiré" $item "de l'inventaire du Nucleus."}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (add ($inv.Get $item) 1)}}
			{{else}}
				{{$inv.Set $item 1}}
			{{end}}
		{{else}}
			{{$item}} ne fait pas parti de l'inventaire du Nucleus.
		{{end}}

	{{else if or (eq $flag "-soin") (eq $flag "-soins")}}
		{{$item := title (index .CmdArgs 1)}}
		{{if $soin.Get $item}}
			{{$soin.Set $item (sub ($soin.Get $item) 1)}}
			{{if le ($soin.Get $item) 0}}
				{{$soin.Del $item}}
			{{end}}
			{{dbSet 0 "soin" $soin}}
			{{$log = joinStr " " $user "a retiré" $item "de l'inventaire du Nucleus."}}
			{{if $inv.Get $item}}
				{{$inv.Set (add ($inv.Get $item) 1)}}
			{{else}}
				{{$inv.Set $item 1}}
			{{end}}
		{{else}}
			{{$item}} ne fait pas parti de l'inventaire du Nucleus.
		{{end}}

	{{else if or (eq $flag "-module") (eq $flag "-modules")}}
		{{$item := title (index .CmdArgs 1)}}
		{{$type := title (reFind `(?i)(Perforant|burst|soutien|altération|schéma|passif)` .Message.Content)}}
		{{if $module.Get $item}}
			{{$module.Set $item (sub ($module.Get $item) 1)}}
			{{if le ($module.Get $item) 0}}
				{{$module.Del $item}}
			{{end}}
			{{dbSet .Server.ID "module" $module}}
			{{$log = joinStr " " $user "a retiré" $item "de l'inventaire du Nucléus"}}
			{{if $inv.Get $item}}
				{{$inv.Set (add ($inv.Get $item) 1)}}
			{{else}}
				{{$inv.Set $item 1}}
			{{end}}
		{{else}}
			{{$item}} ne fait pas parti de l'inventaire du Nucléus.
		{{end}}

	{{else if or (eq $flag "-implant") (eq $flag "-implants")}}
		{{$item := title (index .CmdArgs 1)}}
		{{$type := title (reFind `(?i)(force|résistance|cognition|furtivité|vision)` .Message.Content)}}
		{{if $implant.Get $item}}
			{{$implant.Set $item (sub ($implant.Get $item) 1)}}
			{{if le ($implant.Get $item) 0}}
				{{$implant.Del $item}}
			{{end}}
			{{dbSet .Server.ID "implant" $implant}}
			{{$log = joinStr " " $user "a retiré" $item "de l'inventaire du Nucléus"}}
			{{if $inv.Get $item}}
				{{$inv.Set (add ($inv.Get $item) 1)}}
			{{else}}
				{{$inv.Set $item 1}}
			{{end}}
		{{else}}
			{{$item}} ne fait pas parti de l'inventaire du Nucléus.
		{{end}}

	{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
		{{$x := (toInt (index .CmdArgs 1))}}
		{{$value := $compo.Get "biocomposant"}}
		{{if le $x $value }}
			{{$bc := sub $value $x}}
			{{if lt $bc (toInt 0)}}
				{{$bc = (toInt 0)}}
			{{end}}
			{{$compo.Set "biocomposant" $bc}}
			{{$inv.Set "[C] Biocomposant" (add $x ($inv.Get "[C] Biocomposant"))}}
			{{dbSet .Server.ID "compo" $compo}}
			{{$log = joinStr " " $user "a retiré" $x "biocomposant(s) de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de biocomposants sur le vaisseau pour faire cela.
		{{end}}

	{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
		{{$x := (toInt (index .CmdArgs 1))}}
		{{$value := $compo.Get "cytomorphe"}}
		{{if le $x $value}}
			{{$lc := sub $value $x}}
			{{if lt $lc (toInt 0)}}
				{{$lc = (toInt 0)}}
			{{end}}
			{{$inv.Set "[C] Liquide Cytomorphe" (add $x ($inv.Get "[C] Liquide Cytomorphe"))}}
			{{$compo.Set "cytomorphe" $lc}}
			{{dbSet .Server.ID "compo" $compo}}
			{{$log = joinStr " " $user "a retiré" $x "liquide(s) cytomorphe(s) de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de liquide cytomorphe sur le vaisseau pour faire cela.
		{{end}}

	{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
		{{$x := (toInt (index .CmdArgs 1))}}
		{{$value := $compo.Get "bionotropique"}}
		{{if le $x $value}}
			{{$cb := sub $value $x}}
			{{if lt $cb (toInt 0)}}
				{{$cb = (toInt 0)}}
			{{end}}
			{{$compo.Set "bionotropique" $cb}}
			{{$inv.Set "[C] Cellule Bionotropique" (add $x ($inv.Get "[C] Cellule Bionotropique"))}}
			{{dbSet .Server.ID "compo" $compo}}
			{{$log = joinStr " " $user "a retiré" $x "cellule(s) bionotropique(s) de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de cellule bionotropique sur le vaisseau pour faire cela.
		{{end}}

	{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
		{{$x := (toInt (index .CmdArgs 1))}}
		{{$value := $compo.Get "ferreux"}}
		{{if le $x $value}}
			{{$sf := sub $value $x}}
			{{if lt $sf (toInt 0)}}
				{{$sf = (toInt 0)}}
			{{end}}
			{{$compo.Set "ferreux" $sf}}
			{{dbSet .Server.ID "compo" $compo}}
			{{$inv.Set "[C] Substrat Ferreux" (add $x ($inv.Get "[C] Substrat Ferreux"))}}
			{{$log = joinStr " " $user "a retiré" $x "substrat(s) ferreux de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de substrat ferreux sur le vaisseau pour faire cela.
		{{end}}

	{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
		{{$x := (toInt (index .CmdArgs 1))}}
		{{$value := $compo.Get "universel"}}
		{{if le $x $value}}
			{{$cu := sub $value $x}}
			{{if lt $cu (toInt 0)}}
				{{$cu = (toInt 0)}}
			{{end}}
			{{$compo.Set "universel" $cu}}
			{{$inv.Set "[C] Composant Universel" (add $x ($inv.Get "[C] Composant Universel"))}}
			{{dbSet .Server.ID "compo" $compo}}
			{{$log = joinStr " " $user "a retiré" $x "composant(s) universel(s) de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de composant universel sur le vaisseau pour faire cela.
		{{end}}

	{{else if eq $flag "-chargeur" "-Chargeur"}}
		{{$balle := reFind `(?i)(fusil|pistolet|canon)` .Message.Content}}
		{{$item := print "[CHARGEUR] " (title $balle)}}
		{{$value := $chargeur.Get $item}}
		{{if le (toInt (index .CmdArgs 1)) (toInt $value)}}
			{{$x := sub $value (toInt (index .CmdArgs 1))}}
			{{if lt $x (toInt 0)}}
				{{$x = 0}}
			{{end}}
			{{$chargeur.Set $item $x}}
			{{$inv.Set $item (add ($inv.Get $item) (toInt (index .CmdArgs 1)))}}
			{{dbSet 0 "chargeur_Multi" $chargeur}}
			{{$log = joinStr " " $user "a retiré" (index .CmdArgs 1) $item "de l'inventaire du Nucleus."}}
		{{else}}
			Il n'y a pas assez de {{$item}} sur le vaisseau pour faire cela.
		{{end}}

	{{else}}
		**Usage** : `$vn -use -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|-chargeur) <valeur>`
	{{end}}
{{else}}
	**Usage** : `$vn -(add|use) -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu) <valeur>`
{{end}}

{{$chan := 735938256038002818}}
{{sendMessage $chan $log}}
{{$userEco.Set "Inventory" $inv}}
{{dbSet $id "economy" $userEco}}
{{deleteTrigger 1}}