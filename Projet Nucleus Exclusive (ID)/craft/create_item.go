{{/* Dictionnaire d'Item */}}
	{{/* Recette */ }}
{{$recipe := sdict}}
{{with (dbGet 0 "recipe")}}
	{{$recipe = sdict .Value}}
{{end}}

	{{/* Inventaire du Nucleus */}}

{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}

{{$armes := sdict}}
{{with (dbGet 0 "armelist")}}
	{{$armes = sdict .Value}}
{{end}}

{{$modulist := (dbGet .Server.ID "module")}}
{{$module := cslice}}
{{if $modulist}}
	{{$module = (cslice).AppendSlice $modulist.Value}}
{{else}}
	{{$module = cslice}}
{{end}}

{{$implist := dbGet 0 "implant"}}
{{$implant := cslice}}
{{if $implist}}
	{{$implant = (cslice).AppendSlice $implist.Value}}
{{else}}
	{{$implant = cslice}}
{{end}}

{{$chargeur := sdict}}
{{with (dbGet 0 "chargeur_Multi")}}
	{{$chargeur = sdict .Value}}
{{end}}

{{$soin := sdict}}
{{with (dbGet 0 "soin")}}
	{{$soin = sdict .Value}}
{{end}}

{{/* Personnage */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
	{{dbSetExpire $id "rerollName" $name 3600}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}
{{$user = title $user}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}

{{/* Function */}}
{{if .CmdArgs}}
	{{$item := title (index .CmdArgs 1)}}
	{}
