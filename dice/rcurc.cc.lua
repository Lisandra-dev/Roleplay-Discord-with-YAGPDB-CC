{{$rc:=reFindAllSubmatches `^\$rc` .Message.Content}}
{{$rcc:=reFindAllSubmatches `^\$rcc` .Message.Content}}
{{if .CmdArgs}}
	{{$y := (toFloat (index .CmdArgs 0))}}
	{{if eq $y (toFloat 0)}}
			Merci de rentrer un nombre.
	{{else}}
		{{if $rc}}
			{{$x:= (toInt (mult 1.4 $y))}}
			Vous avez maintenant un bonus de {{$x}}% sur cette compétence non-offensive.
		{{else}}
			{{$x:= (toInt (mult 1.8 $y))}}
			Vous avez maintenant un bonus de {{$x}}% sur cette compétence non-offensive.
		{{end}}
	{{end}}
{{else}}
	**Usage** : `$(rcc|rc) valeur`
{{end}}
{{deleteTrigger 1}}
{{deleteResponse 30}}
