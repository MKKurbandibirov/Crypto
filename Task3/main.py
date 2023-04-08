from rsa_keys import RSAKeys
from cypher import RSA_Cypher
from attack import attack
from utils import *


def task1(bits:int = 128) -> None:
    keys = RSAKeys(bits)

    write_keys(keys)

def task2(bits:int = 128) -> None:
    keys = RSAKeys(bits)
    cypher = RSA_Cypher(keys)

    f = open('text.txt', 'r')
    text = f.read()
    f.close()

    cypher.encryption(text)
    print('--- Encrypted text ---')
    print(cypher.get_encrypted())
    print('--- Decrypted text ---')
    print(cypher.decryption())

    write_txt(cypher)


def task3() -> None:
    attack()

if __name__ == '__main__':
    print(task3())
