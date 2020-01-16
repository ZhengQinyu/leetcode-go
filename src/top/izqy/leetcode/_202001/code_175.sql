# Write your MySQL query statement below
select
    FirstName, LastName, City, State
from
    Person
LEFT JOIN Address ON Address.PersonId = Person.PersonId