{{$matches := reFind `\$(open|close)` .Message.Content}}

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
{{$name := ""}}

{{if eq $matches "$open"}}
	{{dbSet .Server.ID "store" "open"}}
**Génération des items**

	{{if $items}}
		{{range $k, $v := $items}}
			{{$i := sdict ($items.Get $k)}}
			{{$bmin := $i.Get "bmin"}}
			{{$bmax := $i.Get "bmax"}}
			{{$smin := $i.Get "smin"}}
			{{$smax := $i.Get "smax"}}
				{{$buyprice := randInt (toInt $bmin) (toInt $bmax)}}
				{{$sellprice := $i.Get "sellprice"}}
				{{if eq (toInt $smax) 0}}
					{{$sellprice = "Invendable"}}
				{{else}}
					{{$sellprice = randInt (toInt $smin) (toInt $smax)}}
				{{end}}
								
				{{$stock := "♾️"}}
				{{$rstock := randInt 0 100}}
				{{if le $rstock 10}}
					{{$stock = 0}}
				{{else if and (ge $rstock 11) (le $rstock 20)}}
					{{$stock  = randInt 1 5}}
				{{end}}
				{{$i.Set "buyprice" $buyprice}}
				{{$i.Set "sellprice" $sellprice}}
				{{$i.Set "stock" $stock}}
				{{$items.Set $k $i}}
			{{end}}
		{{end}}
		
		{{$serverEco.Set "Items" $items}}
		{{dbSet .Server.ID "economy" $serverEco}}
		
		{{range $k,$v := $items}}
			{{$k}} : {{$v.buyprice}} - {{$v.sellprice}} - {{$v.stock}}
			{{end}}
		
	
		{{$embed := cembed
		"author" (sdict "name" "Sola-UI" "icon_url" "https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png")
		"description" "Un vaisseau marchand vous contacte sur la ligne de communication. Il vous ouvre sa boutique dans un canal privé, n'hésitez pas à aller leur répondre."
		"color" 0xD5B882}}
		{{sendMessage 701395027896565810 $embed}}
	
	{{else}}
	**FERMÉ**
		{{dbSet .Server.ID "store" "close"}}
		{{$embed := cembed
		"author" (sdict "name" "Sola-UI" "icon_url" "https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png")
		"description" "Le vaisseau marchand vient de partir."
		"color" 0xD5B882}}
		{{sendMessage 701395027896565810 $embed}}
	{{end}}
{{deleteTrigger 1}}