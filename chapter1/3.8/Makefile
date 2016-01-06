OPENSSL_DIR=$(shell brew --prefix openssl)
CFLAGS=-I "${OPENSSL_DIR}/include" -I libtomcrypt/src/headers
#LDFLAGS="${OPENSSL_DIR}/lib/libcrypto.a" 
LDFLAGS=-L "${OPENSSL_DIR}/lib" -l crypto

all: decrypt-openssl decrypt-libtom

decrypt-openssl: decrypt-openssl.o

decrypt-libtom: libtom decrypt-libtom.o
	$(CC) -o decrypt-libtom decrypt-libtom.o libtomcrypt/libtomcrypt.a

libtom:
	make -C libtomcrypt -j

clean:
	$(RM) decrypt-openssl decrypt-openssl.o decrypt-libtom decrypt-libtom.o
