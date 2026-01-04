SELECT HACKERS.HACKER_ID, HACKERS.NAME
FROM HACKERS JOIN SUBMISSIONS ON HACKERS.HACKER_ID=SUBMISSIONS.HACKER_ID
JOIN CHALLENGES ON SUBMISSIONS.CHALLENGE_ID=CHALLENGES.CHALLENGE_ID
JOIN DIFFICULTY ON CHALLENGES.DIFFICULTY_LEVEL=DIFFICULTY.DIFFICULTY_LEVEL
WHERE SUBMISSIONS.SCORE=DIFFICULTY.SCORE
GROUP BY HACKERS.NAME, HACKERS.HACKER_ID
HAVING COUNT(HACKERS.HACKER_ID)>1
ORDER BY HACKERS.HACKER_ID ASC, COUNT(CHALLENGES.CHALLENGE_ID) DESC



select
    h.hacker_id, h.name
from 
    hackers h
join
    submissions s on h.hacker_id = s.hacker_id
join
    challenges c on s.challenge_id = c.challenge_id
join
    difficulty d on c.difficulty_level = d.difficulty_level
where 
    s.score = d.score
group by
    h.hacker_id, h.name
having
    count(*) > 1
order by
    count (*) desc, hacker_id asc;