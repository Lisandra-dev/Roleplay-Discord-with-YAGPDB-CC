{{/* Databases */}}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}


{{$desc := "Ton inventaire est vide ! Si le shop est ouvert, tu peux aller acheter des trucs !"}}
{{$footer := ""}}
{{$cslice := cslice}}
{{range $k,$v := $inv}}
	{{$cslice = $cslice.Append (printf " :white_small_square: ** %-10v **  : [%v]" $k $v)}}
{{end}}
{{$author := (joinStr " " "Inventaire de :" (title $user)}}

{{/* Pagination */}}
{{$del := false}}
{{ $action := .Reaction.Emoji.Name }} {{/* The action being ran */}}
{{ $validEmojis := cslice "▶️" "◀️" "🗑️" }} {{/* Valid emojis */}}
{{ $isValid := false }} {{/* Whether this is actually a valid embed / leaderboard embed */}}
{{ $page := 0 }} {{/* The current page */}}
{{ with and (eq .ReactionMessage.Author.ID 204255221017214977) .ReactionMessage.Embeds }} {{/* Checks for validity */}}
	 {{ $embed := structToSdict (index . 0) }}
	 {{ range $k, $v := $embed }}
		 {{- if eq (kindOf $v true) "struct" }}
		 {{- $embed.Set $k (structToSdict $v) }}
		 {{- end -}}
		{{ end }}
	{{if $embed.Author.Name}}
		{{$check := reFind `Inventaire` $embed.Author.Name}}
	{{ if and $check $embed.Footer}} {{/* More checks */}}
	{{ $page = toInt (reFind `\d+` $embed.Footer.Text) }} {{/* We presume that this is valid, and get the page num */}}
		{{ $isValid = true }} {{/* Yay, it is valid */}}
	{{else if eq $embed.Author.Name "Vaisseau Nucleus"}}
		{{$isValid = true}}
		{{$page = 1}}
	{{ end }}
{{ end }}

{{ if and (in $validEmojis $action) $isValid $page }} {{/* Even more checks for validity... */}}
		{{ deleteMessageReaction nil .ReactionMessage.ID .User.ID $action }}
	{{ if eq $action "▶️" }}
		{{ $page = add $page 1 }} {{/* Update page according to emoji */}}
	{{ else if eq $action "◀️"}}
		{{ $page = sub $page 1 }}
	{{else}}
		{{$del = true}}
		{{$page = 1}}
		{{deleteMessage nil .ReactionMessage.ID 1}}
	{{ end }}
{{if $cslice}}
{{/* hell starts */}}
	{{$end = roundCeil (div (toFloat (len $cslice)) 10)}}
	{{$footer = print "Page: " $page "/" $end}}
	{{$start := (mult 10 (sub $page 1))}}
	{{$stop := (mult $page 10)}}
	{{$data := ""}}
	{{if ge $stop (len $cslice)}}
		{{$stop = (len $cslice)}}
	{{end}}
	{{if not (eq $page "0")}}
		{{if and (le $start $stop) (ge (len $cslice) $start) (le $stop (len $cslice))}}
			{{range (seq $start $stop)}}
				{{$data = (print $data "\n" (index $cslice .))}}
			{{else}}
				{{$desc = "Cette page n'existe pas"}}
			{{end}}
			{{$desc = print "" $data ""}}
		{{else}}
			{{$desc = "Cette page n'existe pas"}}
		{{end}}
	{{end}}
{{/* hell ends */}}
{{end}}
{{if eq $del false}}
	{{editMessage nil .ReactionMessage.ID (cembed "author" (sdict "name" $author "icon_url" "https://i.imgur.com/iUmz9Gi.png") "color" 0x8CBAEF "description" $desc "footer" (sdict "text" $footer) )}}
{{else}}
	{{deleteMessage nil .ReactionMessage.ID 1}}
	{{end}}
{{end}}