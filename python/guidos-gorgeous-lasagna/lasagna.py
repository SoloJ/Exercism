"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""
EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


def bake_time_remaining(time_elapsed):
    """
this function returns the time remaining
"""
    return (EXPECTED_BAKE_TIME - time_elapsed)


def preparation_time_in_minutes(number_of_layers):
    """
this function returns the preperation time needed
"""
    return (number_of_layers*PREPARATION_TIME)


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """
this function returns the time elapsed
"""

    return (int(PREPARATION_TIME * number_of_layers + elapsed_bake_time))
