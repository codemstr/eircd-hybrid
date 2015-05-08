# $Id: Makefile.in 715 2006-07-14 21:01:55Z michael $

CC		= gcc
INSTALL		= /usr/bin/install -c
INSTALL_BIN	= ${INSTALL}
RM		= /bin/rm
CFLAGS		=  -Wall -O2 -g
CPPFLAGS	= 
LDFLAGS   	=  -Wl,-export-dynamic
MV		= /bin/mv
RM		= /bin/rm
LN              = /bin/ln
CLOBBER		= 

SSL_LIBS	=  -lssl -lcrypto
SSL_INCLUDES	= 
CRYPT_LIB	= -lcrypt

INCLUDES	= -I../include $(SSL_INCLUDES)

prefix		= $(DESTDIR)/home/pi/eircd
exec_prefix	= $(DESTDIR)${prefix}
bindir		= $(DESTDIR)${exec_prefix}/bin
datarootdir	= $(DESTDIR)${prefix}/share

PROGS		= mkpasswd encspeed

all: $(PROGS)

build: all

# We must link these two against special libs
encspeed: ../include/setup.h encspeed.c
	$(CC) $(CFLAGS) $(INCLUDES) $(LDFLAGS) encspeed.c -o encspeed $(SSL_LIBS)

mkpasswd: ../include/setup.h mkpasswd.c
	$(CC) $(CFLAGS) $(INCLUDES) $(LDFLAGS) mkpasswd.c -o mkpasswd $(CRYPT_LIB)

# Default rule for everything

clean:
	$(RM) -f *.o *~ core *.exe $(PROGS)

distclean: clean
	$(RM) -f Makefile

install: build
	mkdir -p $(bindir)
	@for i in $(PROGS); do \
		if test -f $(bindir)/$$i -a -z "$(CLOBBER)"; then \
			echo $(MV) $(bindir)/$$i $(bindir)/$$i.old; \
			$(MV) $(bindir)/$$i $(bindir)/$$i.old; \
		fi; \
		echo $(INSTALL_BIN) $$i $(bindir); \
		$(INSTALL_BIN) $$i $(bindir); \
	done

.PHONY: install distclean clean build
