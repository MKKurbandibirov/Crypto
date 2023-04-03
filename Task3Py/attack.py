from random import randint
from rsa_keys import RSAKeys
import time
import math


def pollard_attack(n: int) -> int:
    x0 = randint(1, n-2)
    k = 1
    gcd = 0
    for _ in range(1, n):
        x = []
        x.append(x0)     
        i = 1
        for _ in range((2 ** k) + 1, 2 ** (k + 1) + 1):
            x.append(((x[i - 1] ** 2) + 1) % n)
            gcd = math.gcd(n, abs(x[0] - x[i]))

            if gcd > 1:
                return gcd
            i += 1
        x0 = x[i - 1]
        k += 1
        del x

def attack() -> None:
    i = 35
    while True:
        keys = RSAKeys(i)

        st_time = time.time()
        p = pollard_attack(keys.n)
        q = int(keys.n/p)
        end_time = time.time()

        print(f"For bit length = {i} attack took {end_time-st_time}s")

        if end_time-st_time < 60:
            i += 5
        else:
            i += 1
