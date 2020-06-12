All command here can work as stand alone. That's why there are ungrouped.

# Pins :
**Trigger** : Command

**Usage** : `-pins`
> The bot send a message with a gif indicate to click on the pins if the response is pinned.

# Soin - Healing system :
**Trigger** : Command

**Usage** : `-soin`
> The bot send, in french, a message indicate the HP healed.

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

**Usage** : `-edit <channel ID> <Message ID> <new message>`
> Edit a message send by the bot.

# Google :
*Source : The Yagpdb github CC*

**Trigger** : Command

**Usage** : `-google <things to search>`
> Yag will send a google link to search the things in parameters.
