#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <tomcrypt.h>

unsigned char userKey[] = {
    0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01
};

unsigned char ciphertext[] = {
    0x53, 0x9b, 0x33, 0x3b, 0x39, 0x70, 0x6d, 0x14, 0x90, 0x28, 0xcf, 0xe1,
    0xd9, 0xd4, 0xa4, 0x07
};

int main(int argc, char** argv) {
    symmetric_key key;
    int rc;
    

    unsigned char plaintext[sizeof(ciphertext)];
    memset(plaintext, 0, sizeof(plaintext));


    rc = rijndael_setup(userKey, sizeof(userKey), 0, &key);
    if (rc != CRYPT_OK) {
        fprintf(stderr, "ERROR: rijndael_setup returned %d.\n", rc);
        exit(1);
    }

    rc = rijndael_ecb_decrypt(ciphertext, plaintext, &key);
    if (rc != CRYPT_OK) {
        fprintf(stderr, "ERROR: rijndael_ecb_decrypt returned %d.\n", rc);
        exit(1);
    }

    rijndael_done(&key);

    for(int i = 0; i < sizeof(plaintext); i++) {
        putchar(plaintext[i]);
    }

    return 0;
}

