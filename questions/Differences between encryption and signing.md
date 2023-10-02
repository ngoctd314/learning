# What are the differences between encryption and signing? Why should you use digital signatures?

## What does encryption mean & how it works?

Encryption is a process where data in the form of plaintext is transformed into ciphertext. The ciphertext is created using the encryption algorithm where all the patterns of the plaintext are recreated into a new format of the text that can only be deciphered using a private key.

This key ensures that the data is only decrypted by the authorized recipient. In simple terms, the sender of the data will use the public key to write and send an encrypted message while the recipient uses the private key to read the data.

Encryption comes into different types: Symmetric and Asymmetric encryption. Asymmetric is the most secure way to protect sensitive information as it involves creating a pair of public and private keys. The private key is kept secret while the public key is available to everyone to decrypt the sent data.

In symmetric, encryption, there's only one key to encrypt and decrypt the sent data, which makes the information vulnerable to attacks. Using encryption technology, business can encrypt their emails, documents, user payment information, and other sensitive data.

## What does signing mean & how it works?

The signing process is where the sender of the data is authenticated to ensure the data is coming from a legitimate source and doesn't involve malicious code. It also ensure that data received hasn't been altered or attacked during the transit.

In simple terms, the signing process involves digital signatures to bind the source with the data while sending it. Here, the sender uses the private key to write and send the data and the receiver use the public key to check if the data is from the mentioned source or not. 

Thus, when a hacker tries to compromise the data with malicious injection, the public key won't match the private key, letting the recipient know not to trust the source.

The digital signing process is somewhat opposite to how encryption works. Here, the hashing algorithm and the sender's private key are used for digitally signing the data. The process creates the hash digest that can only be recreated using one of the keys from the key pair of the sender.

The sender, then, sends the data, the hash digest, and the public key to the recipient. Now the recipient will use this public key to hash received information, which should match the hash digest sent along with the data. If it matches successfully, the identity of the sender is then confirmed along with the data which ensure it hasn't been modified. 

## Differences between encryption and signing

|Encryption|Signing|
|-|-|
|Encryption technology is used to encoding sensitive information sent in an email or document|Signing is used for verifying the identity of the source of the information sent over the internet|
|There are two encryption methods: Asymmetric and Symmetric|Signing uses the hashing algorithm to generate a hash digest.|
|The sender uses the public key to encrypt the data, receiver use private key to decrypt data|The sender uses the private key while the public key is used by the receiver to verify|
