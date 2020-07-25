{{/* You can change the word between quote, beware : some are the database key, if you change it you must change ALL key
Float are for the number of bullet, you can change it.  */}}

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

{{$img := "https://i.imgur.com/WNuPWCv.png"}}

{{if not (dbGet $id "recharge")}}
  {{if .CmdArgs}}
    {{if eq (index .CmdArgs 0) "fusil"}}
      {{if gt (toFloat (dbGet $id "fusil").Value) (toFloat 11)}}
        {{dbSet $id "cf" 1}}
  	     {{ $embed := cembed
					 	"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
					"description" "Début du rechargement de votre fusil..."}}
  	      {{ $idM := sendMessageRetID nil $embed }}

          {{deleteMessage nil $idM 30}}
      {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
			"description" "Votre fusil n'est pas déchargé."}}
  	  {{ $idM := sendMessageRetID nil $embed }}

      {{deleteMessage nil $idM 30}}
      {{end}}

			{{else if eq (index .CmdArgs 0) "fusil2"}}
	      {{if gt (toFloat (dbGet $id "fusil2").Value) (toFloat 11)}}
	        {{dbSet $id "cf2" 1}}
	  	     {{ $embed := cembed
						 	"author" (sdict "name" $user "icon_url" $img)
 							"color" 0x6CAB8E
						"description" "Début du rechargement de votre deuxième fusil..."}}
	  	      {{ $idM := sendMessageRetID nil $embed }}

	          {{deleteMessage nil $idM 30}}
	      {{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
				"description" "Votre deuxième fusil n'est pas déchargé."}}
		  	  {{ $idM := sendMessageRetID nil $embed }}

		      {{deleteMessage nil $idM 30}}
	      {{end}}

    {{else if eq (index .CmdArgs 0) "pistolet"}}
      {{if gt (toFloat (dbGet $id "pistol").Value) (toFloat 7)}}
        {{dbSet $id "cp" 1}}
  	    {{ $embed := cembed
						"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
					"description" "Début du rechargement de votre pistolet..."}}
  	    {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
				"description" "Votre pistolet n'est pas déchargé. "}}
  	    {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "canon"}}
      {{if gt (toFloat (dbGet $id "canon").Value) (toFloat 19)}}
        {{dbSet $id "ca" 1}}
        	{{ $embed := cembed
						"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
					"description" "Début du rechargement de votre canon..."}}
  	    {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
			"description" "Votre canon n'est pas déchargé."}}
    	  {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "pistolet2"}}
      {{if gt (toFloat (dbGet $id "pistol2").Value) (toFloat 7)}}
        {{dbSet $id "cp2" 1}}
				{{ $embed := cembed
 					"author" (sdict "name" $user "icon_url" $img)
 					"color" 0x6CAB8E
				"description" "Début du rechargement de votre deuxième pistolet..."}}
  	    {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{else}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
			"description" "Votre deuxième pistolet n'est pas déchargé. "}}
  	    {{ $idM := sendMessageRetID nil $embed }}

        {{deleteMessage nil $idM 30}}
      {{end}}
    {{else}}
      {{ $embed := cembed
  	     "titre" "ERREUR"
  	      "description" (joinStr "" "Erreur dans l'arme indiquée par " .User.Mention)}}
  	  {{ $idM := sendMessageRetID nil $embed }}

      {{deleteMessage nil $idM 30}}
    {{end}}
  {{else}}
      {{ $embed := cembed
      	"titre" "ERREUR"
      	"description" (joinStr "" "Merci d'indiquer l'arme que vous rechargez, " .User.Mention)}}
    	{{ $idM := sendMessageRetID nil $embed }}

      {{deleteMessage nil $idM 30}}
  {{end}}
{{else}}
    {{ $embed := cembed
  	"titre" "ERREUR"
  	"description" (joinStr "" .User.Mention ", vous êtes déjà en cours de rechargement.")}}
	{{ $idM := sendMessageRetID nil $embed }}

  {{deleteMessage nil $idM 30}}
{{end}}
{{deleteTrigger 1}}
