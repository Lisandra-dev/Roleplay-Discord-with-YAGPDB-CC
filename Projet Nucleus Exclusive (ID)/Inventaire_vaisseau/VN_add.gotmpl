{{/* Here its more an example and cannot really be used on another roleplay, but you can using it like a base. You can change the name for dictionnary, componant... You can change all name. Just don't forget that if you change the name here, you must change it everywhere.

	Name used in :
	-> Inventory dict : [C] items
	-> Compo dict : diminutif without majuscule
	-> recipe : Complete name with majuscule
	You can use the function without using the compo dictionnary, delete just the "VN reFind" or, if you want to keep place, delete all text in relation with.  */}}


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

{{$flag := reFind `\-(?i)(armes?|modules?|implant(s)|soin(s)|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur)` .Message.Content}}
{{$log := ""}}

{{if .CmdArgs}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 1)}}
			{{$type := reFind `(?i)(poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if $armes.Get $item}}
					{{$armes.Set $item (add ($armes.Get $item) 1)}}
				{{else}}
					{{$armes.Set $item 1}}
				{{end}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$log = joinStr " " $user "a posé" $item "dans l'inventaire du Nucleus."}}
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
				{{$log = joinStr " " $user "a posé" $item "dans l'inventaire du Nucleus."}}
				{{dbSet 0 "soin" $soin}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$item := title (index .CmdArgs 1)}}
			{{$type := title (reFind `(?i)(Perforant|burst|soutien|altération|schéma|passif)` .Message.Content)}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$module.Set $item (add ($module.Get $item) 1)}}
				{{dbSet .Server.ID "module" $module}}
				{{$log = joinStr " " $user "a posé" $item "dans l'inventaire du Nucleus."}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

			{{else if eq $flag "-implant" "-implants"}}
				{{$item := title (index .CmdArgs 1)}}
				{{$type := title (reFind `(?i)(force|résistance|cognition|furtivité|vision)` .Message.Content)}}
				{{if $inv.Get $item}}
					{{$inv.Set $item (sub ($inv.Get $item) 1)}}
					{{if eq ($inv.Get $item) 0}}
						{{$inv.Del $item}}
					{{end}}
					{{$implant.Set $item (add ($implant.Get $item) 1)}}
					{{dbSet .Server.ID "implant" $implant}}
					{{$log = joinStr " " $user "a posé" $item "dans l'inventaire du Nucleus."}}
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
					{{$log = joinStr " " $user "a posé" $x "biocomposant(s) dans l'inventaire du Nucleus."}}
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
				{{$log = joinStr " " $user "a posé" $x "liquide(s) cytomorphe dans l'inventaire du Nucleus."}}
			{{else}}
				{{$user}} n'a pas assez de liquide(s) cytomorphe pour faire cela.
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
				{{$log = joinStr " " $user "a posé" $x "cellule(s) bionotropique dans l'inventaire du Nucleus."}}
			{{else}}
				{{$user}} n'a pas assez de cellule(s) bionotropique pour faire cela.
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
				{{$log = joinStr "" $user "a posé" $x "substrat(s) ferreux dans l'inventaire du Nucleus."}}
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
				{{$log = joinStr " " $user "a posé" $x "composant(s) universel dans l'inventaire du Nucleus."}}
			{{else}}
				{{$user}} n'a pas assez de composants universel pour faire cela.
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
				{{$log = joinStr " " $user "a posé" $q $item "dans l'inventaire du Nucleus."}}
			{{else}}
				{{$user}} n'a pas assez de {{$item}} pour faire cela.
			{{end}}

		{{else}}
		**Usage** : `$vn -add -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur) <valeur>`
		{{end}}
	{{end}}

{{$chan := 735938256038002818}}
{{sendMessage $chan $log}}
{{$userEco.Set "Inventory" $inv}}
{{dbSet $id "economy" $userEco}}
{{deleteTrigger 1}}
