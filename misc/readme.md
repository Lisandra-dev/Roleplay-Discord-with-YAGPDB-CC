All command here can work as stand alone. That's why there are ungrouped.

# Pins

**Trigger** : Command

**Usage** : `$pins`

> Le bot envoie juste un gif.

# Instruction

**Trigger** : Command

**Usage** : `$instruction`

> Commande réservée aux administrateur, elle envoie un message à une personne mentionnée qui lui indique quoi faire une fois que son perso a été validé.

# Bloc

**Trigger** : Command

**Usage** : `-bloc`
> La partie code de la commande précédente. Ce code contient tout ce qui doit être mis pour résumé son personnage.

# Delete HRP message

**Trigger** : Regex `^\((.*)`

**Usage** : /
> Supprime tout le HRP dans les channels RP, c'est à dire toutes les commandes commençant par une parenthèse.

# Quoting By link

*Source : The YAGPDB server*
**Trigger** : Regex `https://discordapp.com/channels\/(\d+)\/(\d+)\/(\d+)`

**Usage** : Send a message with a discord link message in the server.
> Le bot envoie un message citant le message dont le lien a été posté. Il peut même citer les embed de divers bots.

# End

**Trigger** : Command

**Usage** : `$end`

> Le bot envoie une ligne dans le channel où il a été trigger.

# Edit

**Trigger** : Command

**Usage** : `-edit <id> <things>`
> Permet d'éditer un message de YAG, dont des embed.
>
> Source : https://github.com/Jo3-L/yagpdb-cc
> Voir au dessus pour plus d'information.

# Google:

*Source : The Yagpdb github CC*

**Trigger** : Command

**Usage** : `$google <things to search>`

> Permet d'envoyer une recherche google à quelqu'un.

# AFK

**Trigger** : command

**Usage** : `-afk -d duration (en heure) <raison>`

> Permet de donner le rôle AFK à quelqu'un durant x temps, avec une raison optionnel. Un message indiquant son absence est envoyé dans un channel, et le message est automatiquement supprimé si la personne supprime son AFK (avec la commande) ou sort naturellement de l'AFK.

# Viewafk

**Trigger** : Command

**Usage** : Regex : `^\$(viewafk <@!?\d+>)`

> Permet de voir le temps restant d'une AFK, mais aussi la raison.

# Snippet

**Trigger** : Starts with ?

**Usage** : `?(snippet|all|help|bouclier|horloge|notes|ticket|charge|position|dégâts|dégât|dégat|dégat|dés|résumé|arme|armes)`

> Donne certaines informations en rapport avec le nom. Affiche parfois des messages du serveur.

# Support help

**Trigger :** Regex `^\?(malus|soin|poison|support)`

**Usage : ** `?(malus|soin|poison|support)`

> Analogue à la précédente commande, mais était trop longue pour y être rajouté. Donne les infos sur les résultats des supports/altérations, pour tous les résultats possibles.

# Delete emoji

**Trigger** : Regex `((<a?:[\w~]{2,32}:\d{17,19}>)|[\x{1f1e6}-\x{1f1ff}]{2}|\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?(\x{200D}\p{So}\x{fe0f}?[\x{1f3fb}-\x{1f3ff}]?)*|[#\d*]\x{FE0F}?\x{20E3})`

**Usage** : /

> Supprime les émoji "solitaire".

# play or skip :

**Trigger :** Reaction added/removed

> Edit un message de YAG, et rajoute l'utilisateur dans l'embed. Permet d'indiquer rapidement les roles des membres sur le message.

# rm player guide :

**Trigger :** Reaction added / Removed

> Envoie un message aux personnes choisissant d'être joueur, avec un guide pour ne pas être perdu dans les channels.

# Content :

**Trigger** : Command

**Usage** : `$content <id>`
Pour récupérer un message d'un autre channel : `$content -chan <id> <message id>`

> Permet de récupérer, en format "code" le contenu d'un message, qui peut être un embed ou un message "normal".
> Attention, la commande bug sur les messages contenant à la fois un embed et un message texte.
> Le message va se supprimer après 3 minutes. Pratique sur les gros sujets très suivi.
