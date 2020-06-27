{{/* Float corresspond to the number of message what must be send */}}
{{/* len .Message.Content is for the length of each message : the message must be that length for +1 the counter.  */}}

{{ $matches := (index (reFindAllSubmatches `^\((.*)\)|(^\$(.*))|(^\!\!(.+))|((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})` .Message.Content)) }}
{{if not $matches }}
	{{/* Fusil : "cf" and "fusil" key : 6 message of 15 characters */}}

	{{if (dbGet .User.ID "cf")}}
	  {{if and (lt (dbGet .User.ID "cf").Value (toFloat 6)) (gt (len .Message.Content) 15)}} {{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
	    {{$incr := dbIncr .User.ID "recharge" 1}}
		  {{$incr := dbIncr .User.ID "cf" 1}}
	  {{else if and (eq (dbGet .User.ID "cf").Value (toFloat 6))}}{{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
		  {{ $embed := cembed
		      "description" (joinStr "" "Fusil de " $.User.Mention " rechargé.")}}
		  {{ $id := sendMessageRetID nil $embed }}
	    {{deleteMessage nil $id 30}}
	    {{dbDel .User.ID "recharge"}}
		  {{dbDel .User.ID "cf"}}
		  {{dbSet .User.ID "fusil" 0}}
	  {{end}}

	  {{/* CANON : "ca" key and "canon" key : 11 messages of 15 characters */}}
	{{else if (dbGet .User.ID "ca")}}
	  {{if and (lt (dbGet .User.ID "ca").Value (toFloat 11)) (gt (len .Message.Content) 15) }} {{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
	    {{$incr := dbIncr .User.ID "recharge" 1}}
		  {{$incr := dbIncr .User.ID "ca" 1}}
	  {{else if and (eq (dbGet .User.ID "ca").Value (toFloat 11))}} {{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
		  {{ $embed := cembed
		     "description" (joinStr "" "Canon de " $.User.Mention " rechargé.")}}
		  {{ $id := sendMessageRetID nil $embed }}
	    {{deleteMessage nil $id 30}}
	    {{dbDel .User.ID "recharge"}}
		  {{dbDel .User.ID "ca"}}
		  {{dbSet .User.ID "canon" 0}}
	  {{end}}

	  {{/* PISTOLET : "cp" key and "pistol" key : 4 messages of 15 characters */}}

	{{else if (dbGet .User.ID "cp")}}
	  {{if and (lt (dbGet .User.ID "cp").Value (toFloat 4)) (gt (len .Message.Content) 15)}}{{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
	    {{$incr := dbIncr .User.ID "recharge" 1}}
		  {{$incr := dbIncr .User.ID "cp" 1}}
	  {{else if and (eq (dbGet .User.ID "cp").Value (toFloat 4))}} {{/* Lock à un channel : (eq (toString .Channel.ID) (dbGet 0 "run").Value)*/}}
	    {{ $embed := cembed
		     "description" (joinStr "" "Pistolet de " $.User.Mention " rechargé.")}}
		  {{ $id := sendMessageRetID nil $embed }}
	    {{deleteMessage nil $id 30}}
	    {{dbDel .User.ID "recharge"}}
		  {{dbDel .User.ID "cp"}}
		  {{dbSet .User.ID "pistol" 0}}
	  {{end}}
	{{end}}
{{else}}
{{end}}
