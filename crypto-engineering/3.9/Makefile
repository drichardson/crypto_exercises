OPENSSL_DIR=$(shell brew --prefix openssl)
CFLAGS=-I "${OPENSSL_DIR}/include" -I libtomcrypt/src/headers
#LDFLAGS="${OPENSSL_DIR}/lib/libcrypto.a" 
LDFLAGS=-L "${OPENSSL_DIR}/lib" -l crypto

all: encrypt-openssl encrypt-libtom

encrypt-openssl: encrypt-openssl.o

encrypt-libtom: libtom encrypt-libtom.o
	$(CC) -o encrypt-libtom encrypt-libtom.o libtomcrypt/libtomcrypt.a

libtom:
	make -C libtomcrypt -j

clean:
	$(RM) encrypt-openssl encrypt-openssl.o encrypt-libtom encrypt-libtom.o
