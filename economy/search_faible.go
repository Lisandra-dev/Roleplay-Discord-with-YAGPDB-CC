{{$ri :="**Rixolam** : Restaure 10 PV" }}
{{$ban := "**Bandage** : Réduit de 1 tour un malus "}}
{{$eu := "**Eufyrusant** : Augmente les capacités PSI pour 1 tour."}}
{{$im := "**Implant temporaire** : Implant sur une caractéristique pendant 3 tours (ou cycle, si hors combat)."}}
{{$so := "**Soma** : Réduit de 1 tours les cooldown d'un module ou d'un PSI."}}
{{$eu := "**Eufyrusant** : Augmente une compétence de 8% pour 1 tour."}}
{{$ro := "**Sirop de betapropyl** : Divise par deux les effets des malus."}}
{{$gr := "**Grenade Nécrotique** : 5% bonus projectile"}}
{{$fn := "**Liquide antifongique : Annule bouclier mush durant 1Tour**"}}
{{$th := "**Gaz anesthésiant** : Endort la cible si jet de karma raté"}}
{{$sng := "**Sang Ethérique** : 7% bonus sur une attaque"}}
{{$cr := "**Huile carotonixique** : Malus sur toutes les caractéristiques pendant 1 tour"}}
{{$di := "**Huile digestive** : Divise par 2 les effets d'une armure sur 3 tours"}}

{{$d:=""}}
{{$v:="722755391498485800"}}
{{$m:=""}}
{{$k:=""}}
{{$id:=""}}
{{$f:=print "Recherche faite par " (getMember .User.ID).Nick }}
{{$b:=true}}
{{$a:="https://cdn.discordapp.com/attachments/726496591489400872/727978845185245283/download20200605012708.png"}}
{{$l:="https://i.imgur.com/o557fMx.png"}}

{{if .CmdArgs}}
	{{$i:=lower (index .CmdArgs 0)}}
	{{$flag:=reFind `(?i)(analgésiques?|(Armes? biologiques?)|(armes? bio)|Armes?)` (index .CmdArgs 0)}}
	{{$flag =lower $flag}}
	{{if eq $i "rixolam"}}
		{{$d =$ri}}
	{{else if eq $i "bandage" "bandages"}}
		{{$d =$ban}}
	{{else if eq $i "eufyrusant" "eufy"}}
		{{ $d =$eu}}
	{{else if eq $i "soma"}}
		{{$d =$so}}
	{{else if eq $i "sirop" "betapropyl" "sirop de betapropyl"}}
		{{$d =$ro}}
	{{else if eq $i "grenade" "grenade nécrotique" "grenade necrotique" "nécrotique" "necrotique" "necro" "nécro"}}
		{{$d =$gr}}
	{{else if eq $i "liquide antifongique" "liquide" "antifongique"}}
		{{$d =$fn}}
	{{else if eq $i "anesthésiant" "gaz anesthésiant" "gaz"}}
		{{$d =$th}}
	{{else if eq $i "sang etherique" "sang ethérique" "sang" "éthérique" "etherique" "étherique" "ethérique" "sang étherique"}}
		{{$d =$sng}}
	{{else if eq $i "carotoxinique" "caro" "huile carotoxinique"}}
		{{$d =$cr}}
	{{else if eq $i "huile digestive" "digestive"}}
		{{$d =$di}}
	{{else if eq $i "huile"}}
		{{$d =joinStr "" $di "\n" $cr}}
	{{else if eq $flag "analgésique" "analgésiques"}}
		{{$id ="736005431582916638"}}
	{{else if eq $flag "arme biologique" "armes biologiques" "arme biologiques" "armes biologique" "arme bio" "armes bio"}}
		{{$id ="736005737452404778"}}
	{{else if eq $flag "arme" "armes"}}
		{{$id ="736006328668913684"}}
	{{end}}

	{{if $flag}}
		{{$msg:=getMessage $v $id}}
		{{$m =(index $msg.Embeds 0).Description}}
		{{$k:=(print "(https://discordapp.com/channels/" .Guild.ID "/" $v "/" $id ")")}}
		{{$d =(joinStr "" $m )}}
	{{end}}

{{else}}
	{{$d ="Vous cherchez un objet à bas coût ? Mais lequel ?\n\n Pour rappel il existe :\n :white_small_square: Les armes (`?search arme(s)`)\n :white_small_square: Les armes biologiques (`?search \"arme bio\"`) \n :white_small_square: Les analgésiques (`?search analgésique`)"}}
{{end}}

{{$embed:=cembed
	"author" (sdict "name" "[Sola-UI] BDD : Objet bas de gamme" "icon_url" $a)
	"title" "Encyclopédie - Objet"
	"thumbnail" (sdict "url" $l)
	"description" $d
	"footer" (sdict "text" $f )
	"color" 0x94CAF0}}
{{sendMessage nil $embed}}
