--task "Surface Area and Volume of a Box"

-- # write your SQL statement here: 
-- you are given a table 'box' with columns: width (int), height (int), depth (int)
-- return a query with columns: width, height, depth, area (int),  as volume (int)
-- sort results by area ascending, then volume ascending, then width ascending, then height ascending

SELECT WIDTH, HEIGHT, DEPTH, 2 * (WIDTH * HEIGHT + DEPTH * WIDTH + DEPTH * HEIGHT) AS AREA, WIDTH * HEIGHT * DEPTH AS VOLUME
FROM BOX
ORDER BY AREA ASC, VOLUME DESC, WIDTH ASC, HEIGHT ASC;