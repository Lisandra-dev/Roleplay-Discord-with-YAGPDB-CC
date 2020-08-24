{{$yes:=715708401816043520}}
{{$no:=715708440466554921}}
{{$neutral:=742470860031393924}}

{{$react := sdict}}
{{with (dbGet 0 "react")}}
	{{$react = sdict .Value}}
{{end}}

{{$user := (str .User.ID)}}
{{$mid := str .Message.ID}}

{{$userReact := sdict}}
{{with ($react.Get $user)}}
	{{$userReact = sdict .}}
{{end}}

{{if .ReactionAdded}}
	{{if eq .Reaction.Emoji.ID $yes}}
		{{if ($userReact.Get $mid)}}
			{{if eq (str ($userReact.Get $mid)) "n"}}
				{{deleteMessageReaction nil .Message.ID .User.ID ":moins:715708440466554921"}}
			{{else if eq ($userReact.Get $mid) "neutral"}}
				{{deleteMessageReaction nil .Message.ID .User.ID ":blobshrug:742470860031393924"}}
			{{end}}
			{{$userReact.Del $mid}}
		{{end}}
		{{$userReact.Set $mid "y" }}
	{{else if eq .Reaction.Emoji.ID $no}}
	{{if ($userReact.Get $mid)}}
			{{if eq (str ($userReact.Get $mid)) "y"}}
				{{deleteMessageReaction nil .Message.ID .User.ID "this:715708401816043520"}}
			{{else if eq ($userReact.Get $mid) "neutral"}}
				{{deleteMessageReaction nil .Message.ID .User.ID ":blobshrug:742470860031393924"}}
			{{end}}
			{{$userReact.Del $mid}}
		{{end}}
		{{$userReact.Set $mid "n" }}
	{{else if eq .Reaction.Emoji.ID $neutral}}
		{{if ($userReact.Get $mid)}}
			{{if eq (str ($userReact.Get $mid)) "y"}}
				{{deleteMessageReaction nil .Message.ID .User.ID "this:715708401816043520"}}
			{{else if eq (str ($userReact.Get $mid)) "n"}}
				{{deleteMessageReaction nil .Message.ID .User.ID ":moins:715708440466554921"}}
			{{end}}
			{{$userReact.Del $mid}}
		{{end}}
			{{$userReact.Set $mid "neutral"}}
	{{end}}
{{end}}
{{$react.Set $user $userReact}}
{{dbSet 0 "react" $react}}
