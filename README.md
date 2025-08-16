# Encryption/decryption command line tool

Command line tool to encrypt/decrypt files using AES-128 in GCM mode.

Requires Go 1.23.

To build, run `make` in the source folder, it will create the command line tool `bin/cr`

Usage:

    cr encrypt <file> <passphrase>
    cr decrypt <file.enc> <passphrase>

