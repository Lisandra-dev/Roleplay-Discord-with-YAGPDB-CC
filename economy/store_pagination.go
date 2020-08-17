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
			{{if ne (str $stock) "♾️"}}
				{{$stock = $stock}}
			{{end}}
			{{$svente := $symbol }}
			{{if eq (toString $sprice) "Invendable"}}
				{{$svente = ""}}
			{{end}}
			{{$cslice = $cslice.Append (sdict "name" (title $item) "value" (print ":white_small_square: Achat : " $bprice " " $symbol  "\n :white_small_square: Vente: " $sprice " " $svente "\n :white_small_square: Stock : " $stock "\n > " $doc ) "inline" false)}}
			{{$desc = "Hey ! Regarde tout ce que tu peux acheter !"}}
		{{end}}
	{{end}}
{{else}}
	{{$desc = "Le magasin est actuellement fermé..."}}
	{{$name = "Vaisseau Nucleus"}}
	{{$icon = (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}
	{{$thumb = "https://i.imgur.com/FW1ZJDk.png"}}
{{end}}




{{ $action := .Reaction.Emoji.Name }} {{/* The action being ran */}}
{{ $validEmojis := cslice "▶️" "◀️" }} {{/* Valid emojis */}}
{{ $isValid := false }} {{/* Whether this is actually a valid embed / leaderboard embed */}}
{{ $page := 0 }} {{/* The current page */}}
{{ with and (eq .ReactionMessage.Author.ID 204255221017214977) .ReactionMessage.Embeds }} {{/* Checks for validity */}}
	 {{ $embed := structToSdict (index . 0) }}
	 {{ range $k, $v := $embed }}
		 {{- if eq (kindOf $v true) "struct" }}
		 {{- $embed.Set $k (structToSdict $v) }}
		 {{- end -}}
		{{ end }}
	{{ if and (eq $embed.Author.Name "Vaisseau Marchand") $embed.Footer}} {{/* More checks */}}
		{{ $page = reFind `\d+` $embed.Footer.Text }} {{/* We presume that this is valid, and get the page num */}}
		{{ $isValid = true }} {{/* Yay, it is valid */}}
	{{ end }}
{{ end }}

{{ if and (in $validEmojis $action) $isValid $page }} {{/* Even more checks for validity... */}}
		{{ deleteMessageReaction nil .ReactionMessage.ID .User.ID $action }}
	{{ if eq $action "▶️" }}
		{{ $page = add $page 1 }} {{/* Update page according to emoji */}}
	{{ else }}
		{{ $page = sub $page 1 }}
	{{ end }}

	{{$start := ""}}
	{{$stop := ""}}
	{{$end := ""}}
	{{if $cslice}}
		{{$start = (mult 10 (sub $page 1))}}
		{{$stop = (mult $page 10)}}
		{{$end = roundCeil (div (toFloat (len $cslice)) 10)}}
		{{$data := ""}}
		{{if ge $stop (len $cslice)}}
			{{$stop = (len $cslice)}}
		{{end}}
		{{if not (eq $page 0)}}
			{{if and (le $start $stop) (ge (len $cslice) $start) (le $stop (len $cslice))}}
				{{range (seq $start $stop)}}
					{{$fields = $fields.Append (index $cslice .)}}
				{{end}}
				{{$footer = (print "Page: " $page "/" $end)}}
			{{else}}
				{{$desc = "Il n'y a rien ici"}}
				{{$footer = (print "Page: " $page)}}
			{{end}}
		{{else}}
			{{$desc = "Il n'y a rien ici"}}
			{{$footer = (print "Page: " $page)}}
		{{end}}
	{{end}}
	{{editMessage nil .ReactionMessage.ID (cembed "author" (sdict "name" $name "icon_url" $icon) "thumbnail" (sdict "url" $thumb) "color" 0x8CBAEF "description" $desc "fields" $fields "footer" (sdict "text" $footer))}}
{{end}}
