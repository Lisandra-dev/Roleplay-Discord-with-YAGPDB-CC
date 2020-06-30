Please note that most of the notes used here are in French, and are based on science fiction RP. However, you just have to change the names and triggers to make them usable for a majority of RP.

Some commands are not gameplay but aim to improve life on the server.



Thus, we have :
- The "Ungrouped" folder, which contains the "funs" and help commands.
- The note-taking folder, taken entirely from the YAGPDB server. It is composed of a note system for the whole server, and a personal note system whose answers are sent as private messages.
- The weapon system, which must be taken in its entirety, operates on the basis of message length and number. Weapons reload from a certain number of messages sent, and as soon as the correct trigger is used, the weapon will discharge. This makes it possible to count the number of "bullets" without getting confused. You can change the names of the weapons, bullets, messages and lengths.
- Time is one of the most complicated commands that had to be developed. On the whole RP part of the server, it takes into account the number of messages sent and increments a counter. The bot then sends a message in embed with an image symbolizing the time (or more exactly, the cycle) of the day. Here, we have chosen to separate into 4, but it can be easy to increase the number of cycles. It takes 100 messages to change cycle, and a command allows to get the number of messages before the change of cycle.
- The dice system is the new kid on the block! Globally, the system pulls an integer between 1 and 10, and it can remove numbers (passed as argument) or add numbers (passed as argument too). It also takes into account possible comments (type of action...) that it will simply display, with the player's nickname.
In addition, there is also a whole system of "malus" "poison" and "care" which works in the same way but which will actually display the effect directly.
The results will be displayed as a fairly simple embed, with just a color and a description.

In addition, with the "Set" command, you can set the thresholds of people's characteristics and use them to directly indicate whether there was a pass/fail in the actions. For the dice, the default value if the guy didn't indicate the characteristics it's 8. 

Attention however, there is a limit in the DB which is 50*UserCount (or 500*UserCount for premium). So check your DB before integrating its programs.
