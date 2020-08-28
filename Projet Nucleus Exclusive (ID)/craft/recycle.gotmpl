{{/* Dictionnaire d'Item */}}
	{{/* Recette */}}
{{$recipe := sdict}}
{{with (dbGet 0 "recipe")}}
	{{$recipe = sdict .Value}}
{{end}}

	{{/* Inventaire du Nucleus */}}

{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}


{{/* Personnage */}}

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

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}

{{/* Flag */}}
{{$choice := reFind `(\+VN)` .Message.Content }}
{{$quanti := reFind `q\d+` .Message.Content}}
{{$bdg := reFind `[BDG]` .Message.Content}}
{{$balle := reFind `(?i)(chargeur|module|implant|poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content}}

{{/* Variable */}}
{{$q := 1}}
{{if $quanti}}
	{{$q = joinStr "" (split $quanti "q")}}
	{{$q = toInt $q}}
{{end}}

{{$log := ""}}

{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{$reci := $item}}

	{{if eq $balle "chargeur" "CHARGEUR" "Chargeur"}}
		{{$weap := reFind `(?i)(fusil|pistolet|canon)` $item}}
		{{$item = print "[CHARGEUR] " (title $weap)}}
	{{else if eq $balle "module" "Module" "MODULE"}}
		{{$type := reFind `(?i)(Perforant|burst|soutien|altération|schéma|passif)` $item}}
		{{$reci = title (lower $type)}}
	{{else if eq $balle "implant" "Implant" "IMPLANT"}}
		{{$type := reFind `(?i)(force|résistance|cognition|furtivité|vision)` $item}}
		{{$reci = title (lower $type)}}
	{{else if eq $balle "poigne" "Poigne" "épée" "Epée" "Épée" "masse" "Masse" "projectile" "projectile" "grenade" "Grenade" "pistolet" "Pistolet" "fusil" "Fusil" "canon" "Canon" "POIGNE" "ÉPÉE" "MASSE" "PROJECTILE" "GRENADE" "PISTOLET" "FUSIL" "CANON"}}
	{{$type := reFind `(?i)(poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content}}
		{{$reci = title (lower $type)}}
	{{end}}

	{{if $bdg}}
		{{$reci = print "[BDG] " (joinStr "" (split $reci "[BDG]"))}}
	{{end}}

	{{if $recipe.Get $reci}}
		{{$i := sdict ($recipe.Get $reci)}}
		{{$bc := mult ($i.Get "Biocomposant") $q}}
		{{$sf := mult ($i.Get "Substrat Ferreux") $q}}
		{{$lc := mult ($i.Get "Liquide Cytomorphe") $q}}
		{{$cb := mult ($i.Get "Cellule Bionotropique") $q}}

		{{if and ($inv.Get $item) (ge (toInt ($inv.Get $item)) $q)}}
			{{$inv.Set $item (sub ($inv.Get $item) $q)}}
			{{if eq ($inv.Get $item) 0}}
				{{$inv.Del $item}}
			{{end}}
			{{$inv.Set "[C] Biocomposant" (div $bc 2)}}
			{{$inv.Set "[C] Substrat Ferreux" (div $sf 2)}}
			{{$inv.Set "[C] Liquide Cytomorphe" (div $lc 2)}}
			{{$inv.Set "[C] Cellule Bionotropique" (div $cb 2)}}
			{{$user}} a recyclé {{$q}} {{$item}} et obtenu :
			▫️ Biocomposant : {{div $bc 2}}
			▫️ Substrat Ferreux : {{div $sf 2}}
			▫️ Liquide Cytomorphe : {{div $lc 2}}
			▫️ Cellule Bionotropique : {{div $cb 2}}

		{{else if and ($inv.Get $item) (ge (toInt ($inv.Get $item)) $q) $choice}}
			{{$inv.Set $item (sub ($inv.Get $item) $q)}}
			{{if eq ($inv.Get $item) 0}}
				{{$inv.Del $item}}
			{{end}}
			{{$compo.Set "biocomposant" (div $bc 2)}}
			{{$compo.Set "ferreux" (div $sf)}}
			{{$compo.Set "cytomorphe" (div $lc 2)}}
			{{$compo.Set "bionotropique" (div $cb)}}
			{{$log = joinStr " " $user "a recyclé" $q $item "et obtenu :\n ▫️ Biocomposant :" (div $bc) "\n▫️ Substrat Ferreux :" (div $sf 2) "\n▫️ Liquide Cytomorphe :" (div $lc) "\n▫️ Cellule Bionotropique :" (div $cb)}}
		{{else}}
			{{$user}} ne possède pas {{$item}} en quantité suffisante ({{$q}}) pour faire le recyclage.
		{{end}}
	{{else}}
		{{$item}} n'a pas recette attitré.
	{{end}}
{{else}}
	**Usage** : `$recycle "objet" (q[1-100]) (+VN) (-bdg)`
{{end}}

{{if eq (toInt ($inv.Get "[C] Biocomposant")) 0}}
	{{$inv.Del "[C] Biocomposant"}}
{{end}}
{{if eq (toInt ($inv.Get "[C] Substrat Ferreux")) 0}}
	{{$inv.Del "[C] Substrat Ferreux"}}
{{end}}
{{if eq (toInt ($inv.Get "[C] Liquide Cytomorphe")) 0}}
	{{$inv.Del "[C] Liquide Cytomorphe"}}
{{end}}
{{if eq (toInt ($inv.Get "[C] Cellule Bionotropique")) 0}}
	{{$inv.Del "[C] Cellule Bionotropique"}}
{{end}}

{{dbSet 0 "compo" $compo}}
{{$userEco.Set "Inventory" $inv}}
{{dbSet $id "economy" $userEco}}
{{deleteTrigger 1}}
