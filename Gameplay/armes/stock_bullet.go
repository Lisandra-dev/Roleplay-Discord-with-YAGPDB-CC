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


{{$fusil := reFindAllSubmatches `fusil` .Message.Content}}
{{$fusil2 := reFindAllSubmatches `fusil2` .Message.Content}}
{{$pistolet := reFindAllSubmatches `pistolet` .Message.Content}}
{{$pistol2 := reFindAllSubmatches `pistolet2` .Message.Content }}
{{$canon := reFindAllSubmatches `canon` .Message.Content }}

{{$msgc:=""}}
{{$msgp:=""}}
{{$msgp2 := ""}}
{{$msgf := ""}}
{{$msgf2 := ""}}
{{$img := "https://i.imgur.com/YeIsRmw.png"}}

{{if $fusil}}
	{{if (dbGet $id "fusil")}}
		{{$y := (dbGet $id "fusil").Value}}
		{{$x := sub 12 $y}}
		{{if lt $y (toFloat 12)}}
	  	{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges de fusil !")}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre fusil est vide..."}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre fusil est complètement chargé. \n Vous avez 12 charges à disposition !"}}
		{{ $idM := sendMessageRetID nil $embed }}
	{{end}}

{{else if $fusil2 }}
	{{if (dbGet $id "fusil2")}}
		{{$y := (dbGet $id "fusil2").Value}}
		{{$x := sub 12 $y}}
		{{if lt $y (toFloat 12)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $x)) "/12 charges dans votre deuxième fusil !")}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre deuxième fusil est vide..."}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre deuxième fusil est complètement chargé.\n Vous avez 12 charges à disposition !"}}
		{{ $idM := sendMessageRetID nil $embed }}
	{{end}}

{{else if $pistolet}}
	{{if (dbGet $id "pistol")}}
		{{$a := (dbGet $id "pistol").Value}}
		{{$b := sub 8 $a}}
		{{if lt $a (toFloat 8)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $b)) "/8 charges de pistolet !")}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre pistolet est vide."}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{$embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
		"description" "Votre pistolet est complètement chargé.\n Vous avez 8 charge à disposition !"}}
		{{ $idM := sendMessageRetID nil $embed }}
	{{end}}

{{else if $pistol2}}
	{{if (dbGet $id "pistol2")}}
		{{$a := (dbGet $id "pistol2").Value}}
		{{$b := sub 8 $a}}
		{{if lt $a (toFloat 8)}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
					"description" (joinStr "" "Il vous reste " (toString (toInt $b)) "/8 charges dans votre deuxième pistolet !")}}
				{{ $idM := sendMessageRetID nil $embed }}
			{{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
					"description" "Votre deuxième pistolet est vide."}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{$embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre deuxième pistolet est complètement chargé.\n Vous avez 8 charge à disposition !"}}
		{{ $idM := sendMessageRetID nil $embed }}
	{{end}}

{{else if $canon}}
	{{if (dbGet $id "canon")}}
		{{$c := (dbGet $id "canon").Value}}
		{{$d := sub 20 $c}}
		{{if lt $c (toFloat 20)}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" (joinStr "" "Il vous reste " (toString (toInt $d)) "/20 charges de canon !")}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre canon est vide !"}}
			{{ $idM := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{ $embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
			"color" 0x6CAB8E
			"description" "Votre canon est complètement chargé.\n Vous avez 20 charge à disposition !"}}
		{{ $idM := sendMessageRetID nil $embed }}
	{{end}}

{{else}}
		{{if (dbGet $id "canon")}}
			{{$canon := (dbGet $id "canon").Value}}
			{{$canon := (toFloat (sub 20 $canon))}}
			{{if le $canon (toFloat 20)}}
				{{$msgc = (joinStr "" "Il vous reste " (toString (toInt $canon)) "/20 charge dans votre canon !")}}
			{{else}}
				{{$msgc = "Votre canon est vide"}}
			{{end}}
		{{else}}
			{{$msgc = "Votre canon est complètement chargé ! Vous avez 20 charges disponibles actuellement."}}
		{{end}}

		{{if (dbGet $id "fusil")}}
			{{$fusil := (dbGet $id "fusil").Value}}
			{{$fusil := (toFloat (sub 12 $fusil))}}
			{{if le $fusil (toFloat 12)}}
				{{$msgf = (joinStr "" "Il vous reste " (toString (toInt $fusil)) "/12 charge dans votre fusil")}}
			{{else}}
				{{$msgf = "Votre fusil est vide..."}}
			{{end}}
		{{else}}
			{{$msgf = "Votre fusil est complètement chargé ! Vous avez 12 charges disponibles actuellement. "}}
		{{end}}

		{{if (dbGet $id "fusil2")}}
			{{$fusil2 := (dbGet $id "fusil2").Value}}
			{{$fusil2 := (toFloat (sub 12 $fusil2 ))}}
			{{if le $fusil2 (toFloat 12)}}
				{{$msgf2 = (joinStr "" "Il vous reste " (toString (toInt $fusil2)) "/12 charge dans votre deuxième fusil")}}
			{{else}}
				{{$msgf2 = "Votre deuxième fusil est vide..."}}
			{{end}}
		{{else}}
			{{$msgf2 = "Votre deuxième fusil est complètement chargé ! Vous avez 12 charges disponibles actuellement. "}}
		{{end}}

		{{if (dbGet $id "pistol")}}
			{{$pistolet := (dbGet $id "pistol").Value}}
			{{$pistolet := (toFloat (sub 8 $pistolet))}}
			{{if le $pistolet (toFloat 8)}}
				{{$msgp = (joinStr "" "Il vous reste " (toString (toInt $pistolet )) "/8 dans votre pistolet")}}
			{{else}}
				{{$msgp = "Votre pistolet est vide..."}}
			{{end}}
		{{else}}
			{{$msgp = "Votre pistolet est complètement chargé ! Vous avez 8 charges à disposition."}}
		{{end}}
		{{if (dbGet $id "pistol2")}}
			{{$pistol2 := (dbGet $id "pistol2").Value}}
			{{$pistol2 := (toFloat (sub 8 $pistol2))}}
			{{if le $pistol2 (toFloat 8)}}
				{{$msgp2 = (joinStr "" "Il vous reste " (toString (toInt $pistol2 )) "/8 dans votre deuxième pistolet")}}
			{{else}}
				{{$msgp2 = "Votre deuxième pistolet est vide..."}}
			{{end}}
		{{else}}
			{{$msgp2 = "Votre deuxième pistolet est complètement chargé ! Vous avez 8 charges à disposition."}}
		{{end}}

		{{$embed := cembed
			"author" (sdict "name" $user)
			"title" "Charges restantes"
			"description" (joinStr "" ":white_small_square: **Canon** : " $msgc "\n:white_small_square: **Fusil** : " $msgf "\n:white_small_square: **Fusil 2** : " $msgf2 "\n:white_small_square: **Pistolet** : " $msgp "\n:white_small_square: **Pistolet 2** : " $msgp2 "\n")
			"color" (0x38A9BD)
			"thumbnail" (sdict "url" $img)}}
			{{sendMessage nil $embed}}
{{end}}
{{deleteTrigger 1}}
{{deleteResponse 15}}
