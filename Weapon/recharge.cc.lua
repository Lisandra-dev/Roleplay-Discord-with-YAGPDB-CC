{{/* You can change the word between quote, beware : some are the database key, if you change it you must change ALL key
Float are for the number of bullet, you can change it.  */}}

{{if not (dbGet .User.ID "recharge")}}
{{if .CmdArgs}}

{{/* Fist weapon : FUSIL [French word for shotgun]*/}}
  {{if eq (index .CmdArgs 0) "fusil"}}
    {{if gt (toFloat (dbGet .User.ID "fusil").Value) (toFloat 11)}}
      {{dbSet .User.ID "cf" 1}}
      {{dbSet 0 "run" (toString .Channel.ID)}}
	{{ $embed := cembed
	"description" (joinStr "" "Début du rechargement du fusil de " $.User.Mention)}}
	{{ sendMessage nil $embed }}
    {{else}}
	{{ $embed := cembed
	"description" "Votre fusil est chargé"}}
	{{ sendMessage nil $embed }}
    {{end}}


{{/* Second Weapon : PISTOLET*/}}

  {{else if eq (index .CmdArgs 0) "pistolet"}}
    {{if gt (toFloat (dbGet .User.ID "pistol").Value) (toFloat 7)}}
      {{dbSet .User.ID "cp" 1}}
      {{dbSet 0 "run" (toString .Channel.ID)}}
	{{ $embed := cembed
	"description" (joinStr "" "Début du rechargement du pistolet de " $.User.Mention)}}
	{{ sendMessage nil $embed }}    {{else}}
      {{ $embed := cembed
	"description" "Votre pistolet est chargé"}}
	{{ sendMessage nil $embed }}
    {{end}}

    {{/* CANON */}}

  {{else if eq (index .CmdArgs 0) "canon"}}
    {{if gt (toFloat (dbGet .User.ID "canon").Value) (toFloat 19)}}
      {{dbSet .User.ID "ca" 1}}
      {{dbSet 0 "run" (toString .Channel.ID)}}
      	{{ $embed := cembed
	"description" (joinStr "" "Début du rechargement du canon de " $.User.Mention)}}
	{{ sendMessage nil $embed }}
    {{else}}
	{{ $embed := cembed
	"description" "Votre canon est chargé"}}
	{{ sendMessage nil $embed }}
{{end}}


{{/* EMBED FOR SOME ERRORS */}}

{{else}}
{{ $embed := cembed
	"titre" "ERREUR"
	"description" "Erreur dans l'arme indiquée"}}
	{{ sendMessage nil $embed }}
  {{end}}

  {{else}}
    {{ $embed := cembed
	"titre" "ERREUR"
	"description" "Merci d'indiquer l'arme que vous rechargez."}}
	{{ sendMessage nil $embed }}
  {{end}}

{{else}}
    {{ $embed := cembed
	"titre" "ERREUR"
	"description" "Vous êtes déjà en cours de rechargement."}}
	{{ sendMessage nil $embed }}
{{end}}
