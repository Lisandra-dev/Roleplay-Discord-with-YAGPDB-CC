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

* Manage inventory : `$manageinv -(add|remove|set) user item amount`

* Reset/clean inventory : `$manageinv -(reset|clean) user`

  > Reset détruit complètement l'inventaire du joueur, quand clean le vide seulement.



## Item

Ces commandes permettent de rajouter, modifier, ou supprimer un objet dans la boutique.

**Trigger :** Command

**Usage :**

* Add : `$item -add "name" price (none|sellprice) (inf|stockage)(true|false) rareté "description"`
* Delete : `$item -delete "name"`
* Information : `$item -info "name"`

> - Attention à la casse.
>
> - SII veut dire si l'objet peut apparaître ou non dans l'inventaire du joueur.

* `$item -edit "name" -(price|sell|stock|sii|desc|rare|rename|rbuy|rsell) "value"`

Type que doit être `value` :
* <u>Prix :</u> nombre.
* <u>Vendre</u> : `0` ou un nombre.
* <u>Stock</u> :  `inf` ou un nombre
* <u>SII</u> : true ou false
* <u>desc</u> : Chaîne de caractère.
* <u>Rare</u> : Nombre
* <u>Rename</u> : Chaine de caractère
* <u> rbuy & rsell</u> : Ce sont les deux seules commandes qui ont besoin de deux valeurs : une borne "start" et "end" qui permet de choisir aléatoirement un nombre entre ses deux valeurs.

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

> Montre la boutique.
> Que la boutique soit fermée ou ouverte, un message sera affiché.
> S'il y a beaucoup d'objet, l'utilisateur peut indiquer une page. De plus, des émoji seront ajouté à la boutique et permettent de changer de page, ou supprimer le message.

# Inventaire

**Trigger :** Command

**Usage :** `$inv (optional page)`

> Affiche l'inventaire du joueur ou du personnage cible.
> Comme pour le magasin, l'utilisateur peut indiquer une page ou utiliser les émoji.

# Achat et vente :

## Sell

**Trigger :** Command

**Usage :** `$sell "name" amount

> Vend l'objet.

## Buy

**Trigger :** Command

**Usage :** `$buy "name" amount`

> Achete l'objet

# Money

**Trigger :** Command

**Usage :** `$money`

> Affiche le porte monnaie de l'utilisateur.

# VN

**Trigger :** Command

**Usage :** `$vn -(add|use) (-armes?|modules?|BC|LC|SF|CU|bc|lc|cb|sf|cu) "value"`

**Indication sur la valeur :**
* <u>Armes & module</u>: nom
* <u> Implant :</u> Quantité

> La fonction va mettre à jour le message de l'inventaire à chaque fois que quelqu'un utilise la commande pour éditer cet inventaire.
> Le message de l'inventaire doit **obligatoirement** être fait par YAGPD.

# Craft

**Trigger :** Command

**Usage :** `$(add|use) (#nom) (quantité) "nom"`

Pour Use, quantité peut être `all` ce qui permet d'utiliser tous les objets directement.



Pour les rerolls, il faut ajouter `#nom` avant d'indiquer la quantité et le nom.

> Permet à l'utilisateur de se rajouter à lui ou à un reroll un objet qui ne se trouve pas forcément dans le store. Si l'objet se trouve déjà dans son inventaire, alors il augmente la quantité.
>
> Cette dernière est d'ailleurs optionnelle. Non précisé, elle sera de 1. La commande `$use` permet de retirer l'objet de l'inventaire, en précisant un montant ou non.
>
> Le bot est sensible aux fautes d'orthographes et de frappes. Il est donc possible que vous ayez un affichage doublé pour le même objet si vous avez fait une erreur d''orthographe ou de frappe.
>
> *Il est impossible de vérifier toutes les orthographes possibles pour chaque objet, faites donc attention.*
>
> Le bot n'est pas sensible à la casse : une fonction permet de mettre automatiquement la première lettre en majuscule, c'est ce qui est fait ici.

# Recherche :

**Trigger** : Regex `?search`

**Usage** : BEAUCOUP.
-  **Inventaire** : `?search inventaire`
- **Composant** : `?search composant`
  - Biocomposant : `?search bc`
	- Cellule bionotropique : `?search cb`
  - Cellule cytomorphe : `?search lc`
  - Substrat ferreux : `?search sf`
  - Composant universel : `?search cu`
- **Analgésique** : `?search analgésique`
	- Rixolam : `?search rixolam`
	- Bandage : `?search bandage(s)`
	- Eufyrusant : `?search eufy(rusant)`
	- Implant : `?search "Implant(s) temporaire(s)"`
	- Soma : `?search soma`
	- Injection Eufyrusant : `?search Eufy(rusant)`
	- Sirop de betapropyl : `?search (sirop|betapropyl|("sirop de betapropyl"))`
	- Xenox : `?search xenox`
- **Armes biologiques** : `?search "Arme biologique"`
	-  Grenade Nécrotique : `?search (grenade|grenade nécrotique|nécro)`
	-  Liquide antifongique :`?search (antifongique|liquide|"liquide antifongique")`
	-  Gaz anesthésiant de combat : `?search (anesthésiant|gaz|"gaz anesthésiant")`
	-  Sang Etherique : `?search ("sang éthérique"|sang|éthérique)`
	-  Huile carotoxinique : `?search (huile|carotonixique|caro|"huile carotoxinique")`
	-  Huile digestive : `?search (huile|digestive|"huile digestive"`
- **Armes** : `?search arme`
- **Modules** : `?search module`
- **Implants** : `?search implants`
- **Charges** : `?search charge`
- **Objet à bas coût** : `?search faible`
	-  Les armes : `?search faible arme(s)`
	-  Les armes biologiques : `?search faible "arme bio"`
	-  Les analgésiques : `?search faible analgésique`
  - Pour les objets seuls, ce sont exactement les mêmes variables que pour les objets hauts de gammes.

> Cette commande permet de récupérer la description (et UNIQUEMENT la description) d'un objet.
> Elle permet aussi de récupérer les messages avec les catégories d'objets.
