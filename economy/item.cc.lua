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
		{{$name = (index . 1)}}
		{{if ge (len .) 3}}
			{{$flag2 = (reFind `\-(price|sell|stock|sii|desc)` (index . 2))}}
		{{end}}
	{{end}}
{{end}}

{{if eq $flag "-add" }}
	{{$bprice := ""}}
	{{$sprice := ""}}
	{{$stock := ""}}
	{{$desc := ""}}
	{{$sii := true}}
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
							{{$desc = (index . 6)}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}

	{{if and $name $bprice $stock $sprice}}
		{{$items.Set $name (sdict "buyprice" $bprice "sellprice" $sprice "stock" $stock "sii" $sii "desc" $desc)}}
		{{$serverEco.Set "Items" $items}}
		**Created/Edited Item**
		Nom : `{{$name}}`
		Prix d'achat : `{{$bprice}}`
		Prix de vente : `{{$sprice}}`
		Quantité : `{{$stock}}`
		SII: `{{$sii}}`
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
				Description : `{{.desc}}`
			{{end}}
		{{else}}
			Cet objet n'existe pas !
			**Usage** : `$item -info <name>`
		{{end}}
	{{end}}

{{else if eq $flag "-edit"}}
	{{if $name}}
			{{if $name}}
				{{if ($items.Get ( $name))}}
					{{$i := ($items.Get ( $name))}}
					{{if eq flag2 "-price"}}
						{{$buyprice := toInt (index .CmdArgs 3)}}
						{{$items.Set $name $i}}
						{{$i.Set "buyprice" $buyprice}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-sell"}}
						{{$sprice := toInt (index .CmdArgs 3)}}
						{{if eq $sprice 0}}
							{{$sprice = "Invendable"}}
						{{end}}
						{{$items.Set $name $i}}
						{{$i.Set "sellprice" $sprice}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-stock"}}
						{{$stock := toInt (index .CmdArgs 3)}}
						{{if eq $stock "inf"}}
							{{$stock = "♾️"}}
						{{end}}
						{{$items.Set $name $i}}
						{{$i.Set "stock" $stock}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-sii"}}
						{{$sii := (index .CmdArgs 2)}}
						{{$items.Set $name $i}}
						{{$i.Set "sii" $sii}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-desc"}}
						{{$desc := (index .CmdArgs 3)}}
						{{$items.Set $name $i}}
						{{$i.Set "desc" $desc}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-rbuy"}}
						{{$buyprice := randInt (toInt (index .CmdArgs 3)) (toInt (index .CmdArgs 4))}}
						{{$items.Set $name $i}}
						{{$i.Set "buyprice" $buyprice}}
						{{$serverEco.Set "Items" $items}}
					{{else if eq flag2 "-rsell"}}
						{{$sellprice := randInt (toInt (index .CmdArgs 3)) (toInt (index .CmdArgs 4))}}
						{{$items.Set $name $i}}
						{{$i.Set "sellprice" $sprice}}
						{{$serverEco.Set "Items" $items}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
	**Erreur : ** Le produit n'existe pas.
	**Usage : ** `$item -edit "nom" -(price|sell|stock|sii|desc|rbuy|rsell) <value>`
		> Pour rbuy & rsell : A besoin des bornes pour l'aléatoire.
		> pour price & sell : Prix
		> Stock : Montant ou `inf`
		> SII : True / False (signifie si l'objet apparaît dans les inventaires ou non)
		> Description : Chaîne de caractère entre `"`.
	{{end}}
	{{with $i}}
		**Item Info**
		Nom : `{{ $name}}`
		Prix d'achat : `{{.buyprice}}`
		Prix de vente : `{{.sellprice}}`
		Quantité : `{{.stock}}`
		SII: `{{.sii}}`
		Description : `{{.desc}}`
	{{end}}
{{else}}
	**Usage** : `$item -(info|add|delete|edit)`
{{end}}

{{/* Database Updates */}}
{{dbSet .Server.ID "economy" $serverEco}}
