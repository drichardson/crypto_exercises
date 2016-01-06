#include <openssl/aes.h>
#include <openssl/conf.h>
#include <openssl/err.h>
#include <openssl/evp.h>
#include <stdio.h>
#include <string.h>

unsigned char key[] = {
    0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01
};

unsigned char ciphertext[] = {
    0x53, 0x9b, 0x33, 0x3b, 0x39, 0x70, 0x6d, 0x14, 0x90, 0x28, 0xcf, 0xe1,
    0xd9, 0xd4, 0xa4, 0x07
};

int main(int argc, char** argv) {
    unsigned char plaintext[sizeof(ciphertext)];
    memset(plaintext, 0, sizeof(plaintext));

    /*
    // Load the human readable error strings for libcrypto
    ERR_load_crypto_strings();

    // Load all digest and cipher algorithms
    OpenSSL_add_all_algorithms();

    // Load config file, and other important initialisation
    OPENSSL_config(NULL);
    */

    AES_KEY k;
    AES_set_encrypt_key(key, 256, &k);
    //AES_decrypt(ciphertext, plaintext, &k);
    AES_ecb_encrypt(ciphertext, plaintext, &k, AES_DECRYPT);

    for(int i = 0; i < sizeof(plaintext); i++) {
        putchar(plaintext[i]);
    }
    return 0;
}

