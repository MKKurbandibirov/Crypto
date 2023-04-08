from rsa_keys import RSAKeys
from cypher import RSA_Cypher

def write_keys(keys: RSAKeys) -> None:
    private_f = open('private.txt', 'w')
    private_f.write('P: ' + str(keys.p)+'\n')
    private_f.write('Q: ' + str(keys.q)+'\n')
    private_f.write('D: ' + str(keys.d)+'\n')
    private_f.write('N: ' + str(keys.n)+'\n')
    private_f.close()

    public_f = open('public.txt', 'w')
    public_f.write('E: ' + str(keys.e)+'\n')
    public_f.write('N: ' + str(keys.n)+'\n')
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

    en_text_f = open('en_text.txt', 'w')
    en_text_f.write(cypher.get_encrypted())
    en_text_f.close()