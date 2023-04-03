from generator import PrimeGenerator
import math

class RSAKeys:
    def __init__(self, bits=128) -> None:
        generator = PrimeGenerator()
        self.bits = bits

        while True:
            self.p = generator(bits)+1
            self.q = generator(bits)
            self.n = self.p * self.q
            self.fi = (self.p-1) * (self.q-1)
            self.e = 65537
            self.d = y = pow(self.e, -1, self.fi)

            if math.gcd(self.e, self.d) == 1:
                break
