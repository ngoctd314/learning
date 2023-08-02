# RSA Encryption

Reference: https://www.youtube.com/watch?v=wXB-V_Keiu8

RSA is asymmetric algorithm, which is just another way to say "one-way"

RSA works by generating a public and a private key. The public and private keys are generated together and form a key pair.

The public key can be used to encrypt any arbitrary piece of data, but cannot decrypt it.

The private key can be used to decrypt any piece of data that was encrypted by it's corresponding public key.

This means we can give our public key to whoever we want. They can then encrypt any information they want to send us, and the only way to access this information is by using our private key to decrypt it.

## Key generation
