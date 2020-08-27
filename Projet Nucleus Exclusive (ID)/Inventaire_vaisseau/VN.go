{{$imp := ""}}
{{range $k, $v := (dbGet .Server.ID "implant")}}
	{{- $imp = (print  $imp "▫️" $k " : " $v "\n")}}
{{- end}}

{{$modu := ""}}
{{range $k, $v := (dbGet .Server.ID "module" )}}
	{{- $modu = (print $modu "▫️" $k " : " $v "\n")}}
{{- end}}

{{$weapon := ""}}
{{range $k,$v := (dbGet 0 "armelist").Value}}
	{{- $weapon = print $weapon "▫️" $k " : " $v "\n"}}
{{- end}}

{{$soin := ""}}
{{range $k, $v := (dbGet 0 "soin").Value}}
	{{- $soin = print $soin "▫️" $k " : " $v "\n"}}
{{- end}}

{{$compo := sdict}}
{{with (dbGet .Server.ID "compo")}}
	{{$compo = sdict .Value}}
{{end}}

{{$chargeur := sdict}}
{{with (dbGet 0 "chargeur_Multi")}}
	{{$chargeur = sdict .Value}}
{{end}}

{{$bioc := $compo.Get "biocomposant"}}
{{if not $bioc}}
	{{$bioc = 0}}
{{end}}
{{$cyto := $compo.Get "cytomorphe"}}
{{if not $cyto}}
	{{$cyto = 0}}
{{end}}
{{$biono := $compo.Get "bionotropique"}}
{{if not $biono}}
	{{$biono = 0}}
{{end}}
{{$ferreux := $compo.Get "ferreux"}}
{{if not $ferreux}}
	{{$ferreux = 0}}
{{end}}
{{$univ := $compo.Get "universel"}}
{{if not $univ}}
	{{$univ = 0}}
{{end}}
{{$fusil := $chargeur.Get "[CHARGEUR] Fusil"}}
{{if not $fusil}}
	{{$fusil = 0}}
{{end}}
{{$pistolet := $chargeur.Get "[CHARGEUR] Pistolet"}}
{{if not $pistolet}}
	{{$pistolet = 0}}
{{end}}
{{$canon := $chargeur.Get "[CHARGEUR] Canon"}}
{{if not $canon}}
	{{$canon = 0}}
{{end}}
{{$icon := (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}

{{$embed := cembed
"title" "Inventaire du Nucleus"
"thumbnail" (sdict "url" "https://i.imgur.com/Zt2aZl4.png")
"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)
"description" (joinStr " " "**Composant** \n ▫️ Biocomposant : " $bioc "\n▫️ Liquide cytomorphe :" $cyto "\n▫️ Cellule bionotropique :" $biono "\n ▫️ Substrat ferreux :" $ferreux "\n ▫️ Composant universel :" $univ "\n\n **Armes libres** :\n" $weapon "\n\n **Modules libres :**\n" $modu "\n\n **Soin** :" $soin " \n\n **Chargeurs disponibles : **\n ▫️ Pistolet : " $pistolet "\n ▫️ Fusil : " $fusil "\n ▫️ Canon : " $canon "\n\n **Implants disponibles :**" $imp )
"timestamp" currentTime
"color" 0x8CBAEF}}

{{editMessage 722755391498485800 736003604221001790 (complexMessageEdit "embed" $embed "content" "")}}
