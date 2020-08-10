{{$yes:=715708401816043520}}
{{$no:=715708440466554921}}
{{if .ReactionAdded}}
	{{if eq .Reaction.Emoji.ID $yes}}
		{{dbSet .User.ID "y" 1}}
		{{if dbGet .User.ID "n"}}
			{{dbDel .User.ID "n"}}{{deleteMessageReaction nil .Message.ID .User.ID ":moins:715708440466554921"}}
		{{end}}
	{{else if eq .Reaction.Emoji.ID $no}}
		{{dbSet .User.ID "n" 1}}
		{{if dbGet .User.ID "y"}}
			{{dbDel .User.ID "y"}}{{deleteMessageReaction nil .Message.ID .User.ID "this:715708401816043520"}}
		{{end}}
	{{end}}
{{end}}
