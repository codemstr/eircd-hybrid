# $Id: Makefile.in 715 2006-07-14 21:01:55Z michael $
INSTALL		= /usr/bin/install -c
INSTALL_DATA	= ${INSTALL} -m 644
RM		= /bin/rm
prefix		= $(DESTDIR)/home/pi/eircd
datarootdir	= $(DESTDIR)${prefix}/share
mandir          = $(DESTDIR)${datarootdir}/man

all:

install:
	mkdir -p $(mandir)/man8
	$(INSTALL_DATA) ircd.8 $(mandir)/man8/

distclean:
	${RM} -f Makefile

.PHONY: install distclean
