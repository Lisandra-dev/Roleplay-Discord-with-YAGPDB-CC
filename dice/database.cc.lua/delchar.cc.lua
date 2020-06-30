{{$user := (index .CmdArgs 0)}}
{{dbDel $user "force"}}
{{dbDel $user "endurance"}}
{{dbDel $user "agi"}}
{{dbDel $user "preci"}}
{{dbDel $user "intelligence"}}
{{dbDel $user "karma"}}

{{dbDel  $user "i_force"}}
{{dbDel $user "i_endu"}}
{{dbDel $user "i_agi"}}
{{dbDel $user "i_preci"}}
{{dbDel $user "i_intel"}}
{{dbDel $user "i_karma"}}

Done !
