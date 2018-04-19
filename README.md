npuzzle
=======

Overview
-------
- The goal of this project is to programmatically solve the [N-puzzle](https://en.wikipedia.org/wiki/15_puzzle) using the A* algorithm with admissible heuristics
- Heuristics used: Manhattan distance, Manhattan distance + Linear conflict, Misplaced tiles
- Bonus: visualizer => server + front (reactJS) 

Usage
-------
```
client/ npm install && npm start
python assets/generator.py -s 3 > map
server/ make
./npuzzle -f map -greedy -visualizer
```
./npuzzle:<br/>
+ -f string<br/>
    	File containing the puzzle to solve<br/>
+ -greedy<br/>
    	Use greedy search (if not true, uniform-cost search is used by default)<br/>
+ -heuristic string<br/>
    	Heuristics: Manhattan "m", Misplaced Tiles "mt" or Manhattan + Linear Conflict "mlc" (default "mlc")<br/>
+ -visualizer<br/>
    	Launch a web app to visualize results on your browser<br/>



Most effective configurations
-----------------------------
- 3x3: Manhattan distance + Linear conflict in greedy search
- 4x4 and 5x5: Manhattan distance + Linear conflict in uniform-cost search
