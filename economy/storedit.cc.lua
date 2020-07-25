{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{$items := sdict}}
{{if ($serverEco.Get "Items")}}
	{{$items = sdict ($serverEco.Get "Items")}}
{{end}}
{{$i:= sdict}}

{{$name := (index .CmdArgs 0)}}

{{if ($items.Get ( $name))}}
	{{$i = sdict ($items.Get ( $name))}}
	{{if .CmdArgs}}
		{{if eq (index .CmdArgs 1) "-price"}}
			{{$buyprice := toInt (index .CmdArgs 2)}}
			{{$items.Set $name $i}}
			{{$i.Set "buyprice" $buyprice}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-sell"}}
			{{$sprice := toInt (index .CmdArgs 2)}}
			{{if eq $sprice 0}}
				{{$sprice = "Invendable"}}
			{{end}}
			{{$items.Set $name $i}}
			{{$i.Set "sellprice" $sprice}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-stock"}}
			{{$stock := toInt (index .CmdArgs 2)}}
			{{if eq $stock "inf"}}
				{{$stock = "♾️"}}
			{{end}}
			{{$items.Set $name $i}}
			{{$i.Set "stock" $stock}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-sii"}}
			{{$sii := (index .CmdArgs 2)}}
			{{$items.Set $name $i}}
			{{$i.Set "sii" $sii}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-desc"}}
			{{$desc := (index .CmdArgs 2)}}
			{{$items.Set $name $i}}
			{{$i.Set "desc" $desc}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-rbuy"}}
			{{$buyprice := randInt (toInt (index .CmdArgs 2)) (toInt (index .CmdArgs 3))}}
			{{$items.Set $name $i}}
			{{$i.Set "buyprice" $buyprice}}
			{{$serverEco.Set "Items" $items}}
		{{else if eq (index .CmdArgs 1) "-rsell"}}
			{{$sellprice := randInt (toInt (index .CmdArgs 2)) (toInt (index .CmdArgs 3))}}
			{{$items.Set $name $i}}
			{{$i.Set "sellprice" $sprice}}
			{{$serverEco.Set "Items" $items}}
		{{end}}
	{{end}}
{{else}}
	Le produit n'existe pas !
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

{{dbSet .Server.ID "economy" $serverEco}}
