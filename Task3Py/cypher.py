from rsa_keys import RSAKeys

class RSA_Cypher:
    def __init__(self, keys: RSAKeys) -> None:
        self.keys = keys

    def encryption(self, text: str) -> list:
        bytes_text = bytes(text, 'utf-8')

        binary = []
        for byte in bytes_text:
            j = 0
            rev = []
            while byte > 0:
                rev.append(byte%2)
                byte //= 2
                j += 1

            while j < 12:
                binary.append(0)
                j += 1
            
            rev.reverse()
            binary.extend(rev)
        
        print(binary)

        cut = self.keys.bits // 4

        i = 0
        cutted = []
        while i < len(binary):
            cutted.append(binary[i:i+cut])
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
        
        binary = []
        for byte in self.decrypted:
            j = 0
            rev = []
            while byte > 0:
                rev.append(byte%2)
                byte //= 2
                j += 1

            while j < 12:
                binary.append(0)
                j += 1
            
            rev.reverse()
            binary.extend(rev)
        
        print(binary)

        i = 0
        cutted = []
        while i < len(binary):
            cutted.append(binary[i:i+12])
            i += 12
        

        text = []
        for val in cutted:
            i = 0
            sum = 0
            while i < len(val):
                sum += pow(2, len(val)-i-1)*val[i]
                i += 1
            
            text.append(sum)

        return ''