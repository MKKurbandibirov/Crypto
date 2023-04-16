def GetBin(text: str) -> str:
    result = ''

    for char in text:
        result += bin(ord(char))[2:].zfill(8)

    return result


def GetInput() -> str:
    text = input ('Enter the text for hashing: ')
    return GetBin(text)


def InvertBits(binary: str, step: int) -> list:
	result = []

	for i in range(0,len(binary),step):
		result.append(binary[i:i+step])

	return result

