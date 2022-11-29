#!/bin/bash
TODAY=$(date +"%d");
XMAS=25;
if [[ "$TODAY" > "$XMAS" ]]; then
    echo "Xmas is over, so is AoC"
    exit 1
else
    echo "🎄 It's AoC day $TODAY, get ready to code 👩‍💻..."
fi
#create code folder and placeholder files
if [ -d src/day_"$TODAY" ]; then
  echo "🎄 Day $TODAY already created 😲"
  exit 1
fi
mkdir -p src/day_"$TODAY";
touch src/day_"$TODAY"/input.txt;
touch src/day_"$TODAY"/README.md;
cp template/part_x.py src/day_"$TODAY"/part_1.py;
cp template/part_x.py src/day_"$TODAY"/part_2.py;
# create test folder and placeholder files
mkdir -p src/tests/day_"$TODAY";
touch src/tests/day_"$TODAY"/__init__.py;
sed "s/day_xx.part_x/day_$TODAY.part_1/" template/_test_part_x.py >> src/tests/day_"$TODAY"/test_part_1.py;
sed "s/day_xx.part_x/day_$TODAY.part_2/" template/_test_part_x.py >> src/tests/day_"$TODAY"/test_part_2.py;