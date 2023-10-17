# Encryption and Decryption

Encryption is the process by which a readable message is converted to an unreadable from to prevent unauthorized parties from reading it. Decryption is the process of converting an encrypted message back to its original (readable) format. The original message is called the plaintext message. Then encrypted message is called the ciphertext message.  

Digital encryption algorithm work by manipulating the digital context of a plaintext message mathematically, using an encryption algorithm and a digital key to produce a ciphertext version of the message. The sender and recipient can communicate securely if the sender and recipient are the only ones who know the key. 

## Shared Key and Public Key Encryption

**Shared Key Encryption**

Shared key encryption uses one key to encrypt and decrypt messages. For shared key cryptography to work, the sender and the recipient of a message must both have the same key, which must keep secret from everybody else. Shared key encryption/decryption is relatively fast. However, since any one with the shared key can decrypt the information, shared key encryption requires that only the sender and recipient have access to the shared key.

**Public Key Encryption**

Public key encryption uses a pair of complementary keys (a public key and a private key) to encrypt and decrypt messages. The two keys are mathematically related such that a message encoded with one key only be decoded with the other key. Although a user's public and private keys are mathematically related  knowledge of a public key does not make it possible to calculate the corresponding private key.

In public key encryption systems, users make their public ey available to anyone and keep their private key secret. When user wants to send a private message to another user, the sender look up the recipient's public key and uses it to encrypt a message.

When the encrypted message arrives, the recipient uses his or her private key to decrypt the message.

Public key encryption algorithms are mathematically more complex that shared key encryption algorithm. As a result, public key encryption is significantly shower that shared key encryption.
