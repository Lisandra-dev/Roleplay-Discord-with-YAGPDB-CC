{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{if not $matches}}
	{{if (dbGet 0 "plasma")}}
	  {{if and (lt (dbGet 0 "plasma").Value (toFloat 11)) (gt (len .Message.Content) 15) }}
		 {{$incr := dbIncr 0 "plasma" 1}}
	  {{else if and (eq (dbGet 0 "plasma").Value (toFloat 11))  }}
		{{ $embed := cembed
		"description" "Bouclier plasma rechargé"}}
		{{ sendMessage nil $embed }}
	        {{dbDel 0 "plasma"}}
	    {{end}}
	{{end}}
{{else}}
{{end}}
