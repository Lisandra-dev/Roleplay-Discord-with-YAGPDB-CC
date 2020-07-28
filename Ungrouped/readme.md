All command here can work as stand alone. That's why there are ungrouped.

# Pins :
**Trigger** : Command

**Usage** : `-pins`
> The bot send a message with a gif indicate to click on the pins if the response is pinned.


# Instruction :
**Trigger** : Command

**Usage** : `-instruction`
> Send a big message, in french, by an admin, to an user, indicate what he must do after the validation of his character.

# Bloc :
**Trigger** : Command

**Usage** : `-bloc`
> Send a message for resume a characters. It's just the code party from the instruction.

# Delete HRP message :
**Trigger** : Regex `\((.*?)\)`

**Usage** : /
> The bot delete all message between bracket.

# Quoting By link :
*Source : The YAGPDB server*
**Trigger** : Regex `https://discordapp.com/channels\/(\d+)\/(\d+)\/(\d+)`

**Usage** : Send a message with a discord link message in the server.
> The bot send a embed with the message quoted. It deletes the message if the message have only the link. Moreover, the embed have the link to the original, mainly for embed who cannot be quoted.

# End :
**Trigger** : Command

**Usage** : `-end`
> The bot send a message with only a line, and delete the trigger.

# Edit :
**Trigger** : Command

**Usage** : `-edit <id> <things>`
> Edit a message send by the bot.
> From https://github.com/Jo3-L/yagpdb-cc
> See CC for more information

# Google :
*Source : The Yagpdb github CC*

**Trigger** : Command

**Usage** : `-google <things to search>`
> Yag will send a google link to search the things in parameters.

# AFK :

**Trigger** : command

**Usage** : `-afk -d duration (in second) <reason>`

> Set an role AFK in the user who trigger it, for a certain time. Send an embed in the channel for afk with the reason and duration.

# Viewafk

**Trigger** : Command

**Usage** : Regex : `^\$(viewafk <@!?\d+>)`

> View the afk duration for an user and the reason behind it.

# dice info (dinfo)

**Trigger** : command

**Usage** : `-dinfo (soin|poison|malus)`

> View information about the dice

# Snippet

**Trigger** : Starts with ?

**Usage** : `?(snippet|all|help|bouclier|horloge|notes|ticket|charge|position|dégâts|dégât|dégat|dégat|dés|résumé|arme|armes)`

> Quote some message from the discord

# Delete emoji :

**Trigger** : Regex `((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})`

> Delete solo emoji.
