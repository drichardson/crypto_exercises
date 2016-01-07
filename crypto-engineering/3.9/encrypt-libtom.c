#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <tomcrypt.h>

unsigned char userKey[] = {
    0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01
};

unsigned char plaintext[] = {
  0x29, 0x6c, 0x93, 0xfd, 0xf4, 0x99, 0xaa, 0xeb, 0x41, 0x94, 0xba, 0xbc,
  0x2e, 0x63, 0x56, 0x1d
};

int main(int argc, char** argv) {
    symmetric_key key;
    int rc;
    

    unsigned char ciphertext[sizeof(plaintext)];
    memset(ciphertext, 0, sizeof(ciphertext));


    rc = rijndael_setup(userKey, sizeof(userKey), 0, &key);
    if (rc != CRYPT_OK) {
        fprintf(stderr, "ERROR: rijndael_setup returned %d.\n", rc);
        exit(1);
    }

    rc = rijndael_ecb_encrypt(plaintext, ciphertext, &key);
    if (rc != CRYPT_OK) {
        fprintf(stderr, "ERROR: rijndael_ecb_encrypt returned %d.\n", rc);
        exit(1);
    }

    rijndael_done(&key);

    for(int i = 0; i < sizeof(ciphertext); i++) {
        putchar(ciphertext[i]);
    }

    return 0;
}

