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
			{{printf "%T" $bmin}}
			{{printf "%T" $bmax}}
			{{printf "%T" $smin}}
			{{printf "%T" $smax}}
			{{$buyprice := randInt 0 100}}
			{{$sellprice := randInt 0 100}}
			{{$buyprice}}
		{{end}}
	{{end}}
{{end}}