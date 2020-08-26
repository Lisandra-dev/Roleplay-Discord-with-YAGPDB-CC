{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}

{{$armes := sdict}}
{{with (dbGet 0 "armelist")}}
	{{$armes = sdict .Value}}
{{end}}

{{$modulist := (dbGet .Server.ID "module")}}
{{$module := cslice}}
{{if $modulist}}
	{{$module = (cslice).AppendSlice $modulist.Value}}
{{else}}
	{{$module = cslice}}
{{end}}

{{$implist := dbGet 0 "implant"}}
{{$implant := cslice}}
{{if $implist}}
	{{$implant = (cslice).AppendSlice $implist.Value}}
{{else}}
	{{$implant = cslice}}
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

{{$flag := reFind `\-(?i)(armes?|modules?|implant(s)|soin(s)|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur)` .Message.Content}}

{{if .CmdArgs}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 1)}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if $armes.Get $item}}
					{{$armes.Set (add ($armes.Get $item) 1)}}
				{{else}}
					{{$armes.Set $item 1}}
				{{end}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$user}} a posé {{$item}} dans l'inventaire du Nucleus.
				{{dbSet 0 "armelist" $armes}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

		{{else if eq $flag "-soin" "-soins"}}
			{{$item := title (index .CmdArgs 1)}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if $soin.Get $item}}
					{{$soin.Set (add ($soin.Get $item) 1)}}
				{{else}}
					{{$soin.Set $item 1}}
				{{end}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$user}} a posé {{$item}} dans l'inventaire du Nucleus.
				{{dbSet 0 "soin" $soin}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$item := title (index .CmdArgs 1)}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$module = $module.Append $item}}
				{{dbSet .Server.ID "module" $module}}
				{{$user}} a posé {{$item}} dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

			{{else if eq $flag "-implant" "-implants"}}
				{{$item := title (index .CmdArgs 1)}}
				{{if $inv.Get $item}}
					{{$inv.Set $item (sub ($inv.Get $item) 1)}}
					{{if eq ($inv.Get $item) 0}}
						{{$inv.Del $item}}
					{{end}}
					{{$implant = $implant.Append $item}}
					{{dbSet 0 "implant" $implant}}
					{{$user}} a posé {{$item}} dans l'inventaire du Nucleus.
				{{else}}
					{{$user}} ne possède pas {{$item}}.
				{{end}}

		{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
			{{$x := (toInt (index .CmdArgs 1))}}
			{{$value := $compo.Get "biocomposant"}}
			{{$bc := add $value $x}}
			{{if and ($inv.Get "[C] Biocomposant") (ge (toInt ($inv.Get "[C] Biocomposant")) $x)}}
				{{$compo.Set "biocomposant" $bc}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "[C] Biocomposant" (sub ($inv.Get "[C] Biocomposant") $x)}}
				{{if eq ($inv.Get "[C] Biocomposant") 0}}
					{{$inv.Del "[C] Biocomposant"}}
				{{end}}
					{{$user}} a posé {{$x}} biocomposant(s) dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de biocomposants pour faire cela.
			{{end}}

		{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
			{{$x := (toInt (index .CmdArgs 1))}}
			{{$value := $compo.Get "cytomorphe"}}
			{{$lc := add $value $x}}
			{{if and ($inv.Get "[C] Liquide Cytomorphe") (ge (toInt ($inv.Get "[C] Liquide Cytomorphe")) $x)}}
				{{$compo.Set "cytomorphe" $lc}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "[C] Liquide Cytomorphe" (sub ($inv.Get "[C] Liquide Cytomorphe") $x)}}
				{{if eq ($inv.Get "[C] Liquide Cytomorphe") 0}}
					{{$inv.Del "[C] Liquide Cytomorphe"}}
				{{end}}
				{{$user}} a posé {{$x}} liquide cytomorphe dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de liquide(s) cytomorphe(s) pour faire cela.
			{{end}}

		{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
			{{$value := $compo.Get "bionotropique"}}
			{{$x := (toInt (index .CmdArgs 1))}}
			{{$cb := add $value $x}}
			{{if and ($inv.Get "[C] Cellule Bionotropique") (ge (toInt ($inv.Get "[C] Cellule Bionotropique")) $x)}}
				{{$compo.Set "bionotropique" $cb}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "[C] Cellule Bionotropique" (sub ($inv.Get "[C] Cellule Bionotropique") $x)}}
				{{if eq ($inv.Get "[C] Cellule Bionotropique") 0}}
					{{$inv.Del "[C] Cellule Bionotropique"}}
				{{end}}
				{{$user}} a posé {{$x}} cellule bionotropique dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de cellule(s) bionotropique(s) pour faire cela.
			{{end}}

		{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
			{{$value := $compo.Get "ferreux"}}
			{{$x := (toInt (index .CmdArgs 1))}}
			{{$sf := add $value $x}}
			{{if and ($inv.Get "[C] Substrat Ferreux") (ge (toInt ($inv.Get "[C] Substrat Ferreux")) $x)}}
				{{$compo.Set "ferreux" $sf}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "[C] Substrat Ferreux" (sub ($inv.Get "[C] Substrat Ferreux") $x)}}
				{{if eq ($inv.Get "[C] Substrat Ferreux") 0}}
					{{$inv.Del "[C] Substrat Ferreux"}}
				{{end}}
				{{$user}} a posé {{$x}} substrat(s) ferreux dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de substrats ferreux pour faire cela.
			{{end}}

		{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
			{{$value := $compo.Get "universel"}}
			{{$x := (toInt (index .CmdArgs 1))}}
			{{$cu := add $value $x}}
			{{if and ($inv.Get "[C] Composant Universel") (ge (toInt ($inv.Get "[C] Composant Universel")) $x)}}
				{{$compo.Set "universel" $cu}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "[C] Composant universel" (sub ($inv.Get "[C] Composant Universel") $x)}}
				{{if eq ($inv.Get "[C] Composant universel") 0}}
					{{$inv.Del "[C] Composant universel"}}
				{{end}}
				{{$user}} a posé {{$x}} composant(s) universel(s) dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de composants universels pour faire cela.
			{{end}}

		{{else if eq $flag "-chargeur" "-Chargeur"}}
			{{$balle := reFind `(?i)(fusil|pistolet|canon)` .Message.Content}}
			{{$item := print "[CHARGEUR] " (title $balle)}}
			{{$value := $chargeur.Get $item}}
			{{$q:= toInt (index .CmdArgs 1)}}
			{{$x := add $value (toInt (index .CmdArgs 1))}}
			{{if and ($inv.Get $item) (ge ($inv.Get $item) $q)}}
				{{$compo.Set $item $x}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set $item (sub ($inv.Get $item) $q)}}
				{{dbSet 0 "chargeur_Multi" $chargeur}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$user}} a posé {{$q}} {{$item}} dans l'inventaire du Nucleus.
			{{else}}
				{{$user}} n'a pas assez de {{$item}} pour faire cela.
			{{end}}

		{{else}}
		**Usage** : `$vn -add -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur) <valeur>`
		{{end}}
	{{end}}
