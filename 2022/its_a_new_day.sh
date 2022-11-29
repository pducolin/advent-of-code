#!/bin/bash
TODAY=$(date +"%d");
XMAS=25;
if [[ $TODAY > $XMAS ]]; then
    echo "Xmas is over, so is AoC"
    exit 1
else
    echo "🎄 It's AoC day $TODAY, get ready to code 👩‍💻..."
fi
#create code folder and placeholder files
if [ -d day_"$TODAY" ]; then
  echo "🎄 Day $TODAY already created 😲"
  exit 1
fi
cp -R template day_"$TODAY"