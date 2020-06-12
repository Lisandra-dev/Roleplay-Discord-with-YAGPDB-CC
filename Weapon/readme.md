**WARNING** : No command here is independant !! You must take ALL of command listed here.
*The command are mainly in french*



Also, the count is only for "one person" : it recharges the weapon of the guy who start the command. 



# Recharge : 

**Trigger :** Command

**Usage :** `-recharge <fusil|Pistolet|Canon>`

> The bot start a "counter" for the recharge of a db key link to a word.

Here, there are three weapon : Fusil, pistolet and canon. The DBkey are : 

- Fusil : `cf`and `fusil`
- Canon : `ca`and `canon`
- Pistolet : `cp`and `pistol`

If you change for one command, you must change all. You can add more weapon, just C/C one of things in recharge and create a count with a C/C. Don't forget to create your key. 

# Increment message :

**Trigger :** Regex `.*` 

**Usage :** /

> The bot increment the key for each message after the start of the command "recharge"
>
> - For fusil : It needs 6 messages.
> - Canon : 11 messages
> - Pistolet : 4 messages
>
> The increment work only for message with 15 characters length. 

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

