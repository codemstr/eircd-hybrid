# eircd-hybrid

May 7, 2015

This update comes with Evo1 & Evo2 password support for each game.
User count file for udpserv.
Fixed a few bugs with server linking.

TODO List:
Fix ban timing system.
Create udpserv as a module.
Intergrate a securtiy check that client contacted udp server before registering on the ircd.
Create a dynamically generated ircd.motd (Update news, new track files, updates, etc)

-sponji


Compile with options.

./configure --with-nicklen=12 --prefix='/home/user/eircd' && make && make install

Edit your etc/ircd.conf

Launch bin/ircd



