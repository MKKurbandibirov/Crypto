from random import randint
from rsa_keys import RSAKeys
import time
import math

# def pollard_attack(n: int) -> int:
#     x, y, d = 2, 2, 1
#     while d == 1:
#         x = (x * x + 1) % n
#         y = (y * y + 1) % n
#         y = (y * y + 1) % n
#         d = math.gcd(abs(x-y), n)
    
#     if d == n:
#         return 0
#     return

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

def attack() -> dict:
    vals = {35:0, 40:0, 45:0, 46:0, 47:0, 48:0, 49:0, 50:0, 51:0, 52:0, 52:0, 53:0}

    for i in range(10):
        i = 46

        keys = RSAKeys(35)
        st_time = time.time()
        p = pollard_attack(keys.n)
        q = int(keys.n/p)
        end_time = time.time()
        vals[35] += end_time-st_time

        print(f"For bit length = {35} attack took {end_time-st_time}s")

        keys = RSAKeys(40)
        st_time = time.time()
        p = pollard_attack(keys.n)
        q = int(keys.n/p)
        end_time = time.time()
        vals[40] += end_time-st_time

        print(f"For bit length = {40} attack took {end_time-st_time}s")

        keys = RSAKeys(45)
        st_time = time.time()
        p = pollard_attack(keys.n)
        q = int(keys.n/p)
        end_time = time.time()
        vals[45] += end_time-st_time

        print(f"For bit length = {45} attack took {end_time-st_time}s")

        while i <= 53:
            keys = RSAKeys(i)

            st_time = time.time()
            p = pollard_attack(keys.n)
            q = int(keys.n/p)
            end_time = time.time()
            vals[i] += end_time-st_time

            print(f"For bit length = {i} attack took {end_time-st_time}s")

            i += 1
    
    return vals
