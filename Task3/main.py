from rsa_keys import RSAKeys
from cypher import RSA_Cypher
from attack import attack, pollard_attack
from utils import *
from generator import PrimeGenerator
import time


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
    cypher.get_encrypted()

    # print(len(cypher.binary1), len(cypher.binary) / len(cypher.binary1))

    # write_txt(cypher)


def task3() -> None:
    attack()

def task4() -> None:
    r = 0.25
    while r <= 0.6:
        gen = PrimeGenerator()

        p = gen.generate_prime_number(int(100*r))
        q = gen.generate_prime_number(int(100*(1-r)))

        st_time = time.time()
        pollard_attack(p*q)
        print(f"r = {r}: time - {time.time()-st_time}")

        r += 0.025



if __name__ == '__main__':
    task2()
