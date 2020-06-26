{{$malus:=reFindAllSubmatches `^\$malus` .Message.Content}}
{{$soin:= reFindAllSubmatches `^\$soin` .Message.Content}}
{{$poison:= reFindAllSubmatches `^\$poison` .Message.Content}}

{{$d:= (randInt 1 10)}}
{{$v := ""}}
{{$implant := "implant"}}

{{$user:= .Member.Nick}}
  {{if eq (len $user) 0}}
    {{$user = .User.Username}}
  {{end}}

{{if .CmdArgs}}
  {{if ne (toInt (index .CmdArgs 0)) (toInt 0)}}
    {{$i:=(toInt  (index .CmdArgs 0))}}
    {{if ge (toInt $i) (toInt 2)}}
      {{$implant = "implants"}}
    {{end}}
    {{$v = (joinStr "" "(" $d ")")}}
    {{$d = sub $d $i}}
  {{end}}
{{end}}

{{$res:= (joinStr " " "Dé :" (toString $d) (toString $v) )}}

{{if $malus}}
	{{if le $d (toInt 0)}}
**{{$user}}** ▬ **Ultra critique !**
	Votre cible a un malus de +4 à son dé.
		*[{{$res}}]*
	{{else if eq $d (toInt 1) }}
**{{$user}}** ▬ **Réussite critique !**
	Votre cible a un malus de +3 à son dé.
		*[{{$res}}]*
    {{else if eq $d (toInt 2)}}
**{{$user}}** ▬ Votre cible a un malus de +3 à son dé.
	*[{{$res}}]*
	{{else if and (le $d (toInt 5)) (ge $d (toInt 3))}}
**{{$user}}** ▬ Votre cible a un malus de +2 à son dé.
	*[{{$res}}]*
	{{else if  and (le $d (toInt 8)) (ge $d (toInt 6))}}
**{{$user}}** ▬ Votre cible a un malus de +1 à son dé.
	*[{{$res}}]*
	{{else if eq $d (toInt 9)}}
**{{$user}}** ▬ **Echec du malus**...
	*[{{$res}}]*
	{{else if eq $d (toInt 10)}}
**{{$user}}** ▬ **Echec critique du malus ! **
	Votre cible a un bonus de -2 à son dé.
		*[{{$res}}]*
	{{end}}

{{else if $soin}}
	{{if le $d (toInt 0)}}
**{{$user}}** ▬ **Ultra critique**
	Votre cible regagne 8 PV *(+1 si module/PSI)* et obtient un bonus de votre choix."
		*[{{$res}}]*
	{{else if eq $d (toInt 1)}}
**{{$user}}** ▬ **Réussite critique**
	Votre cible regagne 8 PV *(+1 si module/PSI).*
		*[{{$res}}]*
	{{else if eq $d (toInt 2)}}
**{{$user}}** ▬ Votre cible regagne 7 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 3)}}
**{{$user}}** ▬ Votre cible regagne 6 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 4)}}
**{{$user}}** ▬ Votre cible regagne 5 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 5)}}
**{{$user}}** ▬ Votre cible regagne 4 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 6)}}
**{{$user}}** ▬ Votre cible regagne 3 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 7)}}
**{{$user}}** ▬ Votre cible regagne 2 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 8)}}
**{{$user}}** ▬ Votre cible regagne 1 PV *(+1 si module/PSI)*
	*[{{$res}}]*
	{{else if eq $d (toInt 9)}}
**{{$user}}** ▬ **Echec du soin ...**
	*[{{$res}}]*
	{{else if eq $d (toInt 10)}}
**{{$user}}** ▬ **Echec critique du soin !**
	*Votre cible gagne un malus*
		*[{{$res}}]*
	{{end}}

{{else}}
	{{if le $d (toInt 0)}}
**{{$user}}** ▬ **Ultra critique**
	Votre cible a un empoisonnement de 10 PV sur 3 tours (-30 PV en tout).
		*[{{$res}}]*
	{{else if eq $d (toInt 1)}}
**{{$user}}** ▬ **Réussite critique**
	Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).
		*[{{$res}}]*
	{{else if eq $d (toInt 2)}}
**{{$user}}** ▬ **Réussite critique**
	Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).
		*[{{$res}}]*
	{{else if or (eq $d (toInt 3)) (eq $d (toInt 4))}}
	**{{$user}}** ▬ Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).
		*[{{$res}}]*
	{{else if or (eq $d (toInt 5)) (eq $d (toInt 6))}}
	**{{$user}}** ▬ Votre cible a un empoisonnement de 4 PV sur 3 tours (-12 PV en tout).
		*[{{$res}}]*
	{{else if or (eq $d (toInt 7)) (eq $d (toInt 8))}}
**{{$user}}** ▬  Votre cible a un empoisonnement de 2 PV sur 3 tours (-6 PV en tout).
	*[{{$res}}]*
	{{else if eq $d (toInt 9)}}
**{{$user}}** ▬ **Echec du poison**...
	*[{{$res}}]*
	{{else if eq $d (toInt 10)}}
**{{$user}}** ▬ **Echec critique de l'empoisonnement !**
	Votre cible gagne une régénération de 3 PV sur 3 tours (+9 PV en tout).
		*[{{$res}}]*
	{{end}}
{{end}}
{{deleteTrigger 1}}
