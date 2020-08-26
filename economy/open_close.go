{{$matches := reFind `\$(open|close)` .Message.Content}}

{{/* User variable */}}
{{$author := "Sola-UI"}}
{{$chan := "701395027896565810"}}
{{$avatar := "https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png"}}
{{$msg_ouvert := "Un vaisseau marchand vous contacte sur la ligne de communication. Il vous ouvre sa boutique dans un canal privé, n'hésitez pas à aller leur répondre."}}
{{$msg_close := "Le vaisseau marchand vient de partir."}}

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

{{if eq $matches "$open"}}
	{{dbSet .Server.ID "store" "open"}}
	{{if $items}}
		{{range $k, $v := $items}}
			{{$i := sdict ($items.Get $k)}}
			{{$bmin := $i.Get "bmin"}}
			{{$bmax := $i.Get "bmax"}}
			{{$smin := $i.Get "smin"}}
			{{$smax := $i.Get "smax"}}
			{{$buyprice := $i.Get "buyprice"}}
			{{$sellprice := $i.Get "sellprice"}}
			{{if ne (toInt $bmax) 0}}
				{{$buyprice = randInt (toInt $bmin) (toInt $bmax)}}
			{{end}}
			{{if eq (toInt $smax) 0}}
				{{$sellprice = "Invendable"}}
			{{else}}
				{{$sellprice = randInt (toInt $smin) (toInt $smax)}}
			{{end}}

			{{$stock := "♾️"}}
			{{$rstock := randInt 0 100}}
			{{if le $rstock 10}}
				{{$stock = 0}}
			{{else if and (ge $rstock 11) (le $rstock 30)}}
				{{$stock  = randInt 1 5}}
			{{end}}
			{{if eq ($i.Get "rare") "⭐" "⭐⭐" }}
				{{if le $rstock 10}}
					{{$stock = randInt 1 10}}
				{{else if and (ge $rstock 11) (le $rstock 30)}}
					{{$stock  = randInt 10 20}}
				{{end}}
			{{end}}
			{{$i.Set "buyprice" $buyprice}}
			{{$i.Set "sellprice" $sellprice}}
			{{$i.Set "stock" $stock}}
			{{$items.Set $k $i}}
		{{end}}
	{{end}}

	{{$serverEco.Set "Items" $items}}
	{{dbSet .Server.ID "economy" $serverEco}}

	{{if not .CmdArgs}}
		{{$embed := cembed
		"author" (sdict "name" $author "icon_url" $avatar)
		"description" $msg_ouvert
		"color" 0xD5B882}}
		{{sendMessage $chan $embed}}
	{{else}}
		{{$Message := "**Ouverture & Randomisation silencieuse**"}}
		{{sendMessage nil $Message}}
	{{end}}

{{else}}
	**FERMÉ**
	{{dbSet .Server.ID "store" "close"}}
	{{$embed := cembed
	"author" (sdict "name" $author "icon_url" $avatar)
	"description" $msg_close
	"color" 0xD5B882}}
	{{sendMessage $chan $embed}}
{{end}}

{{deleteTrigger 1}}
