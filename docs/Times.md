# Count

**Trigger** : Regex `.*`

**Usage** : /
> Le bot compte le nombre de message envoyé par les joueurs dans les channels. Il incrémente alors diverses variables, et quand l'une d'elle atteint une somme fixée, il envoie un message, reset le cycle, et si le cycle est le 4e, il augmente aussi le jour.
>
> Il est possible de modifier le nombre de cycle dans la commande. 

# Time

**Trigger** : Command

**Usage** : `$time`

> The bot send an embed associated at the time. It also indicates the number of messages sent during the cycle, in pourcent.
>
> Le bot envoie un message indiquant le nombre de message dans le cycle (en footer mais aussi en description avec un pourcentage). Contrairement au "count", le temps est envoyé dans le channel où la commande est envoyé. 

# Reset

**Trigger** : Command

**Usage** : `$reset`

> Reset le cycle actuel, les message, et le jour à respectivement 1, 0 et 1. 
>
> Afin d'éviter des miss click, cette commande a une sécurité et force l'utilisateur à taper une seconde commande aléatoire pour reset. 
>
> 

# Changer le nom d'un channel en fonction du temps

**Trigger** : Minute interval : 5 minutes.

**Usage** : /

> Attention, à cause des limitations discord, il n'est pas possible de rendre cette commande plus rapide. Elle permet de changer le nom d'un channel pour indiquer rapidement le cycle et le jour actuel. 

# Changing number for cycle

**Trigger** : command

**Usage** : `$settm <cycle|msg|speed|day>`

> Cette commande permet de changer les paramètres lié au temps. Chaque modification envoie un embed dans un channel précis (ici celui indiquant le temps). Il permet de fixer le nombre de message dans un cycle, le jour, le n° du cycle, et le nombre de message pour le passage d'un cycle. 

# Settings of timer :

**Trigger** : Command

**Usage** : `-settings`

> Permet de voir tous les paramètres du temps. Ainsi : 
- Le nombre de message actuellement dans le cycle
- Le cycle actuel
- Le jour actuel
- La vitesse du cycle (nombre de message avant le changement)

Même si ses paramètres n'ont pas été modifié par la commande précédente. 

