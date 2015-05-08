# $Id: Makefile.in 715 2006-07-14 21:01:55Z michael $

INSTALL		= /usr/bin/install -c
INSTALL_DATA	= ${INSTALL} -m 644
MV		= /bin/mv
RM		= /bin/rm
TEST		= /usr/bin/test

# no whitespace here, test will mess up
CLOBBER=

prefix		= $(DESTDIR)/home/pi/eircd
sysconfdir	= $(DESTDIR)${prefix}/etc
exec_prefix	= $(DESTDIR)${prefix}
datarootdir	= $(DESTDIR)${prefix}/share

DISTFILES	= example.conf

all:

install:
	-@if $(TEST) $(sysconfdir) -ef `pwd`; then exit; fi; \
	mkdir -p $(sysconfdir); \
	for i in $(DISTFILES); do \
		if test -f $(sysconfdir)/$$i -a -z "$(CLOBBER)"; then \
			echo $(MV) $(sysconfdir)/$$i $(sysconfdir)/$$i.old; \
			$(MV) $(sysconfdir)/$$i $(sysconfdir)/$$i.old; \
		fi; \
		echo $(INSTALL_DATA) $$i $(sysconfdir); \
		$(INSTALL_DATA) $$i $(sysconfdir); \
	done

	@echo "Note: more example configuration files can be found in this directory (etc/)."
	
	@if test ! -f "$(sysconfdir)/ircd.motd"; then \
		echo "Creating generic $(sysconfdir)/ircd.motd"; \
		echo "This is ircd-hybrid MOTD replace it with something better" > $(sysconfdir)/ircd.motd; \
	fi
	
	@for i in d k x rx rk; do \
		if test ! -f "$(sysconfdir)/$${i}line.conf"; then \
			echo touch $(sysconfdir)/$${i}line.conf; \
			touch $(sysconfdir)/$${i}line.conf; \
		fi; \
	done

	@for i in n c; do \
		if test ! -f "$(sysconfdir)/$${i}resv.conf"; then \
			echo touch $(sysconfdir)/$${i}resv.conf; \
			touch $(sysconfdir)/$${i}resv.conf; \
		fi; \
	done

distclean:
	${RM} -f Makefile

.PHONY: install distclean
