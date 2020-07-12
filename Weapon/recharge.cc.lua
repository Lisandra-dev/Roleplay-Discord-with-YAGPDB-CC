{{/* You can change the word between quote, beware : some are the database key, if you change it you must change ALL key
Float are for the number of bullet, you can change it.  */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{if $name}}
	{{$user = $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$img := "https://www.pixenli.com/image/xDMop79G"}}

{{if not (dbGet .User.ID "recharge")}}
  {{if .CmdArgs}}
    {{if eq (index .CmdArgs 0) "fusil"}}
      {{if gt (toFloat (dbGet .User.ID "fusil").Value) (toFloat 11)}}
        {{dbSet .User.ID "cf" 1}}
  	     {{ $embed := cembed
					 	"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
						"description" "Début du rechargement de votre fusil"}}
  	      {{ $id := sendMessageRetID nil $embed }}
          {{deleteMessage nil $id 30}}
      {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre fusil est complètement rechargé. \n Vous avez 12 charges à disposition !"}}
  	  {{ $id := sendMessageRetID nil $embed }}
      {{deleteMessage nil $id 30}}
      {{end}}

			{{else if eq (index .CmdArgs 0) "fusil2"}}
	      {{if gt (toFloat (dbGet .User.ID "fusil2").Value) (toFloat 11)}}
	        {{dbSet .User.ID "cf2" 1}}
	  	     {{ $embed := cembed
						 	"author" (sdict "name" $user "icon_url" $img)
 							"color" 0x6CAB8E
							"description" "Début du rechargement de votre deuxième fusil"}}
	  	      {{ $id := sendMessageRetID nil $embed }}
	          {{deleteMessage nil $id 30}}
	      {{else}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
					"description" "Votre deuxième fusil est complètement rechargé.\n Vous avez 12 charges à disposition !"}}
		  	  {{ $id := sendMessageRetID nil $embed }}
		      {{deleteMessage nil $id 30}}
	      {{end}}

    {{else if eq (index .CmdArgs 0) "pistolet"}}
      {{if gt (toFloat (dbGet .User.ID "pistol").Value) (toFloat 7)}}
        {{dbSet .User.ID "cp" 1}}
  	    {{ $embed := cembed
						"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
  	       "description" "Début du rechargement de votre pistolet"}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{else}}
				{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"color" 0x6CAB8E
					"description" "Votre pistolet est complètement rechargé.\n Vous avez 8 charge à disposition !"}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "canon"}}
      {{if gt (toFloat (dbGet .User.ID "canon").Value) (toFloat 19)}}
        {{dbSet .User.ID "ca" 1}}
        	{{ $embed := cembed
						"author" (sdict "name" $user "icon_url" $img)
						"color" 0x6CAB8E
  	        "description" "Début du rechargement de votre canon"}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{else}}
			{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre canon est complètement rechargé.\n Vous avez 20 charge à disposition !"}}
    	  {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{end}}
    {{else if eq (index .CmdArgs 0) "pistolet2"}}
      {{if gt (toFloat (dbGet .User.ID "pistol2").Value) (toFloat 7)}}
        {{dbSet .User.ID "cp2" 1}}
				{{ $embed := cembed
 					"author" (sdict "name" $user "icon_url" $img)
 					"color" 0x6CAB8E
  	      "description" "Début du rechargement de votre deuxième pistolet"}}
  	    {{ $id := sendMessageRetID nil $embed }}
        {{deleteMessage nil $id 30}}
      {{else}}
			{{$embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"color" 0x6CAB8E
				"description" "Votre deuxième pistolet est complètement rechargé.\n Vous avez 8 charge à disposition !"}}
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
{{deleteTrigger 1}}
