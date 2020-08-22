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
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}

{{$img := "https://i.imgur.com/YeIsRmw.png"}}
{{$pa := $groupe.Get (str $id)}}
{{if not $pa}}
	{{$groupe.Set (str $id) 4}}
	{{$pa = 4}}
{{end}}
{{if gt $pa 0}}
	{{if not (dbGet $id "pistol")}}
	  {{dbSet $id "pistol" 0}}
	  {{$incr := dbIncr $id "pistol" 1}}
	  {{$y := (dbGet $id "pistol").Value}}
	  {{$x := sub 8 $y}}
	    {{if lt $y (toFloat 7)}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
					"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/8 charges de pistolet !")}}
	      {{ $idM := sendMessageRetID nil $embed }}
	  		{{deleteMessage nil $idM 30}}
	  {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre pistolet est vide."}}
	    {{ $idM := sendMessageRetID nil $embed }}

	    {{deleteMessage nil $idM 30}}
	  {{end}}
	{{else}}
	  {{$incr := dbIncr $id "pistol" 1}}
	  {{$y := (dbGet $id "pistol").Value}}
	  {{$x := sub 8 $y}}
	  {{if lt $y (toFloat 8)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/8 charges de pistolet !")}}
	    {{ $idM := sendMessageRetID nil $embed }}

	    {{deleteMessage nil $idM 30}}
	  {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre pistolet est vide."}}
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