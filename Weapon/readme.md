**WARNING** : No command here is independant !! You must take ALL of command listed here.
*The command are mainly in french*
Also, the count is only for "one person" : it recharges the weapon of the guy who start the command.

Note : All these command can be use with an reroll, with adding `#name` at the end.


# PA :

The center of the weapon / cooldown system.

**Trigger :** Coutains`:pa:`

> Counts the PA the user use. At 4, the user use all this PA. Each time he uses a PA, the function looks in the DB if the user recharge something (weapon, or competence). If yes, count it. If the count is good, trigger an embed and reset the recharge value.

# Recharge :
*Fun : it's one of the only command who doesn't use dictionnary because it's more pratical without*

**Trigger :** Command

**Usage :** `$recharge -(fusil|fusil2|pistolet|pistolet2|canon)`

For competence : `$recharge -(attaque|support) "name"`


Recharge work like that :
* Pistol : 2 PA
* Fusil : 4 PA
* Canon : 8 PA
* Competence attaque : 8 PA
* Competence support : 12 PA (Alteration delay : 8 PA, cooldown to 4 PA)

# Fusil count :

**Trigger** : countains `:fusil:`

**Usage :** The user send a message with the trigger.

> The bot count the number of bullets used by the user.

# Pistol count :

**Trigger** : countains `:pistol:`

**Usage :** The user send a message with the trigger.

> The bot count the number of bullets used by the user.

# Canon count :

**Trigger** : countains `:canon:`

**Usage :** The user send a message with the trigger.

> The bot count the number of bullets used by the user.

# Bullet stock

**Trigger** : Command

**Usage :** `-stock <pistolet|pistolet2|fusil|fusil2|canon>`

> The bot send the number of bullet for the person who send it.

# New turn :

**Trigger : ** Command

**Usage : ** `$turn (optional trigger)`

> Reset the PA for all people in the group dictionnary and add +1 to the DBkey for the turn. If argument, reset the turn number to 1.
