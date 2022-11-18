def to_rna(dna_strand):
    complement_map = {
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U'
    }
    output = ''
    for x in dna_strand:
        output += complement_map[x]
    return output
