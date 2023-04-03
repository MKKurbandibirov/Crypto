from rsa_keys import RSAKeys
from cypher import RSA_Cypher

def write_keys(keys: RSAKeys) -> None:
    private_f = open('private.txt', 'w')
    private_f.write(str(keys.p)+'\n')
    private_f.write(str(keys.q)+'\n')
    private_f.write(str(keys.d)+'\n')
    private_f.write(str(keys.n)+'\n')
    private_f.close()

    public_f = open('public.txt', 'w')
    public_f.write(str(keys.e)+'\n')
    public_f.write(str(keys.n)+'\n')
    public_f.close()

def write_txt(cypher: RSA_Cypher) -> None:
    encrypted_f = open('encrypted.txt', 'w')
    for val in str(cypher.encrypted).strip('[]').split(', '):
        encrypted_f.write(str(val)+'\n')
    encrypted_f.close()

    decrypted_f = open('decrypted.txt', 'w')
    for val in str(cypher.decrypted).strip('[]').split(', '):
        decrypted_f.write(str(val)+'\n')
    decrypted_f.close()