{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Seuil & dictionnaire */}}
{{$seuil := sdict}}
{{with (dbGet .Server.ID "seuil")}}
	{{$seuil = sdict .Value}}
{{end}}

{{range $i, $j := $groupe}}

{{if le $j 0}}
	{{$j = 4}}
{{else if gt $j 2}}
	{{$j = 4}}
{{else if le $j 2}}
	{{$j = 6}}
{{end}}
	{{$groupe.Set $i $j}}
{{end}}
Tour +1 !
{{dbSet .Server.ID "groupe" $groupe}}
