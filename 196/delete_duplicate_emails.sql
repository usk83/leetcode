-- 196. Delete Duplicate Emails
-- https://leetcode.com/problems/delete-duplicate-emails/

INSERT INTO Person (Id, Email)
VALUES (3, "john@example.com");

SELECT p1.Id FROM Person p1, Person p2 WHERE p1.Email = p2.Email AND p1.Id > p2.Id;

-- INSERT INTO Person (Id, Email)
-- VALUES (1, "john@example.com"),
-- (2, "bob@example.com"),
-- (3, "john@example.com");

select Id from Person group by Email;

select * from Person where Id not in (
  SELECT min(Id)
  FROM Person
  group by Email
);

delete from `Person` where Id not in (
  SELECT id FROM (
    SELECT min(Id) as id
    FROM Person
    group by Email
  ) as ids
);

Every derived table must have its own alias

DELETE p1 FROM Person p1 WHERE p1.Id NOT IN (SELECT p2.Id FROM Person p2 WHERE p2.Id = 2)
