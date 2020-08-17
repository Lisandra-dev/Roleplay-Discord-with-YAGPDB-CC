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
	{{with .CmdArgs}}
		{{if ge (len .) 3}}
			{{$bprice = toInt (index . 2)}}
			{{if ge (len .) 4}}
				{{if eq (index . 3) "none"}}
					{{$sprice = "Invendable"}}
				{{else}}
					{{$sprice = toInt (index . 3)}}
				{{end}}
				{{if ge (len .) 5}}
					{{if ne (index . 4) "inf"}}
						{{$stock = toInt (index . 4)}}
					{{else}}
						{{$stock = "♾️"}}
					{{end}}
					{{if ge (len .) 6}}
						{{if (reFind `true|false` (index . 5))}}
							{{$sii = (reFind `true|false` (index . 5))}}
							{{if eq $sii "true"}}
								{{$sii = true}}
							{{else}}
								{{$sii = false}}
							{{end}}
						{{end}}
						{{if ge (len .) 7}}
							{{range seq 1 (toInt (index . 6))}}
								{{$rare = print $rare "⭐"}}
							{{end}}
						{{end}}
						{{if ge (len .) 8}}
							{{$desc = (index . 7)}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}

	{{if and $name $bprice $stock $sprice}}
		{{$items.Set $name (sdict "buyprice" $bprice "sellprice" $sprice "stock" $stock "sii" $sii "rare" $rare "desc" $desc)}}
		{{$serverEco.Set "Items" $items}}
		**Created/Edited Item**
		Nom : `{{$name}}`
		Prix d'achat : `{{$bprice}}`
		Prix de vente : `{{$sprice}}`
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
	{{if $name}}
		{{if ($items.Get ( $name))}}
			{{$i := ($items.Get ( $name))}}
			{{with $i}}
				**Item Info**
				Nom : `{{ $name}}`
				Prix d'achat : `{{.buyprice}}`
				Prix de vente : `{{.sellprice}}`
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
				{{$buyprice := toInt (index .CmdArgs 3)}}
				{{$items.Set $name $i}}
				{{$i.Set "buyprice" $buyprice}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-sell"}}
				{{$sprice := toInt (index .CmdArgs 3)}}
				{{if eq $sprice 0}}
					{{$sprice = "Invendable"}}
				{{end}}
				{{$items.Set $name $i}}
				{{$i.Set "sellprice" $sprice}}
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
			{{else if eq $flag2 "-rbuy"}}
				{{$buyprice := randInt (toInt (index .CmdArgs 3)) (toInt (index .CmdArgs 4))}}
				{{$items.Set $name $i}}
				{{$i.Set "buyprice" $buyprice}}
				{{$serverEco.Set "Items" $items}}
			{{else if eq $flag2 "-rsell"}}
				{{$sellprice := randInt (toInt (index .CmdArgs 3)) (toInt (index .CmdArgs 4))}}
				{{$sellprice}}
				{{$items.Set $name $i}}
				{{$i.Set "sellprice" $sellprice}}
`				{{$serverEco.Set "Items" $items}}
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
				Nom : `{{ $rename}}`
				Prix d'achat : `{{.buyprice}}`
				Prix de vente : `{{.sellprice}}`
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
