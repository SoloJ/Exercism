"""Functions for creating, transforming, and adding prefixes to strings."""


def add_prefix_un(word):
    return 'un' + word


def make_word_groups(vocab_words):
    output = vocab_words[0]
    for count in vocab_words[1:]:
        output += ' :: ' + vocab_words[0]+count
    return output


def remove_suffix_ness(word):
    new_word = word[:-4]
    if new_word.endswith('i'):
        new_word = new_word[:-1] + 'y'
    return new_word


def adjective_to_verb(sentence, index):
    word = sentence.split(' ')[index]
    if word.endswith('.'):
        word = word[:-1]
    return word + 'en'
