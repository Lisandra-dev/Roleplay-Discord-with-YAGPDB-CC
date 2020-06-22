{{/* You can change the word between quote, beware : some are the database key, if you change it you must change ALL key
Float are for the number of bullet, you can change it.  */}}

{{if not (dbGet .User.ID "recharge")}}
  {{if .CmdArgs}}
    {{if eq (index .CmdArgs 0) "fusil"}}
      {{if gt (toFloat (dbGet .User.ID "fusil").Value) (toFloat 11)}}
        {{dbSet .User.ID "cf" 1}}
        {{dbSet 0 "run" (toString .Channel.ID)}}
  	     {{ $embed := cembed
  	        "description" (joinStr "" "Début du rechargement du fusil de " $.User.Mention)}}
  	      {{ $id := sendMessageRetID nil $embed }}
          {{deleteMessage nil $id 30}}
      {{else}}
  	  {{ $embed := cembed
  	     "description" (joinStr "" "Le fusil de " .User.Mention " est chargé")}}
  	  {{ $id := sendMessageRetID nil $embed }}
      {{deleteMessage nil $id 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "pistolet"}}
      {{if gt (toFloat (dbGet .User.ID "pistol").Value) (toFloat 7)}}
        {{dbSet .User.ID "cp" 1}}
        {{dbSet 0 "run" (toString .Channel.ID)}}
  	    {{ $embed := cembed
  	       "description" (joinStr "" "Début du rechargement du pistolet de " $.User.Mention)}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{else}}
        {{ $embed := cembed
  	       "description" (joinStr "" "Le pistolet de " .User.Mention " est chargé")}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "canon"}}
      {{if gt (toFloat (dbGet .User.ID "canon").Value) (toFloat 19)}}
        {{dbSet .User.ID "ca" 1}}
        {{dbSet 0 "run" (toString .Channel.ID)}}
        	{{ $embed := cembed
  	         "description" (joinStr "" "Début du rechargement du canon de " $.User.Mention)}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{else}}
    	  {{ $embed := cembed
    	     "description" (joinStr "" "Le canon de" .User.Mention "est chargé")}}
    	  {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{end}}
    {{else}}
      {{ $embed := cembed
  	     "titre" "ERREUR"
  	      "description" (joinStr "" "Erreur dans l'arme indiquée par " .User.Mention)}}
  	  {{ $id := sendMessageRetID nil $embed }}
      {{deleteMessage nil $id 30}}
    {{end}}
  {{else}}
      {{ $embed := cembed
      	"titre" "ERREUR"
      	"description" (joinStr "" "Merci d'indiquer l'arme que vous rechargez, " .User.Mention)}}
    	{{ $id := sendMessageRetID nil $embed }}
      {{deleteMessage nil $id 30}}
  {{end}}
{{else}}
    {{ $embed := cembed
  	"titre" "ERREUR"
  	"description" (joinStr "" .User.Mention ", vous êtes déjà en cours de rechargement.")}}
	{{ $id := sendMessageRetID nil $embed }}
  {{deleteMessage nil $id 30}}
{{end}}
