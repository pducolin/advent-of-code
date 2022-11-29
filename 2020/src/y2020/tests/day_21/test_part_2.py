from src.y2020.day_21.part_2 import solution

EXAMPLE = """mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)"""


def test_solution():
    assert solution(EXAMPLE) == 'mxmxvkd,sqjhc,fvjkl'
