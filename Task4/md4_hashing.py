from utils import InvertBits


def BitwiseLefCase(a: int, s: int) -> int:
    return (a << s) | (a >> (32 - s))


def MethodF(x: int, y: int, z: int) -> int:
    return (x & y | ((~x) & z))


def MethodG(x: int, y: int, z: int) -> int:
    return (x & y) | (x & z) | (y & z)


def MethodH(x: int, y: int, z: int) -> int:
    return (x ^ y ^ z)


def Round1(A: int, B: int, C: int, D: int, Xk: int, s: int) -> int:
    return BitwiseLefCase((A + MethodF(B, C, D) + Xk) % (2**32), s)


def Round2(A: int, B: int, C: int, D: int, Xk: int, s: int) -> int:
    return BitwiseLefCase((A + MethodG(B, C, D) + Xk + 0x5A827999) % (2**32), s)


def Round3(A: int, B: int, C: int, D: int, Xk: int, s: int) -> int:
    return BitwiseLefCase((A + MethodH(B, C, D) + Xk + 0x6ED9EBA1) % (2**32), s)


def MD4(R: list) -> str:
    N = len(R)

    # A = 0x67452301
    # B = 0xefcdab89
    # C = 0x98badcfe
    # D = 0x10325476

    A = 0x67452301
    B = 0xefcdab89
    C = 0x98badcfe
    D = 0x10325476

    X = list()
    for i in range(0, N):
       
        word=InvertBits(R[i],32)
        X.clear()

        for j in word:
           X.append(int(j,2))

        AA = A
        BB = B
        CC = C
        DD = D

        A = Round1(A, B, C, D, X[0], 3);  D = Round1(D, A, B, C, X[1], 7); C = Round1(C, D, A, B, X[2], 11);  B = Round1(B, C, D, A, X[3], 19)
        A = Round1(A, B, C, D, X[4], 3);  D = Round1(D, A, B, C, X[5], 7); C = Round1(C, D, A, B, X[6], 11);  B = Round1(B, C, D, A, X[7], 19)
        A = Round1(A, B, C, D, X[8], 3);  D = Round1(D, A, B, C, X[9], 7); C = Round1(C, D, A, B, X[10], 11); B = Round1(B, C, D, A, X[11], 19)
        A = Round1(A, B, C, D, X[12], 3); D = Round1(D, A, B, C, X[13], 7);C = Round1(C, D, A, B, X[14], 11); B = Round1(B, C, D, A, X[15], 19)

        # A = Round2(A, B, C, D, X[0], 3); D = Round2(D, A, B, C, X[4], 5);C = Round2(C, D, A, B, X[8], 9); B = Round2(B, C, D, A, X[12], 13)
        # A = Round2(A, B, C, D, X[1], 3); D = Round2(D, A, B, C, X[5], 5);C = Round2(C, D, A, B, X[9], 9); B = Round2(B, C, D, A, X[13], 13)
        # A = Round2(A, B, C, D, X[2], 3); D = Round2(D, A, B, C, X[6], 5);C = Round2(C, D, A, B, X[10], 9);B = Round2(B, C, D, A, X[14], 13)
        # A = Round2(A, B, C, D, X[3], 3); D = Round2(D, A, B, C, X[7], 5);C = Round2(C, D, A, B, X[11], 9);B = Round2(B, C, D, A, X[15], 13)

        # A = Round3(A, B, C, D, X[0], 3); D = Round3(D, A, B, C, X[8], 9); C = Round3(C, D, A, B, X[4], 11); B = Round3(B, C, D, A, X[12], 15)
        # A = Round3(A, B, C, D, X[2], 3); D = Round3(D, A, B, C, X[10], 9);C = Round3(C, D, A, B, X[6], 11); B = Round3(B, C, D, A, X[14], 15)
        # A = Round3(A, B, C, D, X[1], 3); D = Round3(D, A, B, C, X[9], 9); C = Round3(C, D, A, B, X[5], 11); B = Round3(B, C, D, A, X[13], 15)
        # A = Round3(A, B, C, D, X[3], 3); D = Round3(D, A, B, C, X[11], 9);C = Round3(C, D, A, B, X[7], 11); B = Round3(B, C, D, A, X[15], 15)
                
        A = (A + AA) % (2 ** 32)
        B = (B + BB) % (2 ** 32)
        C = (C + CC) % (2 ** 32)
        D = (D + DD) % (2 ** 32)
        
    A = ''.join(InvertBits(hex(A)[2:],2)[::-1])
    B = ''.join(InvertBits(hex(B)[2:],2)[::-1])
    C = ''.join(InvertBits(hex(C)[2:],2)[::-1])
    D = ''.join(InvertBits(hex(D)[2:],2)[::-1])
    hash = A + B + C + D

    return  hash