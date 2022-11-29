# Advent of Code 2020
Solutions to Advent of Code 2020

![Christmas image of a cup of hot chocolate](./image.png)

I'm using this project to play with Python, [Testing](##Testing) and [Github Actions](##Github-Actions)

* [Testing](#Testing)
  - [Continuous testing](#Continuous-testing)
* [Github Actions](#Github-Actions) 
  - [flake8](#flake8)
  - [black](black)
* [TIL](#TIL)
  - [Resources](#Resources)
  - [Python learnings](#Python-learnings)

## Testing

[pytest](https://docs.pytest.org/en/stable/) is a testing library for Python. It is widely used and has some advantages compared to the built-in `unitest` (here's a [comparison between pytest and unittest](https://github.com/renzon/pytest-vs-unittest)).

It has a powerfull test discovery, detecting test by default in functions beginning with `test` and in Python files beginning with `test_` or ending with `_test`.

`pytest` recommends putting tests in an extra directory outside the application code. 

```
src/
  __init__.py
  day_01/
      part_1.py
      part_2.py
  day_02/
      part_1.py
  tests/
      __init__.py
      day_01/
          __init__.py
          test_part_1.py
          test_part_2.py
      day_02/
          __init__.py
          test_part_1.py
```

When executing `pytest` from root folder, it will detect the following files starting with `test_`
- src/tests/day_01/test_part_1.py
- src/tests/day_01/test_part_2.py
- src/tests/day_02/test_part_1.py

Thanks to `__init__py` in both `day_01` and `day_02` folders we can use test files with the same name, as `__init__py` declares a [directory as a Python package](https://docs.python.org/3/reference/import.html#regular-packages).

To be able to import local code there is an `__init__.py` also in `tests` and in `src`.

Now `test_part_1.py` can load code from all folders under `src`, using `.` instead of `/`

```py
from src.day_01.part_1 import solution
```

### Continuous testing

To gain efficiency it is possible to run `pytest` on code change and only on impacted tests. `pytest` has two plugins that combined together provide this feature, [pytest-testmon](https://pypi.org/project/pytest-testmon/) and [pytest-watch](https://github.com/joeyespo/pytest-watch). 

`pytest-watch` is a CLI tool that runs `pytest`, and re-run it when a file changes. 

`pytest-testmon` uses [coverage.py](https://coverage.readthedocs.io/en/coverage-5.3.1/) to determine which tests are impacted by changes in code.

Install these plugin using `pip`

```bash
pip install pytest-watch pytest-testmon
```

And run `pytest-watcher` using `pytest --testmon`

```bash
ptw --runner "pytest --testmon"
```

## Github Actions

Github built-in CI/CD is free for public repositories since Aug, 2019. It has many workflow templates, including one for Python applications. To add it and start running linting and tests on Github, click on Actions -> New Workflow -> Python Applications. This will create a new configuration `yaml` under `.github/workflows`, that by defaults execute the actions at every push on `main` branch 

### flake8

[Flake8](https://flake8.pycqa.org/en/latest/) is a widely used Python linter. It supports storing its [configuration](./.flake8) in the root directory in a `.flake8` file, within a `[flake8]` section.

### black

[Black](https://black.readthedocs.io/en/stable/) is a widely used Python formatter. It supports storing its [configuration](./pyproject.toml) in a `TOML` file within a `[tool.black]` section. It can be used to format code on save, or to format all files from the command line:

```bash
black .
```

## TIL

Here's a collection of resources and learnings from 2020 edition

### Resources

* Parsing mathematical expressions with [Dijkstra's](https://en.m.wikipedia.org/wiki/Edsger_W._Dijkstra) [Shunting-yard algorithm](https://en.m.wikipedia.org/wiki/Shunting-yard_algorithm) - from day 18
* [Hexagonal grids](https://www.redblobgames.com/grids/hexagons/) - from day 24

### Python learnings

#### `re.match` vs `re.search` vs `re.fullmatch`

`re.match` looks for the pattern at the beginning of the string, or at a specific position `pos`

```python
import re

a_string = 'something whatever'
pattern_at_the_beginning = 'some'
pattern_in_the_middle = 'g what'
print re.match(pattern_at_the_beginning, a_string) # matches
print re.match(pattern_in_the_middle, a_string) # doesn't match
# using a specific position where to look for the pattern
print re.match(pattern_in_the_middle, a_string, pos=8) # matches
```

`re.search` scans through the string looking for a position where the pattern is matched.

```python
import re

a_string = 'something whatever'
pattern_at_the_beginning = 'some'
pattern_in_the_middle = 'g what'
print re.search(pattern_at_the_beginning, a_string) # matches
print re.search(pattern_in_the_middle, a_string) # matches
```

`re.fullmatch` returns a match object iff the whole string matches the regular expression pattern

```python
import re

a_string = 'something whatever'
pattern_at_the_beginning = 'some'
pattern_in_the_middle = 'g what'
print re.fullmatch(pattern_at_the_beginning, a_string) # doesn't match
print re.fullmatch(pattern_in_the_middle, a_string) # doesn't match
pattern_full = '\w+ \w+' # 1 or more word characters + space + 1 or more word characters
print re.fullmatch(pattern_full, a_string) # matches
```

More on `match` vs `search` in [Python's official documentation](https://docs.python.org/2/library/re.html#search-vs-match)