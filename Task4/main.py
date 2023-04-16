import string
import time
import random
from utils import *
from md4_hashing import MD4


def FillTo512(binary: str) -> str:
    len_in_bin = bin(len(binary))[2:]
    binary += '1'

    while len(binary) % 512 != 448:
        binary += '0'

    binary = ''.join(InvertBits(binary[len(binary)-448:], 8)[::-1])
    binary = ''.join(InvertBits(binary, 32)[::-1])

    len_in_bin = '0' * (64 - len(len_in_bin)) + len_in_bin
    len_in_bin = ''.join(InvertBits(len_in_bin, 32)[::-1])

    return binary + len_in_bin


def GetHashMD4(binary: str):
    binaryBy512 =  FillTo512(binary)
    binaryBlocks = InvertBits(binaryBy512,512)
    hash = MD4(binaryBlocks)
    return hash


def Task2(binary: str, k: int, hash1: str) -> None:
    binaryChange = list(binary)

    for i in range(0,k):
        binaryChange[i] = str(int(binary[i])^1)

    binaryChange = ''.join(binaryChange)
    print(f'Changed {k} bits on the left: {binaryChange} ')
    
    hash2 = GetHashMD4(binaryChange)
    print(f'Changed hash: {hash2} ')
    
    count1 = bin(int(hash1,16) ^ int(hash2,16)).count('1')
    print(f'Changed bits in hash: {count1}')
    print('\n')


def GenerateRandomString(length: int) -> str:
    letters = string.ascii_lowercase
    randomString = ''.join(random.choice(letters) for i in range(length))
    return randomString


def Task3(k: int, L: int) -> None:
    hash_list=list()
    words=list()
    word=GenerateRandomString(L)
    start = time.time()
    hash=hex(int(bin(int(GetHashMD4(GetBin(word)),base=16))[2:k+2],2))[2:]

    while True :
        if hash in hash_list and word not in words:
            break
        hash_list.append(hash)       
        words.append(word)
        word=GenerateRandomString(L)
        hash=hex(int(bin(int(GetHashMD4(GetBin(word)),base=16))[2:k+2],2))[2:]

    i = 0
    while hash != hash_list[i]:
        collision_hash=hash_list[i]
        collision_word=words[i]

    print (f'M = {word}')
    print (f"M' = {collision_word}")
    print (f'h = {hash}')
    print (f"h' = {collision_hash}")
    print(f"N =  {len(words)}")
    print(f'Time: {time.time()-start}')
    print('\n')


def Task4(k: int, password: str) -> None:
    hash_pass = hex(int(bin(int(GetHashMD4(GetBin(password)), 16))[2:k+2], 2))[2:]
    print(f'Hash M: {hash_pass}')
    kol = 0

    while True:
        word = GenerateRandomString(len(password))
        hash = hex(int(bin(int(GetHashMD4(GetBin(word)), 16))[2:k+2], 2))[2:]
        kol += 1
        if hash == hash_pass and password != word:
            print(f"M`: {word}\nhash прообраза: {hash}")
            print(f"Количество попыток найти прообраз= {kol}")
            break

    print('\n')


if __name__ == '__main__':
    # file = open("text.txt")
    # text = file.read()
    # file.close()

    # binary = GetBin(text)
    
    # result = []
    # for i in range(0,len(binary), 128):
    #     result.append(binary[i:i+128])
    
    # res = ''
    # for i in result:
    #     res += bin(int(GetHashMD4(i), 16))[2:]
    
    # out = open('out.txt', 'w')
    # out.write(res)
    # out.close()
    
    # print(len(res))
    
    print("------ Task 2 ------")
    for k in range(1, 20):
        binary = GetBin('abc')
        hash = GetHashMD4(binary)
        Task2(binary=binary, k=k, hash1=hash)

    # print("------ Task 4 ------")
    # for k in range(1, 25):
    #     Task4(k=k, password='Beyond the bound of reasons!')

    # print("------ Task 6 ------")
    # for k in range(8):
    #     binary = GetBin(GenerateRandomString(10**k))
    #     st_time = time.time()
    #     hash = GetHashMD4(binary)
    #     print(f"Time for {10**k}: {time.time()-st_time}")
        

