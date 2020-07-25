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
		Il manque des arguments !
	{{end}}

{{else if eq $flag "-delete"}}
	{{if $name}}
		{{if ($items.Get ( $name))}}
			{{$items.Del ( $name)}}
			{{$serverEco.Set "Items" $items}}
			**Objet supprimé : ** `{{( $name)}}`
		{{else}}
			Cet objet n'existe pas.
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
		{{end}}
	{{end}}

{{else if eq $flag "-edit"}}
	{{$buyprice := ""}}
	{{if eq $flag2 "-price"}}
		{{$buyprice = (toInt (index .CmdArgs 3)) }}
	{{end}}
	{{if $name}}
		{{if ($items.Get ( $name))}}
			{{$i := ($items.Get ( $name))}}
			{{$.Get "buyprice"}}
			{{with $i}}
				{{$.Set "buyprice" $buyprice}}
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
		{{end}}
	{{end}}

{{else}}
	Argument error
{{end}}

{{/* Database Updates */}}
{{dbSet .Server.ID "economy" $serverEco}}
