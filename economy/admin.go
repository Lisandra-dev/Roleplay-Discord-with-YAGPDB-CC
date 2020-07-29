{{/* Databases */}}
{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
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

{{/* Command Body */}}
{{$flag := ""}}
{{if .CmdArgs}}
	{{$flag = (reFind `\$(adminmoney|symbol|manageinv)` (index .CmdArgs 0))}}
{{end}}
{{$mention := ""}}


{{ if eq $flag "$money"}}
	{{$secondflag := ""}}
	{{$target := ""}}
	{{$amount := 0}}
	{{with .CmdArgs}}
		{{$secondflag = (reFind `\-(add|remove|reset|set)` (index . 1))}}
		{{if ge (len .) 3}}
			{{if $name}}
				{{$target = $id}}
				{{$user = $name}}
				{{$mention = $name}}
			{{else}}
				{{with (userArg (index . 2))}}
					{{$target = .}}
					{{$target = $target.ID}}
					{{$mention = joinStr "" "<@" $target ">"}}
					{{$user = (getMember $target).Nick}}
				{{end}}
			{{end}}
			{{if ge (len .) 4}}
				{{$amount = toInt (index . 3)}}
			{{end}}
		{{end}}
	{{end}}
	{{if and $secondflag $target}}
		{{$userEco := sdict}}
		{{with (dbGet $target "economy")}}
			{{$userEco = sdict .Value}}
		{{end}}
		{{$bal := (toInt ($userEco.Get "balance"))}}
		{{if eq $secondflag "-add"}}
			{{if $amount}}
				{{$bal = add $bal $amount}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/ATSj8fe.png")
				"description" (joinStr " " $mention " a gagné" $amount "<:klix:724739705954107405>\n Vous avez donc actuellement" $bal "<:klix:724739705954107405>")
				"color" 0x8CBAEF}}
				{{sendMessage nil $embed}}
			{{else}}
				Vous avez besoin de spécifier une valeur !
				**Usage : ** `$adminmoney -(add|remove|reset|set) user amount`
			{{end}}
		{{else if eq $secondflag "-remove"}}
			{{if $amount}}
				{{$bal = sub $bal $amount}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/ATSj8fe.png")
				"description" (joinStr " " $mention " a perdu" $amount "<:klix:724739705954107405>\n Vous avez donc actuellement" $bal "<:klix:724739705954107405>")
				"color" 0x8CBAEF}}
				{{sendMessage nil $embed}}
			{{else}}
				Vous avez besoin de spécifier une valeur !
				**Usage : ** `$adminmoney -(add|remove|reset|set) user amount`
			{{end}}
		{{else if eq $secondflag "-reset"}}
			{{$bal = 0}}
			Reset de {{$mention}}.
		{{else if eq $secondflag "-set"}}
			{{if $amount}}
				{{$bal = $amount}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" "https://i.imgur.com/ATSj8fe.png")
				"description" (joinStr " "  "La balance de" $mention "a été fixé à :" $amount "<:klix:724739705954107405>\n Vous avez donc actuellement" $bal "<:klix:724739705954107405>")
				"color" 0x8CBAEF}}
				{{sendMessage nil $embed}}
			{{else}}
				Vous avez besoin de spécifier une valeur !
				**Usage : ** `$adminmoney -(add|remove|reset|set) user amount`
			{{end}}
		{{end}}
		{{$userEco.Set "balance" $bal}}
		{{dbSet $target "economy" $userEco}}
	{{end}}

{{else if eq $flag "$symbol"}}
	{{$symbol := ""}}
	{{if $serverEco.Get "symbol"}}
		{{$symbol = $serverEco.Get "symbol"}}
	{{end}}
	{{with .CmdArgs}}
		{{if ge (len .) 2}}
			{{if ne (index . 1) "none"}}
				{{$symbol = joinStr " " (slice . 1)}}
			{{else}}
				{{$symbol = ""}}
			{{end}}
		{{end}}
	{{end}}
	{{$serverEco.Set "symbol" $symbol}}
	Le symbole de la monnaie est maintenant : `{{or $symbol "none"}}`

