-- https://www.hackerrank.com/challenges/the-report/problem

-- You are given two tables: Students and Grades. Students contains three columns ID (integer), Name (string) and Marks (integer). Grades contains Grade (integer, from 0 to 10), Min_Mark and Max_Mark.
-- Ketty gives Eve a task to generate a report containing three columns: Name, Grade and Mark. Ketty doesn't want the NAMES of those students who received a grade lower than 8. The report must be in descending order by grade -- i.e. higher grades are entered first. If there is more than one student with the same grade (8-10) assigned to them, order those particular students by their name alphabetically. Finally, if the grade is lower than 8, use "NULL" as their name and list them by their grades in descending order. If there is more than one student with the same grade (1-7) assigned to them, order those particular students by their marks in ascending order.
-- Write a query to help Eve. 

-- Sample output:
-- Maria 10 99
-- Jane 9 81
-- Julia 9 88 
-- Scarlet 8 78
-- NULL 7 63
-- NULL 7 68


SELECT
    CASE 
        WHEN Students.Marks < 70 THEN NULL
        ELSE Students.Name
    END AS Name,
    Grade, Students.Marks
FROM Students JOIN Grades
    ON Students.Marks BETWEEN Grades.Min_Mark AND Grades.Max_Mark

ORDER by Grades.Grade DESC, Students.Name ASC

