from rsa_keys import RSAKeys

class RSA_Cypher:
    def __init__(self, keys: RSAKeys) -> None:
        self.keys = keys

    def encryption(self, text: str) -> list:
        while len(text) % 8:
            text += ' '
        bytes_text = bytes(text, 'utf-8')

        self.binary1 = []
        for byte in bytes_text:
            j = 0
            rev = []
            while byte > 0:
                rev.append(byte%2)
                byte //= 2
                j += 1

            while j < 8:
                self.binary1.append(0)
                j += 1
            
            rev.reverse()
            self.binary1.extend(rev)

        cut = self.keys.bits // 4

        i = 0
        cutted = []
        while i < len(self.binary1):
            cutted.append(self.binary1[i:i+cut])
            i += cut

        self.source = []
        self.encrypted = []
        for val in cutted:
            i = 0
            sum = 0
            while i < len(val):
                sum += pow(2, len(val)-i-1)*val[i]
                i += 1
            
            self.source.append(sum)
            self.encrypted.append(pow(sum, self.keys.e, self.keys.n))

        return self.encrypted
    
    def decryption(self) -> str:
        self.decrypted = []
        for val in self.encrypted:
            self.decrypted.append(pow(val, self.keys.d, self.keys.n))

        self.binary2 = []
        for byte in self.decrypted:
            j = 0
            rev = []
            while byte > 0:
                rev.append(byte%2)
                byte //= 2
                j += 1

            while j < self.keys.bits // 4:
                self.binary2.append(0)
                j += 1
            
            rev.reverse()
            self.binary2.extend(rev)

        byte = 0
        cutted = []
        while byte < len(self.binary2):
            cutted.append(self.binary2[byte:byte+8])
            byte += 8

        text = []
        for val in cutted:
            byte = 0
            sum = 0
            while byte < len(val):
                sum += pow(2, len(val)-byte-1)*val[byte]
                byte += 1
            
            text.append(sum)

        return ''.join(map(chr, text))

    def get_encrypted(self) -> str:
        binary = []
        for byte in self.encrypted:
            j = 0
            rev = []
            while byte > 0:
                rev.append(byte%2)
                byte //= 2
                j += 1

            while j < self.keys.bits // 4:
                binary.append(0)
                j += 1
            
            rev.reverse()
            binary.extend(rev)
        
        byte = 0
        cutted = []
        while byte < len(binary):
            cutted.append(binary[byte:byte+8])
            byte += 8

        text = []
        for val in cutted:
            byte = 0
            sum = 0
            while byte < len(val):
                sum += pow(2, len(val)-byte-1)*val[byte]
                byte += 1
            
            text.append(sum)

        return ''.join(map(chr, text))
