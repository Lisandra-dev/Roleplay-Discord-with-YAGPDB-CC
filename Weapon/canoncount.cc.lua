
{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{if $name}}
	{{$user = $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$img := "https://www.pixenli.com/image/Dkm5e-vR"}}

{{if not (dbGet .User.ID "canon")}}
  {{dbSet .User.ID "canon" 0}}
  {{$incr := dbIncr .User.ID "canon" 1}}
  {{$y := (dbGet .User.ID "canon").Value}}
  {{$x := sub 20 $y}}
  {{if lt $y (toFloat 20)}}
		{{ $embed := cembed
		"author" (sdict "name" $user "icon_url" $img)
		"color" 0x6CAB8E
		"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/20 charges de canon !")}}
    {{ $id := sendMessageRetID nil $embed }}
  	{{deleteMessage nil $id 30}}
  {{else}}
		{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre canon est vide !"}}
  		{{deleteMessage nil $id 30}}
  {{end}}
{{else}}
  {{$incr := dbIncr .User.ID "canon" 1}}
  {{$y := (dbGet .User.ID "canon").Value}}
  {{$x := sub 20 $y}}
  {{if lt $y (toFloat 20)}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/20 charges de canon !")}}
    {{ $id := sendMessageRetID nil $embed }}
    {{deleteMessage nil $id 30}}
  {{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre canon est vide !"}}
    {{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
  {{end}}
{{end}}
