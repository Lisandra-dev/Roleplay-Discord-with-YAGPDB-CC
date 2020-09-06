Merci à Shadownight#9025 de me permettre d'utiliser ses commandes.

Toutes les commandes pour acheter, vendre, ouvrir l'inventaire et le porte-monnaie peuvent être utiliser avec `#name` à **la fin de la commande**
Pour les admins, l'user peut être un reroll, en ajoutant `#`à la place de `user`.

Toutes ses commandes utilisent des dictionnaires ou des listes, sauf pour l'ouverture / fermeture du shop, qui est une simple entrée de BD.

# Administration : 🧰

Ces commandes sont uniquement contrôlés et contrôlables par les administrateurs.

## Management :

Ces commandes permettent de gérer l'économie générale : argent des joueurs, leur inventaire, symbole de la monnaie...

**Trigger :** Command

**Usage :**

* Money : `$adminmoney -(add|remove|reset|set) user amount`
 > L'utilisateur sera averti si un admin change son porte-monnaie.

* Symbol : `$symbol <symbol>`

* Manage inventory : `$manageinv -(add|remove|set) user (@rerollaccount) item amount`

* Reset/clean inventory : `$manageinv -(reset|clean) user (@rerollaccount)`

  > Reset détruit complètement l'inventaire du joueur, quand clean le vide seulement.

* Manage équipement : `$managestuff -(add|rm|reset|init) user (@user_reroll) item amount`

> Permet de rajouter (sans passer par les casiers) ou retirer des objets de l'inventaire d'un joueur, tout en vérifiant les emplacements disponibles. 
> L'initialisation doit se faire pour TOUS les nouveaux joueurs. Elle est primordiale.

## Item

Ces commandes permettent de rajouter, modifier, ou supprimer un objet dans la boutique.

**Trigger :** Command

**Usage :**

* Add : `$item -add "name" pricemin pricemax (sellmin|none) (sellmax|0) (inf|stockage)(true|false) rareté "description"`
* Delete : `$item -delete "name"`
* Information : `$item -info "name"`

> - Attention à la casse.
>
> - SII veut dire si l'objet peut apparaître ou non dans l'inventaire du joueur.

* `$item -edit "name" -(price|sell|stock|sii|desc|rare|rename) "value"`

Type que doit être `value` :
* <u>Prix :</u> Prix minimum et prix maximum.
* <u>Vendre</u> : `0 0` ou un prix minimum et maximum.
* <u>Stock</u> :  `inf` ou un nombre
* <u>SII</u> : true ou false
* <u>desc</u> : Chaîne de caractère.
* <u>Rare</u> : Nombre
* <u>Rename</u> : Chaine de caractère

## Open & close

**Trigger :** Regex `\$(close|open)`

**Usage :** `$close` ou `$open`

> - Ouvre ou ferme la boutique.
> - Si la boutique est fermée, les utilisateurs ne peuvent ni vendre ou acheter.
> - La boutique aura un message spécial si la boutique est fermée.
> - Un message est envoyé dans un channel pour indiqué l'ouverture/fermeture de la boutique.


# Store

**Trigger :** Command

**Usage :** `$store (optional page)`

Pour la pagination : Trigger : Réaction → Added Reaction

> Montre la boutique.
> Que la boutique soit fermée ou ouverte, un message sera affiché.
> S'il y a beaucoup d'objet, l'utilisateur peut indiquer une page. De plus, des émoji seront ajouté à la boutique et permettent de changer de page, ou supprimer le message.


# Echange :

**Trigger :** Commande

**Usage :** `$give (@mention_cible|&reroll_cible) (@joueur du reroll) <argument> (#reroll_donneur)`

> Les arguments entre parenthèses sont optionnels.
> En fonction de ce que vous voulez faire, l'argument change :
> - Donner de l'argent : nombre
> - Donner un objet : Nom de l'objet + Quantité éventuelle

# Équipement

## Affichage

### Commande

**Trigger : ** Command

**Usage** :** `$inv (optional page)`

### Pagination

**Trigger : ** Reaction : Added

**Usage** : Les émoji permettent de tourner les pages.

# Casier

  ## Equiper : 

**Trigger** : Commande

**Usage** : `$equip "objet" (quantité)`

> Retire du casier l'objet ciblé et l'équipe. 
> Compte les places restantes. 
> Les composants ne prennent pas de place et sont reconnus par leurs initiales.
> Les armes sont reconnues par leur types.
> Les sacs sont reconnus par leurs noms "simples" : nul besoin de rajouter le sigle "[E]".

  ## Ranger : 

**Trigger** : Commande

**Usage** : `$ranger "objet" (quantité)`

> Comme précédemment, mais en inversé : Retire un objet de l'inventaire et le remet dans le casier, tout en comptant les places restantes.

# Achat et vente :

## Sell

**Trigger :** Command

**Usage :** `$sell "name" amount`

> Vend l'objet.
> Pour les composants, name peut être bc/lc/cu/cf/cb
> La fonction reconnait automatiquement pour les chargeurs, il suffit d'indiquer chargeur + type d'arme.


## Buy

**Trigger :** Command

**Usage :** `$buy "name" amount (+VN)` 

> Achete l'objet et le met dans votre casier.
> Pour les composants, name peut être bc/lc/cu/cf/cb
> La fonction reconnait automatiquement pour les chargeurs, il suffit d'indiquer chargeur + type d'arme.
> `+VN` indique que l'achat doit être remis dans l'inventaire commun. 


# Money

**Trigger :** Command

**Usage :** `$money`

> Affiche le porte monnaie de l'utilisateur.

# VN

## Add :

**Trigger :** Regex ^\$VNadd

## Use :

**Trigger** : Regex ^\$VNuse

<br>

**Usage :** `$(vnadd|vnuse) (-armes?|modules?|soin(s)|stuff|bc|lc|cb|sf|cu|chargeur) "value" #nom`

**Indication sur la valeur :**
* <u>Armes, module, soin, stuff </u>: nom
* <u> Composants :</u> Quantité
* <u> Chargeur : </u> Quantité suivi du type d'arme tel que : `valeur (pistolet|fusil|canon)`

> La fonction va mettre à jour le message de l'inventaire à chaque fois que quelqu'un utilise la commande pour éditer cet inventaire.
> Le message de l'inventaire doit **obligatoirement** être fait par YAGPD.
> Les objets seront retiré/rajouté dans le casier de la personne qui l'utilise.
