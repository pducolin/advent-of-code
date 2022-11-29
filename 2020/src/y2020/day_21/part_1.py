from collections import namedtuple

Food = namedtuple('Food', ['ingredients', 'allergenes'])

IngredientAllergene = namedtuple('IngredientAllergene', ['ingredient', 'allergene'])


def parse_ingredients(data):
    ingredients = set()
    for ingredient in data.split(' '):
        ingredients.add(ingredient)
    return ingredients


def parse_allergenes(data):
    allergenes = set()
    for allergene in data.split(','):
        allergenes.add(allergene.strip())
    return allergenes


def parse_food(data):
    if '(contains' not in data:
        # only ingredients
        return parse_ingredients(data), set()
    ingredients_data, allergenes_data = data.split('(contains')
    return parse_ingredients(ingredients_data.strip()), parse_allergenes(allergenes_data[:-1])


def parse_food_list(data):
    food_list = []
    for food_data in data:
        ingredients, allergenes = parse_food(food_data)
        food_list.append(Food(ingredients, allergenes))
    return food_list


def get_all_ingredients(food_list):
    all_ingredients = set()
    for food in food_list:
        for ingredient in food.ingredients:
            all_ingredients.add(ingredient)
    return all_ingredients


def get_all_allergenes(food_list):
    all_allergenes = set()
    for food in food_list:
        for allergene in food.allergenes:
            all_allergenes.add(allergene)
    return all_allergenes


def solution(data):
    food_list = parse_food_list(data.splitlines())
    # find ingredient - allergenes matches
    ingredients_allergenes = []
    ingredients_matched = set()
    allergenes_matched = set()
    all_allergenes = list(get_all_allergenes(food_list))
    potential_matches = {}
    while len(all_allergenes) > 0 or len(potential_matches) > 0:
        current_allergene = all_allergenes.pop(0)
        # find list of food with that allergene
        foods = [food for food in food_list if current_allergene in food.allergenes]
        possible_ingredients = set(foods[0].ingredients)
        for food in foods:
            possible_ingredients &= food.ingredients
        possible_ingredients -= ingredients_matched
        if len(possible_ingredients) == 1:
            # found match
            ingredient = possible_ingredients.pop()
            ingredients_allergenes.append(IngredientAllergene(ingredient, current_allergene))
            ingredients_matched.add(ingredient)
            allergenes_matched.add(current_allergene)
        else:
            all_allergenes.append(current_allergene)

    ingredients_with_no_allergenes = get_all_ingredients(food_list) - ingredients_matched
    counter = 0
    for ingredient in ingredients_with_no_allergenes:
        for food in food_list:
            if ingredient in food.ingredients:
                counter += 1
    return counter
