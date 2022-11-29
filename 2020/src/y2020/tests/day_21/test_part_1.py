from src.y2020.day_21.part_1 import (
    get_all_allergenes,
    get_all_ingredients,
    parse_allergenes,
    parse_food,
    parse_food_list,
    parse_ingredients,
    solution,
)


EXAMPLE = """mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)"""


def test_parse_ingredients():
    ingredients = 'abc def ghi lmno'
    assert parse_ingredients(ingredients) == {'abc', 'def', 'ghi', 'lmno'}


def test_parse_allergenes():
    allergenes = 'abc, def'
    assert parse_allergenes(allergenes) == {'abc', 'def'}
    allergenes = 'abc'
    assert parse_allergenes(allergenes) == {'abc'}


def test_parse_food():
    food = 'abc defg (contains 123, 4567)'
    ingredients, allergenes = parse_food(food)
    assert ingredients == {'abc', 'defg'}
    assert allergenes == {'123', '4567'}


def test_parse_food_list():
    food_list = parse_food_list(EXAMPLE.splitlines())
    assert food_list[0].ingredients == {'mxmxvkd', 'kfcds', 'sqjhc', 'nhms'}
    assert food_list[0].allergenes == {'dairy', 'fish'}


def test_solution():
    assert solution(EXAMPLE) == 5


def test_get_all_ingredients():
    food_list = parse_food_list(EXAMPLE.splitlines())
    assert get_all_ingredients(food_list) == {'mxmxvkd', 'kfcds', 'sqjhc', 'nhms', 'trh', 'fvjkl', 'sbzzf'}


def test_get_all_allergenes():
    food_list = parse_food_list(EXAMPLE.splitlines())
    assert get_all_allergenes(food_list) == {'dairy', 'fish', 'soy'}
