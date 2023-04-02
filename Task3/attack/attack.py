from random import randint
import time

def GCD(a, b: int) -> int:
    while b != 0:
        t = a % b
        a, b = b, t
    return a


def Pollard_Attack(n: int) -> int:
    x0 = randint(1, n-2)
    k = 1
    gcd = 0
    for _ in range(1, n):
        x = list()
        x.append(x0)     
        z = 1
        for _ in range((2 ** k) + 1, 2 ** (k + 1) + 1):
            x.append(((x[z - 1] ** 2) + 1) % n)
            gcd = GCD(n, abs(x[0] - x[z]))

            if gcd>1:
                return gcd
            z += 1
        x0 = x[z - 1]
        k += 1
        del x


if __name__ == '__main__':
    file = open('N.txt', 'r')
    ns = file.read().split('\n')
    Ns = list(map(int, ns[:len(ns)-1]))
    file.close()

    for N in Ns:
        st_time = time.time()
        P = Pollard_Attack(N)
        Q = int(N/P)
        end_time = time.time()

        # print(f"N: {N}\nP: {P}\nQ: {Q}")

        print(f"Work time: {end_time-st_time}s")
