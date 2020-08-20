	{{/* Databases */}}
	{{$store := (dbGet .Server.ID "store").Value}}
	{{$serverEco := sdict}}
	{{with (dbGet .Server.ID "economy")}}
		{{$serverEco = sdict .Value}}
	{{end}}
	{{$symbol := ""}}
	{{if $serverEco.Get "symbol"}}
		{{$symbol = $serverEco.Get "symbol"}}
	{{end}}
	{{$desc := ""}}
	{{$name := "Vaisseau Marchand"}}
	{{$icon := "https://i.imgur.com/u80Pg6O.png"}}
	{{$thumb := "https://i.imgur.com/ACrlyqV.png"}}

	{{/* Store/Shop */}}
	{{$cslice := cslice}}
	{{$fields := cslice}}
	{{$footer := ""}}
	{{$items := sdict}}
	{{if or (eq $store "open") (hasRoleName "Staff") }}
		{{if ($serverEco.Get "Items")}}
			{{$items = sdict ($serverEco.Get "Items")}}
		{{end}}
		{{$desc = "La boutique est actuellement vide"}}
		{{if $items}}
			{{range $k,$v := $items}}
				{{$item := $k}}
				{{$bprice := $v.buyprice}}
				{{$sprice := $v.sellprice}}
				{{$stock := $v.stock}}
				{{$doc := $v.desc}}
				{{$rare := $v.rare}}
				{{if ne (str $stock) "♾️"}}
					{{$stock = $stock}}
					{{if eq $stock 0}}
						{{$stock = "**En rupture...**"}}
					{{end}}
				{{end}}
				{{$svente := $symbol }}
				{{if eq (toString $sprice) "Invendable"}}
					{{$svente = ""}}
				{{end}}
				{{$cslice = $cslice.Append (sdict "name" (title $item) "value" (print ":white_small_square: Achat : " $bprice " " $symbol  "\n :white_small_square: Vente: " $sprice " " $svente "\n :white_small_square: Rareté : " $rare "\n :white_small_square: Stock : " $stock "\n > " $doc ) "inline" false)}}
				{{$desc = "Hey ! Regarde tout ce que tu peux acheter !"}}
			{{end}}
		{{end}}
	{{else}}
		{{$desc = "Le magasin est actuellement fermé..."}}
		{{$name = "Vaisseau Nucleus"}}
		{{$icon = (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}
		{{$thumb = "https://i.imgur.com/FW1ZJDk.png"}}
	{{end}}

	{{/* hell starts */}}
	{{$page := ""}}
	{{$start := ""}}
	{{$stop := ""}}
	{{$end := ""}}
	{{if $cslice}}
		{{$page = "1"}}{{if .CmdArgs}}{{$page = toInt (index .CmdArgs 0)}}{{$page = toString $page}}{{end}}
		{{$start = (mult 10 (sub $page 1))}}
		{{$stop = (mult $page 10)}}
		{{$end = roundCeil (div (toFloat (len $cslice)) 10)}}
		{{$data := ""}}
		{{if ge $stop (len $cslice)}}
			{{$stop = (len $cslice)}}
		{{end}}
		{{if not (eq $page "0")}}
			{{if and (le $start $stop) (ge (len $cslice) $start) (le $stop (len $cslice))}}
				{{range (seq $start $stop)}}
					{{$fields = $fields.Append (index $cslice .)}}
				{{end}}
				{{$footer = (print "Page: " $page "/" $end)}}
			{{else}}
				{{$desc = "Cette page n'existe pas"}}
				{{$footer = (print "Page: " $page "/" $end)}}
			{{end}}
		{{else}}
			{{$desc = "Cette page n'existe pas"}}
			{{$footer = (print "Page: " $page "/" $end)}}
		{{end}}
	{{end}}
	{{/* hell ends */}}

	{{$id := sendMessageRetID nil (cembed "author" (sdict "name" $name "icon_url" $icon) "thumbnail" (sdict "url" $thumb) "color" 0x8CBAEF "description" $desc "fields" $fields "footer" (sdict "text" $footer))}}
	{{addMessageReactions nil $id "◀️" "▶️" "🗑️"}}
	{{deleteTrigger 1}}
