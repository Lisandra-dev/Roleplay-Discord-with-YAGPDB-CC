Thanks a lot to Shadownight#9025 to permit me to use it !
Each command for buy, sell, inventory and money can be use with a reroll with `#name`at **the end of the command**.
For admin, user can be reroll, with adding the `#` at the User place.

All these command use dictionnary and cslice, unless of store open/close, who use an simple db entry. 

# ADMIN

**Trigger : ** Command

**Usage :**

* Money : `$admin -money -(add|remove|reset|set) user amount`

* Symbol : `$admin -symbol <symbol>`

* Manage inventory : `$admin -manageinv -(add|remove|set) user item amount`

* Reset/clean inventory : `$admin -manageinv -(reset|clean) user`

  > - Reset détruit complètement l'inventaire du joueur, quand clean le vide seulement.



# Item :

**Trigger : ** Command

**Usage :**

* Add : `$item -add "name" price (none|sellprice) (inf|stockage)(true|false)"description"`
* Delete : `$item -delete "name"`
* Information : `$item -info "name"`

> - Beware of case sensitive
>
> - SII meaning that the item can be in the inventory or not, like an "auto-use item"

# Edit item :

**Trigger :** Regex

**Usage :** `$item -edit "name" -(price|sell|stock|sii|desc|rbuy|rsell) "value" `

* For price : Edit the price of the item, value are number.
* Sell can be 0 or number.
* Stock can be inf or number
* SII : true or false
* desc : string
* rbuy and rsell rand a number between two numbers, they need to have two argument.

# Store :

**Trigger :** Command

**Usage :** `$store (optional page)`

> Print the embed of the store.
>
> If store are close (with `$setup (close|open)`print the embed. Else, print an embed where the store is closed.
>
> If the dictionary have too many entry, user can choose a page.

# Inventory :

**Trigger :** Command

**Usage :** `$inv (optional page)`



# Sell :

**Trigger : ** Command

**Usage :** `$sell "name" amount`



# Buy :

**Trigger :** Command

**Usage :** `$buy "name" amount`



# Money :

**Trigger :** Command

**Usage : ** `$money`

> Print the balance of the user.

# Open - Close :

**Trigger : ** Command

**Usage :** `$setup (open|close)`

> Open or close the shop. If the shop is closed, user cannot buy or sell anything.

# VN :

**Trigger :** Command

**Usage : ** `$vn -(add|use) (-armes?|modules?|BC|LC|SF|CU|bc|lc|cb|sf|cu) "value"`

> For armes and module : Value is the name
>
> For the implant, value is the quantity.
>
> This function update an embed created with YAG and update & edit it each time someone edit the inventory of the guild.
