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
