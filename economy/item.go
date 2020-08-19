{{/* Databases */}}
{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Managing items */}}

{{$items := sdict}}
{{if ($serverEco.Get "Items")}}
	{{$items = sdict ($serverEco.Get "Items")}}
{{end}}
{{$flag := ""}}
{{$flag2 := ""}}
{{$name := ""}}
{{with .CmdArgs}}
	{{if ge (len .) 2}}
		{{$flag = (reFind `\-(add|edit|delete|info)` (index . 0))}}
		{{$name = title (index . 1)}}
		{{if ge (len .) 3}}
			{{$flag2 = (reFind `\-(price|sell|stock|sii|desc|rare|rbuy|rsell|rename)` (index . 2))}}
		{{end}}
	{{end}}
{{end}}

{{if eq $flag "-add" }}
	{{$bprice := ""}}
	{{$sprice := ""}}
	{{$stock := ""}}
	{{$desc := ""}}
	{{$sii := true}}
	{{$rare := "⭐"}}
	{{$smin := $sprice}}
	{{$smax := $sprice}}
	{{$bmin := $bprice}}
	{{$bmax := $bprice}}
	{{with .CmdArgs}}
		{{if ge (len .) 4}}
			{{$bmin = toInt (index . 2)}}
			{{$bmax = toInt (index . 3)}}
			{{if eq $bmax 0}}
				{{$buyprice = $bmin}}
			{{else}}
				{{$buyprice = toInt 0}}
			{{end}}
			{{if ge (len .) 5}}
				{{if eq (index . 4) "none" "0" }}
					{{$sprice = "Invendable"}}
					{{$smin = 0}}
					{{$smax = 0}}
				{{else}}
					{{$smin = toInt (index . 4)}}
					{{$smax = toInt (index . 5)}}
					{{$sprice = 0}}
				{{end}}
				{{if ge (len .) 7}}
					{{if ne (index . 6) "inf"}}
						{{$stock = toInt (index . 6)}}
					{{else}}
						{{$stock = "♾️"}}
					{{end}}
					{{if ge (len .) 8}}
						{{if (reFind `true|false` (index . 7))}}
							{{$sii = (reFind `true|false` (index . 7))}}
							{{if eq $sii "true"}}
								{{$sii = true}}
							{{else}}
								{{$sii = false}}
							{{end}}
						{{end}}
						{{if ge (len .) 9}}
							{{range seq 1 (toInt (index . 8))}}
								{{$rare = print $rare "⭐"}}
							{{end}}
						{{end}}
						{{if ge (len .) 10}}
							{{$desc = (index . 9)}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}

	{{if and $name $bprice $stock $sprice}}
		{{$items.Set $name (sdict "buyprice" $bprice "bmin" $bmin "bmax" $bmax "sellprice" $sprice "smin" $smin "smax" $smax "stock" $stock "sii" $sii "rare" $rare "desc" $desc)}}
		{{$serverEco.Set "Items" $items}}
		**Created/Edited Item**
		Nom : `{{$name}}`
		Prix d'achat min : `{{$bmin}}`
		Prix d'achat max : `{{$bmax}}`
		Prix de vente min : `{{$smin}}`
		Prix de vente max : `{{$smax}}` 
		Quantité : `{{$stock}}`
		SII: `{{$sii}}`
		Rareté : `{{$rare}}`
		Description : `{{$desc}}`
	{{else}}
	**Usage** : `$item -add <name> price (none|price) (number|inf) (true|false) description`
	{{end}}

{{else if eq $flag "-delete"}}
	{{if $name}}
		{{if ($items.Get ( $name))}}
			{{$items.Del ( $name)}}
			{{$serverEco.Set "Items" $items}}
			**Objet supprimé : ** `{{( $name)}}`
		{{else}}
			Cet objet n'existe pas.
			**Usage** : `$item -delete <name>`
		{{end}}
	{{end}}

