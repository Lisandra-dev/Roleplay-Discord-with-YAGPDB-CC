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
	{{if ($items.Get ( "[C] Biocomposant"))}}
		{{$i := sdict ($items.Get ( "[C] Biocomposant"))}}
		{{$bp := randInt 75 125}}
		{{$sp := randInt 50 60}}
		{{$items.Set "[C] Biocomposant" $i}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
	{{end}}
	
	{{if ($items.Get ( "[C] Substrat Ferreux"))}}
		{{$i := sdict ($items.Get ( "[C] Substrat Ferreux"))}}
		{{$bp := randInt 125 175}}
		{{$sp := randInt 80 100}}
		{{$items.Set "[C] Substrat Ferreux" $i}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
	{{end}}
	
	{{if ($items.Get ( "[C] Liquide Cytomorphe"))}}
		{{$i := sdict ($items.Get ( "[C] Liquide Cytomorphe"))}}
		{{$bp := randInt 250 330}}
		{{$sp := randInt 100 150}}
		{{$items.Set "[C] Liquide Cytomorphe" $i}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
	{{end}}
	
	{{if ($items.Get ( "[C] Cellule Bionotropique"))}}
		{{$i := sdict ($items.Get ( "[C] Cellule Bionotropique"))}}
		{{$bp := randInt 300 380}}
		{{$sp := randInt 150 200}}
		{{$items.Set "[C] Cellule Bionotropique" $i}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
	{{end}}
	
	{{if ($items.Get ( "[C] Composant Universel"))}}
		{{$i := sdict ($items.Get ( "[C] Composant Universel"))}}
		{{$bp := randInt 1200 1400}}
		{{$sp := randInt 500 600}}
		{{$items.Set "[C] Composant Universel" $i}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
	{{end}}
	
	{{if ($items.Get ( "[E] Sac À Dos"))}}
		{{$i := sdict ($items.Get ( "[E] Sac À Dos"))}}
		{{$bp := randInt 1000 1200}}
		{{$sp := randInt 500 550}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[E] Sac À Dos" $i}}

	{{end}}
	
	{{if ($items.Get ( "[E] Sacoche"))}}
		{{$i := sdict ($items.Get ( "[E] Sacoche"))}}
		{{$bp := randInt 500 800}}
		{{$sp := randInt 300 350}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[E] Sacoche" $i}}
	{{end}}
	
	{{if ($items.Get ("[E] Sacoche Ceinture"))}}
		{{$i := sdict ($items.Get ("[E] Sacoche Ceinture"))}}
		{{$bp := randInt 300 350}}
		{{$sp := randInt 150 200}}
		{{$i.Set "sellprice" $sp}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[E] Sacoche Ceinture" $i}}
	{{else}}
		not found
	{{end}}
	
	{{if ($items.Get ( "[M] Désimplantation"))}}
		{{$i := sdict ($items.Get ( "[M] Désimplantation"))}}
		{{$bp := randInt 800 1000}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[M] Désimplantation" $i}}
	{{end}}
	
		{{if ($items.Get ( "[M] Remodelage Humain"))}}
		{{$i := sdict ($items.Get ( "[M] Remodelage Humain"))}}
		{{$bp := randInt 2000 3000}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[M] Remodelage Humain" $i}}
	{{end}}
	
		{{if ($items.Get ( "[M] Régénération Humaine"))}}
		{{$i := sdict ($items.Get ( "[M] Régénération Humaine"))}}
		{{$bp := randInt 8000 10000}}
		{{$i.Set "buyprice" $bp}}
		{{$items.Set "[M] Régénération Humaine" $i}}
	{{end}}
	
	{{$serverEco.Set "Items" $items}}
	
	{{dbSet .Server.ID "economy" $serverEco}}
	
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