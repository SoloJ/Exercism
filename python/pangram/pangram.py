def is_pangram(sentence):

    alphabet = 'abcdefghijklmnopqrstuvwxyz'
    sentence = sentence.lower()
    count = 0
    for x in sentence:
        if x in alphabet:
            alphabet = alphabet.replace(x, '')
    if len(alphabet) == 0:
        return True
    return False
