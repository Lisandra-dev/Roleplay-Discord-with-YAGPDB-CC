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
{{if $cslice}}
{{/* hell starts */}}
	{{$page := "1"}}
	{{if .CmdArgs}}
		{{$index := 0}}
		{{if ge (len .CmdArgs) 2}}
			{{$index = 1}}
		{{end}}
		{{$page = or (toInt (index .CmdArgs $index)) 1}}
		{{$page = toString $page}}
	{{end}}
	{{$footer = print "Page: " $page}}
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
{{$author := (joinStr " " "Inventaire de :" (title $user)}}

{{sendMessage nil (cembed "author" (sdict "name" $author "icon_url" "https://i.imgur.com/iUmz9Gi.png") "color" 0x8CBAEF "description" $desc "footer" (sdict "text" $footer) )}}
{{deleteTrigger 1}}
