{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}

{{$armelist := dbGet .Server.ID "arme"}}
{{$armes := cslice}}
{{if $armelist}}
	{{$armes = (cslice).AppendSlice $armelist.Value}}
{{else}}
	{{$armes = cslice}}
{{end}}

{{$modulist := (dbGet .Server.ID "module")}}
{{$module := cslice}}
{{if $modulist}}
	{{$module = (cslice).AppendSlice $modulist.Value}}
{{else}}
	{{$module = cslice}}
{{end}}

{{$chargeur := sdict}}
{{with (dbGet 0 "chargeur_Multi")}}
	{{$chargeur = sdict .Value}}
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

{{$flag2 := reFind `-(add|use)` .Message.Content}}
{{$flag := reFind `\-(?i)(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur)` .Message.Content}}

{{if .CmdArgs}}
	{{if eq $flag2 "-add"}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 2)}}
			{{if $inv.Get $item}}
				{{$armes = $armes.Append $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{dbSet .Server.ID "arme" $armes}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$item := title (index .CmdArgs 2)}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (sub ($inv.Get $item) 1)}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
				{{$module = $module.Append $item}}
				{{dbSet .Server.ID "module" $module}}
			{{else}}
				{{$user}} ne possède pas {{$item}}.
			{{end}}

		{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "biocomposant"}}
			{{$bc := add $value $x}}
			{{if and ($inv.Get "biocomposant") (ge ($inv.Get "biocomposant") $x)}}
				{{$compo.Set "biocomposant" $bc}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "biocomposant" (sub ($inv.Get "biocomposant") $x)}}
				{{if eq ($inv.Get "biocomposant") 0}}
					{{$inv.Del "biocomposant"}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de biocomposants pour faire cela.
			{{end}}

		{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "cytomorphe"}}
			{{$lc := add $value $x}}
			{{if and ($inv.Get "cytomorphe") (ge ($inv.Get "cytomorphe") $x)}}
				{{$compo.Set "cytomorphe" $lc}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "cytomorphe" (sub ($inv.Get "cytomorphe") $x)}}
				{{if eq ($inv.Get "cytomorphe") 0}}
					{{$inv.Del "cytomorphe"}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de liquide cytomorphe pour faire cela.
			{{end}}

		{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
			{{$value := $compo.Get "bionotropique"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$cb := add $value $x}}
			{{if and ($inv.Get "bionotropique") (ge ($inv.Get "bionotropique") $x)}}
				{{$compo.Set "bionotropique" $cb}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "bionotropique" (sub ($inv.Get "bionotropique") $x)}}
				{{if eq ($inv.Get "bionotropique") 0}}
					{{$inv.Del "bionotropique"}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de cellule bionotropique pour faire cela.
			{{end}}

		{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
			{{$value := $compo.Get "ferreux"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$sf := add $value $x}}
			{{if and ($inv.Get "ferreux") (ge ($inv.Get "ferreux") $x)}}
				{{$compo.Set "ferreux" $sf}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "ferreux" (sub ($inv.Get "ferreux") $x)}}
				{{if eq ($inv.Get "ferreux") 0}}
					{{$inv.Del "ferreux"}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de substrats ferreux pour faire cela.
			{{end}}


		{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
			{{$value := $compo.Get "universel"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$cu := add $value $x}}
			{{if and ($inv.Get "universel") (ge ($inv.Get "universel") $x)}}
				{{$compo.Set "universel" $cu}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "universel" (sub ($inv.Get "universel") $x)}}
				{{if eq ($inv.Get "universel") 0}}
					{{$inv.Del "universel"}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de composants universels pour faire cela.
			{{end}}

		{{else if eq $flag "-chargeur" "-Chargeur"}}
			{{$balle := reFind `(?i)(fusil|pistolet|canon)` .Message.Content}}
			{{$item := print "[CHARGEUR] " $balle}}
			{{$value := $chargeur.Get $item}}
			{{$q:= toInt (index .CmdArgs 2)}}
			{{$x := add $value (toInt (index .CmdArgs 2))}}
			{{if and ($inv.Get $item) (ge ($inv.Get $item) $q)}}
				{{$compo.Set $item $x}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set $item (sub ($inv.Get $item) $q)}}
				{{dbSet 0 "chargeur_Multi" $chargeur}}
				{{if eq ($inv.Get $item) 0}}
					{{$inv.Del $item}}
				{{end}}
			{{else}}
				{{$user}} n'a pas assez de {{$item}} pour faire cela.
			{{end}}

		{{else}}
		**Usage** : `$vn -add -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur) <valeur>`

		{{end}}

	{{else if eq $flag2 "-use"}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 2)}}
			{{$newarme := cslice}}
			{{range $armes}}
				{{- if ne . $item }}
					{{- $newarme = $newarme.Append .}}
				{{- else if eq . $item}}
					{{- $inv.Set $item (add ($inv.Get $item) 1)}}
				{{- end}}
			{{- end}}
			{{dbSet .Server.ID "arme" $newarme}}

		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$mod := title (index .CmdArgs 2)}}
			{{$newmod := cslice}}
			{{range $module}}
				{{- if ne . $mod }}
					{{- $newmod = $newmod.Append .}}
				{{- else if eq . $mod}}
					{{- $inv.Set $mod (add ($inv.Get $mod) 1)}}
				{{- end}}
			{{- end}}
			{{dbSet .Server.ID "module" $newmod}}

		{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "biocomposant"}}
			{{if le $x $value }}
				{{$bc := sub $value $x}}
				{{if lt $bc (toInt 0)}}
					{{$bc = (toInt 0)}}
				{{end}}
				{{$compo.Set "biocomposant" $bc}}
				{{$inv.Set "biocomposant" (add $x ($inv.Get "biocomposant"))}}
				{{dbSet .Server.ID "compo" $compo}}
			{{else}}
				Il n'y a pas assez de biocomposants sur le vaisseau pour faire cela.
			{{end}}

		{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "cytomorphe"}}
			{{if le $x $value}}
				{{$lc := sub $value $x}}
				{{if lt $lc (toInt 0)}}
					{{$lc = (toInt 0)}}
				{{end}}
				{{$inv.Set "cytomorphe" (add $x ($inv.Get "cytomorphe"))}}
				{{$compo.Set "cytomorphe" $lc}}
				{{dbSet .Server.ID "compo" $compo}}
			{{else}}
				Il n'y a pas assez de liquide cytomorphe sur le vaisseau pour faire cela.
			{{end}}

		{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "bionotropique"}}
			{{if le $x $value}}
				{{$cb := sub $value $x}}
				{{if lt $cb (toInt 0)}}
					{{$cb = (toInt 0)}}
				{{end}}
				{{$compo.Set "bionotropique" $cb}}
				{{$inv.Set "bionotropique" (add $x ($inv.Get "bionotropique"))}}
				{{dbSet .Server.ID "compo" $compo}}
			{{else}}
				Il n'y a pas assez de cellule bionotropique sur le vaisseau pour faire cela.
			{{end}}

		{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "ferreux"}}
			{{if le $x $value}}
				{{$sf := sub $value $x}}
				{{if lt $sf (toInt 0)}}
					{{$sf = (toInt 0)}}
				{{end}}
				{{$compo.Set "ferreux" $sf}}
				{{dbSet .Server.ID "compo" $compo}}
				{{$inv.Set "ferreux" (add $x ($inv.Get "ferreux"))}}
			{{else}}
				Il n'y a pas assez de substrats ferreux sur le vaisseau pour faire cela.
			{{end}}

		{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "universel"}}
			{{if le $x $value}}
				{{$cu := sub $value $x}}
				{{if lt $cu (toInt 0)}}
					{{$cu = (toInt 0)}}
				{{end}}
				{{$compo.Set "universel" $cu}}
				{{$inv.Set "universel" (add $x ($inv.Get "universel"))}}
				{{dbSet .Server.ID "compo" $compo}}
			{{else}}
				Il n'y a pas assez de composants universels sur le vaisseau pour faire cela.
			{{end}}

		{{else if eq $flag "-chargeur" "-Chargeur"}}
			{{$balle := reFind `(?i)(fusil|pistolet|canon)` .Message.Content}}
			{{$item := print "[CHARGEUR] " $balle}}
			{{$value := $chargeur.Get $item}}
			{{if le (index .CmdArgs 2) $value}}
				{{$x := sub $value (toInt (index .CmdArgs 2))}}
				{{if lt $x (toInt 0)}}
					{{$x = 0}}
				{{end}}
				{{$chargeur.Set $item $x}}
				{{$inv.Set $item (add ($inv.Get $item) (toInt (index .CmdArgs)))}}
				{{dbSet 0 "chargeur_Multi" $chargeur}}
			{{else}}
				Il n'y a pas assez de {{$item}} sur le vaisseau pour faire cela.
			{{end}}

		{{else}}
		**Usage** : `$vn -use -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|-chargeur) <valeur>`
		{{end}}
	{{end}}
{{else}}
**Usage** : `$vn -(add|use) -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu) <valeur>`
{{end}}


