{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}

{{/* Groupe dictionnaire */}}
{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Get player */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id:= .User.ID}}
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
{{$user = title $user}}
{{$img := "https://i.imgur.com/YeIsRmw.png"}}
{{/* get PA */}}
{{$pa := $groupe.Get (str $id)}}
{{if gt $pa 0}}
	{{if not (dbGet .User.ID "fusil2")}}
		{{dbSet .User.ID "fusil2" 0}}
	 	{{$incr := dbIncr .User.ID "fusil2" 1}}
		{{$y := (dbGet .User.ID "fusil2").Value}}
		{{$x := sub 12 $y}}
		{{if lt $y (toFloat 12)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges dans votre deuxième fusil !")}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre deuxième fusil est vide..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{else}}
	 	{{$incr := dbIncr .User.ID "fusil2" 1}}
		{{$y := (dbGet .User.ID "fusil2").Value}}
		{{$x := sub 12 $y}}
		{{if lt $y (toFloat 12)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges dans votre deuxième fusil !")}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre deuxième fusil est vide..."}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
		{{end}}
	{{end}}
{{else}}
	{{$embed := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"color"  0x6CAB8E
	"description" "Vous n'avez pas les PA pour faire cette action"}}
	{{ $idM := sendMessageRetID nil $embed }}
	{{deleteMessage nil $idM 30}}
{{end}}
