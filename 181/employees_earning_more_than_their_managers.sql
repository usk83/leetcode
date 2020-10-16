-- 181. Employees Earning More Than Their Managers
-- https://leetcode.com/problems/employees-earning-more-than-their-managers/

-- # SQL Schema
-- Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, ManagerId int);
-- Truncate table Employee;
-- insert into Employee (Id, Name, Salary, ManagerId) values ('1', 'Joe', '70000', '3');
-- insert into Employee (Id, Name, Salary, ManagerId) values ('2', 'Henry', '80000', '4');
-- insert into Employee (Id, Name, Salary, ManagerId) values ('3', 'Sam', '60000', 'None');
-- insert into Employee (Id, Name, Salary, ManagerId) values ('4', 'Max', '90000', 'None');

SELECT
  Name AS Employee
FROM
  Employee AS e
JOIN
  ( SELECT Id, Salary FROM Employee ) AS m
  ON
    e.ManagerId = m.Id AND e.Salary > m.Salary
;
