# PvZ I-Zombie Solver
I love the Plant vs Zombies game and one of my favourite puzzles is the I-Zombie infinite level. I typically get stuck at level 16 so I have decided to build a script to help me determine minimum zombies to use to clear any given level.

The game consists of a grid of 9 columns and 5 rows. the first five columns are filled with plants of different characteristics and zombies can only be placed on cells after the 5th row (except for bungee zombies).

Every zombie has a sun cost associated with using it. The goal is to figure out the minimum number of sun to use to eat all brains.

## Board Representation
Given we have 5 by 6 matrix, we denote rows as letters A to E and columns as numbers 1 to 6.

|  | 0 | 1 | 2 | 3 | 4 | 5 | 6 |
|--|--|--|--|--|--|--|--|
| A | 0:0 (A0) | 0:1 (A1) | 0:2 (A2) | 0:3 (A3) | 0:4 (A4) | 0:5 (A5)| 0:6 (A6) |
| B | 1:0 (B0) | 1:1 (B1) | 1:2 (B2) | 1:3 (B3) | 1:4 (B4) | 1:5 (B5)| 1:6 (B6) |
| C | 2:0 (C0) | 2:1 (C1) | 2:2 (C2) | 2:3 (C3) | 2:4 (C4) | 2:5 (C5)| 2:6 (C6) |
| D | 3:0 (D0) | 3:1 (D1) | 3:2 (D2) | 3:3 (D3) | 3:4 (D4) | 3:5 (D5)| 3:6 (D6) |
| E | 4:0 (E0) | 4:1 (E1) | 4:2 (E2) | 4:3 (E3) | 4:4 (E4) | 4:5 (E5)| 4:6 (E6) |

Furthermore, because objects do not just jump from cell to cell we further sub-divide the cells into a submatrix. The submatrix dimensions is variable but it will always be an odd number and a multple of 3. so (3x3, 9x9, 15x15, etc.).
The resolution specified will determine this submatrix size and it helps us fine tune the level of detail we want for the simulation.
If resolution is lower than fastest object in the simulation then we will experience some objects disappearing and appearing within frames which will be undesireable as it will make the simulation inaccurate and we will miss potential object interactions. For now, when testing only zombie movements a 3x3 resolution is sufficient. But for full game simulation 15x15 is required as the fastest object in the game (Pea) is 15 sub matrix cells per frame. So it should take 15 frames for a pea to move 1 cell. 15 is chosen because pea is 15 times faster than the slowest object.

The submatrix co-ordinates are denoted by x,y, and z axis. The x axis is the row, the y axis is the column and the z axis is the height.
### Placing And Moving Zombies
Zombies can be placed on any of the rows but can be placed only on the 6th column.
They will also be placed in the middle of the cell. For a 3x3 matrix that will be sub cell  1,1 for 15 by 15 it will be subcell 7,7.
The zombie will be expected to move from where it's placed to the end of the board on the 0th column in a straight line.

## Dimension Estimations
Based on [this reddit post](https://www.reddit.com/r/PlantsVSZombies/comments/a8u8gw/scaling_of_plants_and_zombies_in_pvz1/)
For 9 cells total length in cm is 10621100 micrometer so each cell is 1180122 micrometer. Subcell will be 1180122/RESOLUTION micrometer.
The tallest object (Squash Plant) the pole vault zombie can jump is 807900 micro meter. We will approximate this to 810000 micrometer.
It takes the fastest object (Pea) 250 milli seconds to clear a cell.
Frame per milli second is RESOLUTION / 250 milli seconds.
Accelleration due to gravity is 9.8 micrometer per square milisecond

Frames is equivalent to time and subcells is equivalent to distance
