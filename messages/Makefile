# $Id: Makefile.in 831 2006-12-19 15:52:19Z michael $
#
RM		= /bin/rm
MV		= /bin/mv
INSTALL		= /usr/bin/install -c
INSTALL_DATA	= ${INSTALL} -m 644
prefix		= $(DESTDIR)/home/pi/eircd
messagedir	= $(prefix)/messages
datarootdir	= $(DESTDIR)${prefix}/share

CLOBBER		= 

LANGFILES     = ayb.lang               \
                custom.lang            \
		ircd-brazilian_pt.lang \
                ircd-croatian.lang     \
                ircd-danish.lang       \
                ircd-dutch.lang        \
		ircd-french.lang       \
                ircd-german.lang       \
                ircd-italian.lang      \
                ircd-norwegian.lang    \
                ircd-polish.lang       \
                ircd-russian.lang      \
                ircd-spanish.lang      \
                ircd-swedish.lang

all:

install:
	mkdir -p $(messagedir)
	-@echo    "Installing language files to $(messagedir)...";          \

	@for i in $(LANGFILES); do                                          \
		if test -f $(messagedir)/$$i -a -z "$(CLOBBER)"; then       \
			echo $(MV) $(messagedir)/$$i $(messagedir)/$$i.old; \
			$(MV) $(messagedir)/$$i $(messagedir)/$$i.old;      \
		fi;                                                         \
		echo $(INSTALL_DATA) $$i $(messagedir);                     \
		$(INSTALL_DATA) $$i $(messagedir);                          \
	done

distclean:
	${RM} -f Makefile

.PHONY: install distclean
