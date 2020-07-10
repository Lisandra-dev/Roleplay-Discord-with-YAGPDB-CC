{{if not (dbGet 0 "plasma")}}
      {{dbSet 0 "plasma" 1}}
	{{ $embed := cembed
	"description" (joinStr "" "Début du rechargement bouclier plasma")}}
	{{ sendMessage nil $embed }}
    {{else}}
     	{{ $embed := cembed
	"description" "Le bouclier est déjà en cours de rechargement"}}
	{{ sendMessage nil $embed }}
    {{end}}