{{else if eq $flag "-info"}}
	{{if (reFind `\- all` .Message.Content)}}
		{{range $k, $v := $items}}
			**Nom** : {{$k}}
			**Prix d'achat** : {{$v.buyprice}}
			**Prix de vente** : {{$v.sellprice}}
			**Stock** : {{$v.stock}}
			**Rareté** : {{$v.rare}}
			**Description** : {{$v.desc}}
		{{end}}		
	{{else if not (reFind `\-all` .Message.Content)}}
		{{if ($items.Get ( $name))}}
			{{$i := ($items.Get ( $name))}}
			{{with $i}}
				**Item Info**
				Nom : `{{ $name}}`
				Prix d'achat min : `{{.bmin}}`
				Prix d'achat max : `{{.bmax}}`
				Prix de vente min : `{{.smin}}`
				Prix de vente max : `{{.smax}}` 
				Quantité : `{{.stock}}`
				SII: `{{.sii}}`
				Rareté : `{{.rare}}`
				Description : `{{.desc}}`
			{{end}}
		{{else}}
			Cet objet n'existe pas !
			**Usage** : `$item -info <name>`
		{{end}}
	{{end}}

{{else if eq $flag "-edit"}}
	{{if $name}}
		{{if ($items.Get ( $name))}}
			{{$i := sdict ($items.Get ( $name))}}
			{{$rename := $name}}
			{{if eq $flag2 "-price"}}
				{{$bmin := toInt (index .CmdArgs 3)}}
				{{$bmax := toInt (index .CmdArgs 4)}}		
				{{$items.Set $name $i}}
				{{$i.Set "bmax" $bmax}}
				{{$i.Set "bmin" $bmin}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-sell"}}
				{{$sprice := $i.Get "sellprice"}}
				{{$smin := toInt (index .CmdArgs 3)}}
				{{$smax := toInt (index .CmdArgs 4)}}
				{{if eq $smax 0}}
					{{$sprice = "Invendable"}}
					{{$smin = 0}}
					{{$smax = 0}}
				{{else}}
					{{$sprice = $smax}}
				{{end}}
				{{$items.Set $name $i}}
				{{$i.Set "sellprice" $sprice}}
				{{$i.Set "smin" $smin}}
				{{$i.Set "smax" $smax}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-stock"}}
				{{$stock := toInt (index .CmdArgs 3)}}
				{{if eq $stock 0}}
					{{$stock = "♾️"}}
				{{end}}
				{{$items.Set $name $i}}
				{{$i.Set "stock" $stock}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-sii"}}
				{{$sii := (index .CmdArgs 3)}}
				{{$items.Set $name $i}}
				{{$i.Set "sii" $sii}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-desc"}}
				{{$desc := (index .CmdArgs 3)}}
				{{$items.Set $name $i}}
				{{$i.Set "desc" $desc}}
				{{$serverEco.Set "Items" $items}}
`			{{else if eq $flag2 "-rare"}}
				{{$rare := "⭐"}}
				{{range seq 1 (toInt (index .CmdArgs 3))}}
					{{$rare = print $rare "⭐"}}
				{{end}}
				{{$items.Set $name $i}}
				{{$i.Set "rare" $rare}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-rename"}}
				{{$rename = title (index .CmdArgs 3)}}
				{{$items.Del $name}}
				{{$i.Set "name" $rename}}
				{{$items.Set $rename $i}}
				{{$serverEco.Set "Items" $items}}
			{{end}}
			{{with $i}}
				**Item Info**
				Nom : `{{ $name}}`
				Prix d'achat min : `{{.bmin}}`
				Prix d'achat max : `{{.bmax}}`
				Prix de vente min : `{{.smin}}`
				Prix de vente max : `{{.smax}}` 
				Quantité : `{{.stock}}`
				SII: `{{.sii}}`
				Rareté : `{{.rare}}`
				Description : `{{.desc}}`
			{{end}}
		{{else}}
		**Erreur : ** Le produit n'existe pas.
		**Usage : ** `$item -edit "nom" -(price|sell|stock|sii|rare|desc|rbuy|rsell|rename) <value>`
			> Pour rbuy & rsell : A besoin des bornes pour l'aléatoire.
			> pour price & sell : Prix
			> Stock : Montant ou `inf`
			> SII : True / False (signifie si l'objet apparaît dans les inventaires ou non)
			> Rareté : Nombre
			> Rename : Nouveau nom, si plusieurs mots, entre `"`
			> Description : Chaîne de caractère entre `"`.
		{{end}}
	{{end}}

{{else}}
	**Usage** : `$item -(info|add|delete|edit)`
{{end}}

{{/* Database Updates */}}
{{dbSet .Server.ID "economy" $serverEco}}