{{$weapon := ""}}
{{$modu := ""}}
{{$armeval := (dbGet .Server.ID "arme").Value}}
{{range $armeval}}
	{{- $weapon = (print  $weapon "▫️" . "\n")}}
{{- end}}
{{$modval := (dbGet .Server.ID "module").Value}}
{{range $modval}}
	{{- $modu = (print $modu "▫️" . "\n")}}
{{- end}}

{{$bioc := $compo.Get "biocomposant"}}
{{$cyto := $compo.Get "cytomorphe"}}
{{$biono := $compo.Get "bionotropique"}}
{{$ferreux := $compo.Get "ferreux"}}
{{$univ := $compo.Get "universel"}}
{{$fusil := $chargeur.Get "[CHARGEUR] Fusil"}}
{{if not $fusil}}
	{{$fusil = 0}}
{{end}}
{{$pistolet := $chargeur.Get "[CHARGEUR] Pistolet"}}
{{if not $pistolet}}
	{{$pistolet = 0}}
{{end}}
{{$canon := $chargeur.Get "[CHARGEUR] Canon"}}
{{if not $canon}}
	{{$canon = 0}}
{{end}}
{{$icon := (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}

{{$embed := cembed
"title" "Inventaire du Nucleus"
"thumbnail" (sdict "url" "https://i.imgur.com/Zt2aZl4.png")
"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)
"description" (joinStr " " "**Composant** \n ▫️ Biocomposant : " $bioc "\n▫️ Liquide cytomorphe :" $cyto "\n▫️ Cellule bionotropique :" $biono "\n ▫️ Substrat ferreux :" $ferreux "\n ▫️ Composant universel :" $univ "\n\n **Armes libres** :\n" $weapon "\n\n **Modules libres :**\n" $modu "\n\n **Chargeurs disponibles : **\n ▫️ Pistolet : " $pistolet "\n ▫️ Fusil : " $fusil "\n ▫️ Canon : " $canon  )
"footer" (sdict "name" "Edition |")
"timestamp" currentTime
"color" 0x8CBAEF}}

{{editMessage 722755391498485800 736003604221001790 (complexMessageEdit "embed" $embed "content" "")}}
{{$userEco.Set "Inventory" $inv}}
{{dbSet $target "economy" $userEco}}
