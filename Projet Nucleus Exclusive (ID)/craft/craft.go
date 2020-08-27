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
{{$choice := reFind `(\+VN|-inv)` .Message.Content }}
{{$quanti := reFind `q\d+` .Message.Content}}
{{$univ := reFind `-cu` .Message.Content}}
{{$bdg := reFind `-bdg` .Message.Content}}
{{$balle := reFind `(?i)(chargeur|module|implant|poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content}}
{{$arme := reFind `(?i)(poigne|épée|masse|projectile|grenade|pistolet|fusil|canon)` .Message.Content}}

{{/* Variable */}}

{{$q := 1}}
{{if $quanti}}
	{{$q = joinStr "" (split $quanti "q")}}
	{{$q = toInt $q}}
{{end}}

{{$log := ""}}

{{$bcp := 0}}
{{$sfp := 0}}
{{$cbp := 0}}
{{$cup := 0}}
{{$lcp := 0}}
{{$bc_vn := 0}}
{{$sf_vn := 0}}
{{$cb_vn := 0}}
{{$cu_vn := 0}}
{{$lc_vn := 0}}

{{if ($inv.Get "[C] Biocomposant")}}
	{{$bcp = $inv.Get "[C] Biocomposant"}}
{{end}}
{{if ($inv.Get "[C] Substrat Ferreux")}}
	{{$sfp =  $inv.Get "[C] Substrat Ferreux"}}
{{end}}
{{if ($inv.Get "[C] Liquide Cytomorphe")}}
	{{$lcp = $inv.Get "[C] Liquide Cytomorphe"}}
{{end}}
{{if ($inv.Get "[C] Cellule Bionotropique")}}
	{{$cbp = $inv.Get "[C] Cellule Bionotropique"}}
{{end}}
{{if ($inv.Get "[C] Composant Universel")}}
	{{$cup = $inv.Get "[C] Composant Universel"}}
{{end}}

{{if ($compo.Get "biocomposant")}}
	{{$bc_vn = $compo.Get "biocomposant"}}
{{end}}
{{if ($compo.Get "cytomorphe")}}
	{{$lc_vn = $compo.Get "cytomorphe"}}
{{end}}
{{if ($compo.Get "ferreux")}}
	{{$sf_vn = $compo.Get "ferreux"}}
{{end}}
{{if ($compo.Get "bionotropique")}}
	{{$cb_vn = $compo.Get "bionotropique"}}
{{end}}
{{if ($compo.Get "universel")}}
	{{$cu_vn = $compo.Get "universel"}}
{{end}}


{{/* Function */}}
{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 0)}}
	{{$reci := $item}}
	{{if eq $balle "chargeur" "CHARGEUR" "Chargeur"}}
		{{$weap := reFind `(?i)(fusil|pistolet|canon)` $item}}
		{{$reci = print "[CHARGEUR] " (title $weap)}}
		{{$item = $reci}}
	{{else if eq $balle "module" "Module" "MODULE"}}
		{{$type := reFind `(?i)(Perforant|burst|soutien|altération|schéma|passif)` $item}}
		{{$reci = title (lower $type)}}
	{{else if eq $balle "implant" "Implant" "IMPLANT"}}
		{{$type := reFind `(?i)(force|résistance|cognition|furtivité|vision)` $item}}
		{{$reci = title (lower $type)}}
	{{else if and $balle $arme (eq $balle $arme)}}
		{{$reci = title (lower $arme)}}
	{{end}}

	{{if $bdg}}
		{{$reci = print "[BDG] " $reci}}
		{{$item = print "[BDG] " $item}}
	{{end}}

	{{if ($recipe.Get $reci)}}
		{{$i := sdict ($recipe.Get $reci)}}
		{{$bc := mult ($i.Get "Biocomposant") $q}}
		{{$sf := mult ($i.Get "Substrat Ferreux") $q}}
		{{$lc := mult ($i.Get "Liquide Cytomorphe") $q}}
		{{$cb := mult ($i.Get "Cellule Bionotropique") $q}}

		{{if and (ge $bcp $bc) (ge $sfp $sf) (ge $lcp $lc) (ge $cbp $cb) (not $univ)}}
			{{$user}} a fabriqué {{$q}} {{$item}} en utilisant ses propres ressources.
			{{if ($inv.Get $item)}}
				{{$inv.Set $item (add ($inv.Get $item) $q)}}
			{{else}}
				{{$inv.Set $item $q}}
			{{end}}
			{{$bcp = sub $bcp $bc}}
			{{if le $bcp 0}}
				{{$inv.Del "[C] Biocomposant"}}
			{{else}}
				{{$inv.Set "[C] Biocomposant" $bcp}}
			{{end}}
			{{$cbp = sub $cbp $cb}}
			{{if le $cbp 0}}
				{{$inv.Del "[C] Cellule Bionotropique"}}
			{{else}}
				{{$inv.Set "[C] Cellule Bionotropique" $cbp}}
			{{end}}
			{{$sfp = sub $sfp $sf}}
			{{if le $sfp 0}}
				{{$inv.Del "[C] Substrat Ferreux"}}
			{{else}}
				{{$inv.Set "[C] Substrat Ferreux" $sfp}}
			{{end}}
			{{$lcp = sub $lcp $lc}}
			{{if le $lcp 0}}
				{{$inv.Del "[C] Liquide Cytomorphe"}}
			{{else}}
				{{$inv.Set "[C] Liquide Cytomorphe" $lcp}}
			{{end}}

		{{else if and (eq $choice "+VN") (not $univ)}}
			{{$all_bc := add $bc_vn $bcp}}
			{{$all_sf := add $sf_vn $sfp}}
			{{$all_cb := add $cb_vn $cbp}}
			{{$all_lc := add $lc_vn $lcp}}
			{{if and (ge $all_bc $bc) (ge $all_sf $sf) (ge $all_cb $cb) (ge $all_lc $lc)}}
				{{if ($inv.Get $item)}}
					{{$inv.Set $item (add ($inv.Get $item) $q)}}
				{{else}}
					{{$inv.Set $item $q}}
				{{end}}
				{{$reste_bc := sub $all_bc $bc}}
				{{$reste_sf := sub $all_sf $sf}}
				{{$reste_cb := sub $all_cb $cb}}
				{{$reste_lc := sub $all_lc $lc}}
				{{$reste_bc}}{{$all_bc}}{{$bc}}
				{{$inv.Del "[C] Biocomposant"}}
				{{$inv.Del "[C] Substrat Ferreux"}}
				{{$inv.Del "[C] Liquide Cytomorphe"}}
				{{$inv.Del "[C] Cellule Bionotropique"}}
				{{$log = joinStr " " $user "a fabriqué" $q $item "en utilisant son inventaire et celui du Nucleus. \n\n Il reste :\n ▫️" $reste_bc "Biocomposant \n ▫️" $reste_sf "Substrat Ferreux \n ▫️" $reste_lc "Liquide cytomorphe \n ▫️" $reste_cb "Cellule bionotropique" }}
				{{if le $reste_bc 0}}
					{{$compo.Del "biocomposant"}}
				{{else}}
					{{$compo.Set "biocomposant" $reste_bc}}
				{{end}}
				{{if le $reste_sf 0}}
					{{$compo.Del "ferreux"}}
				{{else}}
					{{$compo.Set "ferreux" $reste_sf}}
				{{end}}
				{{if le $reste_cb 0}}
					{{$compo.Del "bionotropique"}}
				{{else}}
					{{$compo.Set "bionotropique" $reste_cb}}
				{{end}}
				{{if le $reste_lc 0}}
					{{$compo.Del "cytomorphe"}}
				{{else}}
					{{$compo.Set "cytomorphe" $reste_lc}}
				{{end}}
			{{else}}
				Le vaisseau et {{$user}} n'ont pas les composants nécessaires pour faire {{$item}}
			{{end}}

		{{else if and (eq $choice "-inv") (not $univ)}}
			{{if and (ge $bc_vn $bc) (ge $sf_vn $sf) (ge $lc_vn $lc) (ge $cb_vn $cb)}}
				{{if ($inv.Get $item)}}
					{{$inv.Set $item (add ($compo.Get $item) $q)}}
				{{else}}
					{{$inv.Set $item $q}}
				{{end}}
				{{$bc_vn = sub $bc_vn $bc}}
				{{if le $bc_vn 0}}
					{{$compo.Del "biocomposant"}}
				{{else}}
					{{$compo.Set "biocomposant" $bc_vn}}
				{{end}}
				{{$cb_vn = sub $cb_vn $cb}}
				{{if le $cb_vn 0}}
					{{$compo.Del "bionotropique"}}
				{{else}}
					{{$compo.Set "bionotropique" $cb_vn}}
				{{end}}
				{{$sf_vn = sub $sf_vn $sf}}
				{{if le $sf_vn 0}}
					{{$compo.Del "ferreux"}}
				{{else}}
					{{$compo.Set "ferreux" $sf_vn}}
				{{end}}
				{{$lc_vn = sub $lc_vn $lc}}
				{{if le $lc_vn 0}}
					{{$compo.Del "cytomorphe"}}
				{{else}}
					{{$compo.Set "cytomorphe" $lc_vn}}
				{{end}}
				{{$log = joinStr " " $user "a fabriqué" $q $item "en utilisant l'inventaire du Nucleus.\n\n Il reste : \n ▫️ " $bc_vn "Biocomposant \n ▫️" $sf_vn "Substrat ferreux \n ▫️" $lc_vn "Liquide cytomorphe \n ▫️" $cb_vn "Cellule bionotropique."}}
			{{else}}
				Le Nucleus n'a pas les composants nécessaires pour fabriquer {{$q}} {{$item}}.
			{{end}}

		{{else if and $univ (eq $choice "+VN")}}
			{{if or (ge (mult $cu_vn $q) $q ) (ge (mult $cup $q) $q)}}
				{{if ($inv.Get $item)}}
					{{$inv.Set $item (add ($compo.Get $item) $q)}}
				{{else}}
					{{$inv.Set $item $q}}
				{{end}}
				{{$inv.Del "[C] Composant Universel"}}
				{{$compo.Set "universel" (sub $cu_vn (mult $cu_vn $q))}}
				{{if le ($compo.Get "universel") 0}}
					{{$compo.Del "universel"}}
				{{end}}
				{{$log = joinStr " " $user "a fabriqué" $q $item "en utilisant" $q "composant(s) universel(s)\n Il reste" ($compo.Get "universel") "Composant universel sur le vaisseau."}}
			{{else}}
				Le Nucleus et {{$user}} n'ont pas assez de composants universels pour fabriquer {{$q}} {{$item}}.
			{{end}}

		{{else if and $univ (eq $choice "-inv")}}
			{{if (ge (mult $cu_vn $q) $q)}}
				{{if ($inv.Get $item)}}
					{{$inv.Set $item (add ($compo.Get $item) $q)}}
				{{else}}
					{{$inv.Set $item $q}}
				{{end}}
				{{$compo.Set "universel" (sub $cu_vn (mult $cu_vn $q))}}
				{{if le ($compo.Get "universel") 0}}
					{{$compo.Del "universel"}}
				{{end}}
				{{$log = joinStr " " $user "a fabriqué" $q $item "en utilisant" $q "composants universels du vaisseau. \n Il reste" ($compo.Get "universel") "sur le vaisseau." }}
			{{else}}
				Le Nucleus n'a pas assez de composants universels pour fabriquer {{$q}} {{$item}}
			{{end}}

		{{else if $univ }}
			{{if (ge (mult $cup $q) $q)}}
				{{if ($inv.Get $item)}}
					{{$inv.Set $item (add ($inv.Get $item) $q)}}
				{{else}}
					{{$inv.Set $item $q}}
				{{end}}
				{{$inv.Set "universel" (sub $cup (mult $cup $q))}}
				{{if le ($inv.Get "universel") 0}}
					{{$inv.Del "universel"}}
				{{end}}
				{{$user}} a fabriqué {{$q}} {{$item}} en utilisant ses propres composants universels.
			{{else}}
				{{$user}} n'a pas assez de composants universels pour fabriquer {{$q}} {{$item}}
			{{end}}
		{{else}}
			{{$user}} ne peut pas fabriquer {{$q}} {{$item}} car ne possèdant pas les ressources nécessaires.
		{{end}}
	{{else}}
		L'objet ne possède pas de recette.
	{{end}}
{{else}}
	**Usage** : `$craft "objet" (q[1-100]) (+VN|-inv) (-bdg) (-cu) (#reroll)`
{{end}}
{{$userEco.Set "Inventory" $inv}}
{{dbSet $id "economy" $userEco}}
{{dbSet .Server.ID "compo" $compo}}
{{$chan := 735938256038002818}}
{{sendMessage $chan $log}}
{{deleteTrigger 1}}
