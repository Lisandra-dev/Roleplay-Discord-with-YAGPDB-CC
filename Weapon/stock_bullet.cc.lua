{{$user := .Member.Nick}}
{{if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$msgc:=""}}
{{$msgp:=""}}
{{$msgp2 := ""}}
{{$msgf := ""}}
{{$msgf2 := ""}}

{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "fusil"}}
		{{if (dbGet .User.ID "fusil")}}
		{{$y := (dbGet .User.ID "fusil").Value}}
		{{$x := sub 12 $y}}
		{{if lt $y (toFloat 12)}}
      {{ $embed := cembed
      "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) "/12 charges de fusil !")}}
      {{ $id := sendMessageRetID nil $embed }}

		{{else}}
      {{ $embed := cembed
        "description" (joinStr "" .User.Mention ", votre fusil est vide...")}}
      {{ $id := sendMessageRetID nil $embed }}

		{{end}}
	{{else}}
		{{ $embed := cembed
      "description" (joinStr "" .User.Mention ", votre fusil est complètement chargé. Vous avez 12 charges à disposition !")}}
		{{ $id := sendMessageRetID nil $embed }}

	{{end}}

	{{else if (dbGet .User.ID "fusil2")}}
		{{if (dbGet .User.ID "fusil2")}}
			{{$y := (dbGet .User.ID "fusil2").Value}}
			{{$x := sub 12 $y}}
			{{if lt $y (toFloat 12)}}
				{{ $embed := cembed
					"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) "/12 charges dans votre deuxième fusil !")}}
				{{ $id := sendMessageRetID nil $embed }}

			{{else}}
				{{ $embed := cembed
					"description" (joinStr "" .User.Mention ", votre deuxième fusil est vide...")}}
				{{ $id := sendMessageRetID nil $embed }}

			{{end}}
		{{else}}
			{{ $embed := cembed
				"description" (joinStr "" .User.Mention ", votre deuxième fusil est complètement chargé. Vous avez 12 charges à disposition !")}}
			{{ $id := sendMessageRetID nil $embed }}
			{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet"}}
		{{if (dbGet .User.ID "pistol")}}
		{{$a := (dbGet .User.ID "pistol").Value}}
		{{$b := sub 8 $a}}
		{{if lt $a (toFloat 7)}}
			{{ $embed := cembed
				"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $b)) "/8 charges de pistolet !")}}
			{{ $id := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"description" (joinStr "" .User.Mention ", votre pistolet est vide.")}}
			{{ $id := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{$embed := cembed
			"description" (joinStr "" .User.Mention ", votre pistolet est complètement chargé. Vous avez 8 charge à disposition !")}}
		{{ $id := sendMessageRetID nil $embed }}
	{{end}}

	{{else if eq (index .CmdArgs 0) "pistolet2"}}
		{{if (dbGet .User.ID "pistol2")}}
			{{$a := (dbGet .User.ID "pistol2").Value}}
			{{$b := sub 8 $a}}
			{{if lt $a (toFloat 7)}}
				{{ $embed := cembed
					"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $b)) "/8 charges dans votre deuxième pistolet !")}}
				{{ $id := sendMessageRetID nil $embed }}
			{{else}}
				{{ $embed := cembed
					"description" (joinStr "" .User.Mention ", votre deuxième pistolet est vide.")}}
				{{ $id := sendMessageRetID nil $embed }}
			{{end}}
		{{else}}
			{{$embed := cembed
				"description" (joinStr "" .User.Mention ", votre deuxième pistolet est complètement chargé. Vous avez 8 charge à disposition !")}}
			{{ $id := sendMessageRetID nil $embed }}
		{{end}}
	{{else if eq (index .CmdArgs 0) "canon"}}
		{{if (dbGet .User.ID "canon")}}
		{{$c := (dbGet .User.ID "canon").Value}}
		{{$d := sub 20 $c}}
		{{if lt $c (toFloat 20)}}
			{{ $embed := cembed
				"description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $d)) "/20 charges de canon !")}}
			{{ $id := sendMessageRetID nil $embed }}
		{{else}}
			{{ $embed := cembed
				"description" (joinStr "" .User.Mention ", votre canon est vide !")}}
			{{ $id := sendMessageRetID nil $embed }}
		{{end}}
	{{else}}
		{{ $embed := cembed
			"description" (joinStr "" .User.Mention ",votre canon est complètement chargé. Vous avez 20 charge à disposition !")}}
		{{ $id := sendMessageRetID nil $embed }}
	{{end}}

	{{else}}
		**Usage : ** `$stock <canon|pistolet|fusil>`
	{{end}}
{{else}}
	{{if (dbGet .User.ID "canon")}}
		{{$canon := (dbGet .User.ID "canon").Value}}
		{{$canon := (toFloat (sub 20 $canon))}}
		{{if lt $canon (toFloat 20)}}
			{{$msgc = (joinStr "" "Il vous reste " (toString (toInt $canon)) "/20 charge dans votre canon !")}}
		{{else}}
			{{$msgc = "Votre canon est vide"}}
		{{end}}
	{{else}}
		{{$msgc = "Votre canon est complètement chargé ! Vous avez 20 charges disponibles actuellement."}}
	{{end}}

	{{if (dbGet .User.ID "fusil")}}
		{{$fusil := (dbGet .User.ID "fusil").Value}}
		{{$fusil := (toFloat (sub 12 $fusil))}}
		{{if lt $fusil (toFloat 12)}}
			{{$msgf = (joinStr "" "Il vous reste " (toString (toInt ($fusil)) "/12 charge dans votre fusil"))}}
		{{else}}
			{{$msgf = "Votre fusil est vide..."}}
		{{end}}
	{{else}}
		{{$msgf = "Votre fusil est complètement chargé ! Vous avez 12 charges disponibles actuellement. "}}
	{{end}}
	{{if (dbGet .User.ID "fusil2")}}
		{{$fusil2 := (dbGet .User.ID "fusil2").Value}}
		{{$fusil2 := (toFloat (sub 12 $fusil2 ))}}
		{{if lt $fusil2 (toFloat 12)}}
			{{$msgf2 = (joinStr "" "Il vous reste " (toString (toInt $fusil2)) "/12 charge dans votre deuxième fusil")}}
		{{else}}
			{{$msgf2 = "Votre deuxième fusil est vide..."}}
		{{end}}
	{{else}}
		{{$msgf2 = "Votre deuxième fusil est complètement chargé ! Vous avez 12 charges disponibles actuellement. "}}
	{{end}}

	{{if (dbGet .User.ID "pistol")}}
		{{$pistolet := dbGet .User.ID "pistol"}}
		{{$pistolet := (toFloat (sub 8 $pistolet))}}
		{{if lt $pistolet (toFloat 7)}}
			{{$msgp = (joinStr "" "Il vous reste " (toString (toInt $pistolet )) "/8 dans votre pistolet")}}
		{{else}}
			{{$msgp = "Votre pistolet est vide..."}}
		{{end}}
	{{else}}
		{{$msgp = "Votre pistolet est complètement chargé ! Vous avez 8 charges à disposition."}}
	{{end}}
	{{if (dbGet .User.ID "pistol2")}}
		{{$pistol2 := dbGet .User.ID "pistol2"}}
		{{$pistol2 := (toFloat (sub 8 $pistol2))}}
		{{if lt $pistol2 (toFloat 7)}}
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
		"description" (joinStr " " ":white_small_square: **Canon** :" $msgc "\n:white_small_square: **Fusil** :" $msgf "\n:white_small_square: **Fusil 2** :" $msgf2 "\n:white_small_square: **Pistolet** : " $msgp "\n:white_small_square: **Pistolet 2** :" $msgp2 "\n")
		"color" (0x38A9BD)}}
		{{sendMessage nil $embed}}
{{end}}
{{deleteTrigger 1}}
{{deleteResponse 15}}
