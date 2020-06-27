{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{$arg := (dbGet .User.ID "comp").Value}}
{{if not $matches}}
	{{if (dbGet .User.ID $arg)}}
		{{if and (lt (dbGet .User.ID $arg).Value (toFloat 3)) (gt (len .Message.Content) 15) }}
				{{$incr := dbIncr .User.ID $arg 1}}
			{{else if (eq (dbGet .User.ID $arg).Value (toFloat 3))  }}
				{{ $embed := cembed
					"description" (joinStr "" "Votre compétence " $arg " est de nouveau utilisable")}}
				{{ $id := sendMessageRetID nil $embed }}
				{{deleteMessage nil $id 30}}
				{{dbDel .User.ID $arg}}
		{{end}}
	{{end}}
{{else}}
{{end}}