{{else if eq $flag "$manageinv"}}
	{{$secondflag := ""}}
	{{$target := ""}}
	{{$item := ""}}
	{{$amount := 1}}
	{{with .CmdArgs}}
		{{if ge (len .) 2}}
			{{$secondflag = (reFind `\-(add|remove|reset|clean|set)` (index . 1))}}
			{{if ge (len .) 3}}
			{{if $name}}
				{{$target = $id}}
				{{$user = $name}}
				{{$mention = $name}}
			{{else}}
				{{with (userArg (index . 2))}}
					{{$target = .}}
					{{$target = $target.ID}}
					{{$mention = joinStr "" "<@" $target ">"}}
					{{$user = (getMember $target).Nick}}
				{{end}}
			{{end}}
				{{if ge (len .) 4}}
					{{$item = (index . 3)}}
					{{if ge (len .) 5}}
						{{$amount = or (toInt (index . 4)) 1}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
	{{if not $secondflag}}
		**Usage** : `$manageinv -(add|tremove|reset|clean|set) user item amount`
	{{else}}
		{{$userEco := sdict}}
		{{if $target}}
		{{with (dbGet $target "economy")}}
			{{$userEco = sdict .Value}}
		{{end}}
	{{else}}
		**Usage** : `$manageinv -(add|tremove|reset|clean|set) user item amount`
	{{end}}
	{{if eq $secondflag "-clean"}}
		{{if $target}}
			{{$items := sdict}}
			{{with ($serverEco.Get "Items")}}
				{{$items = sdict .}}
			{{end}}
			{{$itemscslice := cslice}}
			{{$inv := sdict}}
			{{with ($userEco.Get "Inventory")}}
				{{$inv = sdict .}}
			{{end}}
			{{range $k,$v := $items}}
				{{$itemscslice = $itemscslice.Append $k}}
			{{end}}
			{{range $k,$v := $inv}}
				{{if not (in $itemscslice $k)}}
					{{$inv.Del $k}}
				{{end}}
			{{end}}
			Inventaire de : {{$mention}} Nettoyé !
			{{$userEco.Set "Inventory" $inv}}
		{{end}}

	{{else if eq $secondflag "-reset"}}
		{{if $target}}
			{{$userEco.Del "Inventory"}}
			Reset de l'inventaire de : {{$mention}}.
	{{end}}

	{{else if eq $secondflag "-add"}}
		{{$items := sdict}}
		{{with ($serverEco.Get "Items")}}
			{{$items = sdict .}}
		{{end}}
		{{$inv := sdict}}
		{{with ($userEco.Get "Inventory")}}
			{{$inv = sdict .}}
		{{end}}
		{{if $items.Get $item}}
			{{if ($items.Get $item).sii}}
				{{$inv.Set $item (add (toInt ($inv.Get $item)) $amount)}}
				{{$userEco.Set "Inventory" $inv}}
				Ajout de : {{$amount}} {{$item}} à l'inventaire de {{$mention}}.
			{{else}}
			Cet item ne peut pas apparaître chez quelqu'un !
			{{end}}
		{{else}}
			Cet item n'existe pas.
		{{end}}

		{{else if eq $secondflag "-remove"}}
			{{if $target}}
				{{$items := sdict}}
				{{with ($serverEco.Get "Items")}}
					{{$items = sdict .}}
				{{end}}
				{{$inv := sdict}}
				{{with ($userEco.Get "Inventory")}}
					{{$inv = sdict .}}
				{{end}}
				{{if $inv.Get $item}}
					{{$value := (sub (toInt ($inv.Get $item)) $amount)}}
					{{if ne $value 0}}
						{{$inv.Set $item $value}}
					{{else}}
						{{$inv.Del $item}}
					{{end}}
					{{$userEco.Set "Inventory" $inv}}
					{{$amount}} : {{$item}} de l'inventaire de {{$mention}} a été retiré.
				{{else}}
					Cet objet n'existe pas.
				{{end}}
			{{end}}
		{{else if eq $secondflag "-set"}}
			{{if $target}}
				{{$items := sdict}}
				{{with ($serverEco.Get "Items")}}
					{{$items = sdict .}}
				{{end}}
				{{$inv := sdict}}
				{{with ($userEco.Get "Inventory")}}
					{{$inv = sdict .}}
				{{end}}
				{{if $items.Get $item}}
					{{if ($items.Get $item).sii}}
						{{$value := $amount}}
						{{if ne $value 0}}
							{{$inv.Set $item $value}}
						{{else}}
							{{$inv.Del $item}}
						{{end}}
						{{$userEco.Set "Inventory" $inv}}
						{{$amount}} : {{$item}} a été mis dans l'inventaire de {{$mention}}
					{{else}}
						Cet item ne peut pas être vu.
					{{end}}
					{else}}
						L'objet n'existe pas.
					{{end}}
				{{end}}
			{{end}}
			{{if $target}}
				{{dbSet $target "economy" $userEco}}
			{{end}}
		{{end}}
	{{else}}
	**Usage** :
:white_small_square: __Money__ : `$adminmoney -(add|remove|reset|set) user amount`
:white_small_square: __Symbol__ : `$symbol <symbol>`
:white_small_square: __Manage inventory__ : `$manageinv -(add|remove|set) user item amount`
:white_small_square: __Reset inventory__ : `$manageinv -reset user`
{{end}}

{{/* Database Updates */}}
{{dbSet .Guild.ID "economy" $serverEco}}
