{{/* Float corresspond to the number of message what must be send */}}
{{/* len .Message.Content is for the length of each message : the message must be that length for +1 the counter.  */}}

{{/* Fusil : "cf" and "fusil" key : 4 message of 15 characters */}}
{{/* Fusil : "cf2" and "fusil2" key : 4 message of 15 characters */}}

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

{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{if not $matches}}
{{if (dbGet $id "cf")}}
	{{if and (lt (dbGet $id "cf").Value (toFloat 4)) (gt (len .Message.Content) 15)}}
		{{$incr := dbIncr $id "recharge_cf" 1}}
		{{$incr := dbIncr $id "cf" 1}}
	{{else if and (eq (dbGet $id "cf").Value (toFloat 4))}}
		{{ $embed := cembed
			"description" (joinStr "" "Fusil de " $.User.Mention " rechargé.")}}
		{{ $idM := sendMessageRetID nil $embed }}
		{{deleteMessage nil $idM 30}}
		{{dbDel $id "recharge_cf"}}
		{{dbDel $id "cf"}}
		{{dbDel $id "fusil"}}
	{{end}}

	{{else if (dbGet $id "cf2")}}
		{{if and (lt (dbGet $id "cf2").Value (toFloat 4)) (gt (len .Message.Content) 15)}}
			{{$incr := dbIncr $id "recharge_cf2" 1}}
			{{$incr := dbIncr $id "cf2" 1}}
		{{else if and (eq (dbGet $id "cf2").Value (toFloat 4))}}
			{{ $embed := cembed
				"description" (joinStr "" "Deuxième fusil de " $.User.Mention " rechargé.")}}
			{{ $idM := sendMessageRetID nil $embed }}
			{{deleteMessage nil $idM 30}}
			{{dbDel $id "recharge_cf2"}}
			{{dbDel $id "cf2"}}
			{{dbDel $id "fusil2"}}
		{{end}}

  {{/* CANON : "ca" key and "canon" key : 7 messages of 15 characters */}}

{{else if (dbGet $id "ca")}}
	{{if and (lt (dbGet $id "ca").Value (toFloat 7)) (gt (len .Message.Content) 15) }}
		{{$incr := dbIncr $id "recharge_ca" 1}}
		{{$incr := dbIncr $id "ca" 1}}
	{{else if and (eq (dbGet $id "ca").Value (toFloat 7))}}
	  	{{ $embed := cembed
	     		"description" (joinStr "" "Canon de " $.User.Mention " rechargé.")}}
	  	{{ $idM := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $idM 30}}
   		{{dbDel $id "recharge_ca"}}
		{{dbDel $id "ca"}}
		{{dbDel $id "canon"}}
  	{{end}}

  {{/* PISTOLET : "cp" key and "pistol" key : 2 messages of 15 characters */}}
	{{/* PISTOLET 2 : "cp2" key and "pistol2" key : 2 messages of 15 characters */}}

{{else if (dbGet $id "cp")}}
	{{if and (lt (dbGet $id "cp").Value (toFloat 2)) (gt (len .Message.Content) 15)}}
	{{$incr := dbIncr $id "recharge_cp" 1}}
	{{$incr := dbIncr $id "cp" 1}}
	{{else if and (eq (dbGet $id "cp").Value (toFloat 2))}}
    		{{ $embed := cembed
	     		"description" (joinStr "" "Pistolet de " $.User.Mention " rechargé.")}}
	  	{{ $idM := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $idM 30}}
		{{dbDel $id "recharge_cp"}}
		{{dbDel $id "cp"}}
		{{dbDel $id "pistol"}}
	{{end}}

{{else if (dbGet $id "cp2")}}
	{{if and (lt (dbGet $id "cp2").Value (toFloat 2)) (gt (len .Message.Content) 15)}}
	{{$incr := dbIncr $id "recharge_cp2" 1}}
	{{$incr := dbIncr $id "cp2" 1}}
	{{else if and (eq (dbGet $id "cp2").Value (toFloat 2))}}
    		{{ $embed := cembed
	     		"description" (joinStr "" "Deuxième pistolet de " $.User.Mention " rechargé.")}}
	  	{{ $idM := sendMessageRetID nil $embed }}
    		{{deleteMessage nil $idM 30}}
		{{dbDel $id "recharge_cp2"}}
		{{dbDel $id "cp2"}}
		{{dbDel $id "pistol2"}}
	{{end}}
{{end}}
{{else}}
{{end}}
