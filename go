
#!/user/bin/perl
make distclean
./configure --prefix=$HOME/ircd --with-nicklen=12
make
make install
