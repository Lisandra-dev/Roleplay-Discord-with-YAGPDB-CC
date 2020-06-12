# Count :
**Trigger** : Regex `.*`

**Usage** : /
> The bot count the message sended by user in the channel, increment after some number and post an embed when the count is equal to the value. You can change the number of message and add some cycle if you want.

# Time :
**Trigger** : Command

**Usage** : `-time`
> The bot send an embed associated at the time. It also indicates the number of messages sent during the cycle, in pourcent.
> If you change the number of message for changing the cycle, you can create some calcul to indicate a real pourcentage. For example, for 200, you can divise by 2 the amount of message. For 300, 3, etc. Real number of message are in the footer.
> Also, the time message is send where you trigger the command !

# Reset :
**Trigger** : Command

**Usage** : `-reset`
> Reset the timing count.

# Changing the channel name with the cycle number :
**Trigger** : Minute interval : 5 minutes.

**Usage** : /

> Don't change the time, because discord block the bot changing a channel name under 5 min : it changes change each 5 min.
