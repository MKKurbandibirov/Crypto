from random import randrange, getrandbits

class PrimeGenerator:
    def __call__(self, bits: int) -> int:
        return self.generate_prime_number(bits)

    def miller_rabin_test(self, num: int, rounds: int) -> bool:
        if num == 2 or num == 3:
            return True
        if num <= 1 or num % 2 == 0:
            return False
        s = 0
        r = num - 1
        while r & 1 == 0:
            s += 1
            r //= 2
        for _ in range(rounds):
            a = randrange(2, num - 1)
            x = pow(a, r, num)
            if x != 1 and x != num - 1:
                j = 1
                while j < s and x != num - 1:
                    x = pow(x, 2, num)
                    if x == 1:
                        return False
                    j += 1
                if x != num - 1:
                    return False
        return True


    def generate_prime_candidate(self, bits):
        p = getrandbits(bits)
        p |= (1 << bits - 1) | 1
        return p


    def generate_prime_number(self, length=128):
        p = 4
        while not self.miller_rabin_test(p, rounds=128):
            p = self.generate_prime_candidate(length)
        return p

