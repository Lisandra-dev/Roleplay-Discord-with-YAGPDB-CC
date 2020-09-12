Toutes les commandes listées peuvent être utiliser avec un reroll, en indiquant `#nom` à la fin de la commande. En outre, les emoji peuvent être indiqué à la fin de la commande.

# PA

## PA et PA2

Le centre du système d'arme.

**Trigger :** Coutains `:pa:`

**Trigger :** Countains `:2pa:`

> Compte les PA d'un utilisateur à chaque trigger. A quatre, l'utilisateur a utilisé tous ses PA. A chaque fois que la commande est lancé, elle recherche dans la base de données si l'utilisateur recharge quelque chose (arme ou compétence). Si oui, elle va faire `+1`. Si cette valeur atteint le seuil, alors elle va indiquer le rechargement complet puis supprimer la clé de la base de données. En outre, pour les armes, elle va reset le compte et indiquer donc le maximum de charge.

## Recharges

### Compétences

**Trigger :** Commande

**Usage :** `$cd (long|court) "nom"`

- Capacité courte : 4 PA
- Capacité longue : 8 PA

En général, les actions à CD courts sont les compétences offensives (burst, perforant, grosse attaque liés à un pouvoir). Celles ayant un CD long sont les compétences de supports (altérations, malus, boost...).

### Recharge manuelle des armes :

* **Il faut absolument que l'utilisateur ait des chargeurs dans son inventaire et que les noms soient indiquent à la fois dans l'inventaire, la boutique, et la fonction.**

 > Pour des raisons de praticités, les chargeurs portent le nom de l'arme. Si dans la boutique, des armes (différente d'un chargeur) doivent y être mis, je conseille de rajouter un indicateur pour les ARMES (**et non les chargeurs**), tel que [A] ou [E] par exemple.
 > Le nom peut être cependant changé, dans ce cas, il faut changer la ligne indiqué en indiquant le nom exacte des chargeurs se trouvant dans la boutique.

* Les fichiers `pa` et `2pa` ne contiennent la recharge que pour les compétences.
* Le fichier `competence` ne change pas entre les deux versions.
* Note : Recharger coûte 2PA, la fonction le vérifie donc.
* Le compte des `pa` et `2pa` peut être indiqué directement dans la commande.

**Trigger :** Commande

**Usage :** `$recharge (fusil|fusil2|pistolet|pistolet2|canon)`

## PA restant :

**Trigger** : Command

**Usage** : `$pa`

→ Permet au joueur de voir ses PA restants, tout simplement.

## Turn

**Trigger :**Command

**Usage :** `$turn (random argument optional)`

> Reset pour tout les membres du groupe les PA, et rajoute +1 au nombre de turn, sauf si un argument est passé en paramètre. Dans ce cas, le tour retourne à 1, et le groupe est reset. 

# Armes

## Fusil, fusil2, pistolet, pistolet2, canon

**Trigger :** Contains :

* Pistolet : `:pistol:` & `:pistol2:`

* Fusil : `:fusil` & `:fusil2`

* Canon : `:canon:`

  > Le bot compte les balles utilisés par l'utilisateur quand il l'indique avec l'émoji associé.
  > Les emoji peuvent être indiqué directement dans la commande de dés.

## Bullet stock

**Trigger :** Command

**Usage :** `$stock -(pistolet|pistol2|canon|fusil|fusil2)`

> Indique le nombre de charge restante dans l'arme. Si l'utilisateur n'a pas utilisé son arme, indique le nombre de charge maximal.

## Delweapon

**Trigger :** Command

**Usage :** `$delweapon <mention|#name>`


> Permet de supprimer dans la base de données toutes les données des armes relative à la personne passée en paramètre.


# Bouclier

Le bouclier fonctionne un peu comme le temps : il va compter les messages envoyés dans un channel jusqu'à atteindre un seuil. 
## Activation

**Trigger** : Command

**Usage :** `$bouclier`

> Le bot commence un compte de message pour recharger le bouclier.  

## Recharge

**Trigger :** Regex `.*`

**Usage : /**

> Compte tous les messages après le début de la recharge. Il faut 11 messages pour la terminer. 
