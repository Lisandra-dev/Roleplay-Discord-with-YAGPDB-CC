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

{{$flag := reFind `\-(?i)(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|chargeur)` .Message.Content}}

{{if .CmdArgs}}
	{{if or (eq $flag "-arme") (eq $flag "-armes")}}
		{{$item := title (index .CmdArgs 1)}}
		{{$newarme := cslice}}
		{{range $armes}}
			{{- if ne . $item }}
				{{- $newarme = $newarme.Append .}}
			{{- else if eq . $item}}
				{{- $inv.Set $item (add ($inv.Get $item) 1)}}
				{{$user}} a retiré {{$item}} de l'inventaire du Nucleus.
			{{- end}}
		{{- end}}
		{{dbSet .Server.ID "arme" $newarme}}

	{{else if or (eq $flag "-module") (eq $flag "-modules")}}
		{{$mod := title (index .CmdArgs 1)}}
		{{$newmod := cslice}}
		{{range $module}}
			{{- if ne . $mod }}
				{{- $newmod = $newmod.Append .}}
			{{- else if eq . $mod}}
				{{- $inv.Set $mod (add ($inv.Get $mod) 1)}}
				{{$user}} a retiré {{$mod}} de l'inventaire du Nucleus.
			{{- end}}
		{{- end}}
		{{dbSet .Server.ID "module" $newmod}}

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
			{{$user}} a retiré {{$x}} biocomposant(s) de l'inventaire du Nucleus.
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
			{{$user}} a retiré {{$x}} liquide(s) cytomorphe(s) de l'inventaire du Nucleus.
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
			{{$user}} a retiré {{$x}} cellule(s) bionotropique(s) de l'inventaire du Nucleus.
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
			{{$user}} a retiré {{$x}} substrat(s) ferreux de l'inventaire du Nucleus.
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
			{{$user}} a retiré {{$x}} composant(s) universel(s) de l'inventaire du Nucleus.
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
			{{$user}} a retiré {{index .CmdArgs 1}} {{$item}} de l'inventaire du Nucleus.
		{{else}}
			Il n'y a pas assez de {{$item}} sur le vaisseau pour faire cela.
		{{end}}

	{{else}}
	**Usage** : `$vn -use -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu|-chargeur) <valeur>`
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
{{dbSet $id "economy" $userEco}}
