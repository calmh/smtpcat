smtpcat: $(wildcard *.go)
	gccgo -static-libgo -static-libgcc -lnsl -lsocket -o $@ $^

