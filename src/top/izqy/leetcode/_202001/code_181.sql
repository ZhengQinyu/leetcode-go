# Write your MySQL query statement below

select e1.Name AS Employee from Employee e1 join Employee e2 ON e2.Id = e1.ManagerId where e1.ManagerId is not null and e1.Salary > e2.Salary