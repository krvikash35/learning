

* confidential
* Integrity
* Authentication
* Anti-Replay

Encryption method
* symmetric
* asymetric
* hashing


Hashing Algo
* Formulae that on being applied to given input produces an output. `input-->formulae=digest`
* It is used for integrity checks i.e message has not been changed or altered.
* Sender uses secret key along with msg to create hash digest then it sents msg & digest. Reciever recreate hash digest using secret key & message. If its matches then msg has not been altered and msg has been sent by right party that is having secret key.
* Algo should exihibit below characteristic for industry standared
    * Mathematcally almost impossible to extract original msg from digest. Also know as `one-way encryption`
    * Even a minor change in input should cause digest to change
    * Digest size should alwasy remain constant, it should not grow/srink as per size or input.
Digest Lengh
```
MD5     128Bits
SHA1    160Bits
SHA384  384Bits
SHA256  256Bits
```

```bash
echo "hello world" | md5sum
echo "hello world" | sha1sum
```

Symetric encryption Algorithm
* Both party uses same secret key
* Big question is how to trasnfer key from one party to other, answer is `key exchange alog`
* Less CPU intensive, encrypted data size does not change, so good for large size of data.

encryption and key sizes
```
DES 56  56Bits
3DES    168Bits
AES     128Bits
AES192  192Bits
AES256  256Bits
```

Asymetric encryption Algorithm
* Both party has its own two key, private & public
* Private key alwasy lies with sender while it share the public key reciever.
* Sender uses reciever's public key to encrypt data, reciever uses its own private key to decrypt data.
* High CPU intensive, encrypted data size increase, so good for small size of data. Used in symmetric key exchange, signature(hasing)

Authentication
* username & password
* pre shared keys
* digital certificate. SSL/TLS used in https
    * it contain public key of asymmetric key value pair and only one party can contain its associated private key.
    * browser uses it to verify if the authenticity of site. it encrypt random text with site's public key and send it to site server. server decrypt using its private key and send decrypted text back to browser. Browser compare returned text if its same then pass.