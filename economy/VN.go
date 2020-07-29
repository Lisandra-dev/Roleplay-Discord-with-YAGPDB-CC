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

{{$flag2 := reFind `-(add|use)` .Message.Content}}
{{$flag := reFind `\-(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu)` .Message.Content}}

{{if .CmdArgs}}
	{{if eq $flag2 "-add"}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 2)}}
			{{$armes = $armes.Append $item}}
			{{dbSet .Server.ID "arme" $armes}}
		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$mod := title (index .CmdArgs 2)}}
			{{$module = $module.Append $mod}}
			{{dbSet .Server.ID "module" $module}}
		{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "biocomposant"}}
			{{$bc := add $value $x}}
			{{$compo.Set "biocomposant" $bc}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "cytomorphe"}}
			{{$lc := add $value $x}}
			{{$compo.Set "cytomorphe" $lc}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
			{{$value := $compo.Get "bionotropique"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$cb := add $value $x}}
			{{$compo.Set "bionotropique" $cb}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
			{{$value := $compo.Get "ferreux"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$sf := add $value $x}}
			{{$compo.Set "ferreux" $sf}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
			{{$value := $compo.Get "universel"}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$cu := add $value $x}}
			{{$compo.Set "universel" $cu}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else}}
		**Usage** : `$vn -add -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu) <valeur>`
		Pour les armes : Donnez le nom.
		Pour les composants, donner la valeur.
		{{end}}

	{{else if eq $flag2 "-use"}}
		{{if or (eq $flag "-arme") (eq $flag "-armes")}}
			{{$item := title (index .CmdArgs 2)}}
			{{$newarme := cslice}}
			{{range $armes}}
				{{- if ne . $item }}
					{{- $newarme = $newarme.Append .}}
				{{- end}}
			{{- end}}
			{{dbSet .Server.ID "arme" $newarme}}
		{{else if or (eq $flag "-module") (eq $flag "-modules")}}
			{{$mod := title (index .CmdArgs 2)}}
			{{$newmod := cslice}}
			{{range $module}}
				{{- if ne . $mod }}
					{{- $newmod = $newmod.Append .}}
				{{- end}}
			{{- end}}
			{{dbSet .Server.ID "module" $newmod}}
		{{else if or (eq $flag "-bc") (eq $flag "-BC")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "biocomposant"}}
			{{$bc := sub $value $x}}
			{{if lt $bc (toInt 0)}}
				{{$bc = (toInt 0)}}
			{{end}}
			{{$compo.Set "biocomposant" $bc}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-LC") (eq $flag "-lc")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "cytomorphe"}}
			{{$lc := sub $value $x}}
			{{if lt $lc (toInt 0)}}
				{{$lc = (toInt 0)}}
			{{end}}
			{{$compo.Set "cytomorphe" $lc}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-CB") (eq $flag "-cb")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "bionotropique"}}
			{{$cb := sub $value $x}}
			{{if lt $cb (toInt 0)}}
				{{$cb = (toInt 0)}}
			{{end}}
			{{$compo.Set "bionotropique" $cb}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-sf") (eq $flag "-SF")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "ferreux"}}
			{{$sf := sub $value $x}}
			{{if lt $sf (toInt 0)}}
				{{$sf = (toInt 0)}}
			{{end}}
			{{$compo.Set "ferreux" $sf}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else if or (eq $flag "-CU") (eq $flag "-cu")}}
			{{$x := (toInt (index .CmdArgs 2))}}
			{{$value := $compo.Get "universel"}}
			{{$cu := sub $value $x}}
			{{if lt $cu (toInt 0)}}
				{{$cu = (toInt 0)}}
			{{end}}
			{{$compo.Set "universel" $cu}}
			{{dbSet .Server.ID "compo" $compo}}
		{{else}}
		**Usage** : `$vn -use -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu) <valeur>`
		Pour les armes : Donnez le nom.
		Pour les composants, donner la valeur.
		*Note : Pour les modules ou les armes dont le nom n'est pas reconnu, il n'y aura pas de soucis : rien ne sera supprimé, mais le poste sera tout de même mis à jour, sans aucune édition, du coup.*
		{{end}}
	{{end}}
{{else}}
**Usage** : `$vn -(add|use) -(armes?|modules?|BC|LC|CB|SF|CU|bc|lc|cb|sf|cu) <valeur>`
Pour les armes : Donnez le nom.
Pour les composants : donner la valeur.
{{end}}


{{$weapon := ""}}
{{$modu := ""}}
{{$armeval := (dbGet .Server.ID "arme").Value}}
{{range $armeval}}
	{{- $weapon = (print  $weapon ":white_small_square:" . "\n")}}
{{- end}}
{{$modval := (dbGet .Server.ID "module").Value}}
{{range $modval}}
	{{- $modu = (print $modu ":white_small_square:" . "\n")}}
{{- end}}

{{$bioc := $compo.Get "biocomposant"}}
{{$cyto := $compo.Get "cytomorphe"}}
{{$biono := $compo.Get "bionotropique"}}
{{$ferreux := $compo.Get "ferreux"}}
{{$univ := $compo.Get "universel"}}
{{$icon := (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}

{{$embed := cembed
"title" "Inventaire du Nucleus"
"thumbnail" (sdict "url" "https://i.imgur.com/Zt2aZl4.png")
"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)
"description" (joinStr " " "**Composant** \n :white_small_square: Biocomposant : " $bioc "\n:white_small_square: Liquide cytomorphe :" $cyto "\n:white_small_square: Cellule bionotropique :" $biono "\n :white_small_square: Substrat ferreux :" $ferreux "\n :white_small_square: Composant universel :" $univ "\n\n **Armes libres** :\n" $weapon "\n\n **Modules libres :**\n" $modu )
"footer" (sdict "name" "Edition |")
"timestamp" currentTime
"color" 0x8CBAEF}}

{{editMessage 722755391498485800 736003604221001790 (complexMessageEdit "embed" $embed "content" "")}}
