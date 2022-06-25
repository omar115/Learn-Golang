* serial linking
so, we are using http function to check status,
of each links. How?
we take the slice and iteration through it. After
then we get status. We can call it Sequencial searching.

* What is the problem?
- the problem is it taking time after one request proceed.
So let's think like this, we have 7 links and each link takes
some amount of seconds time. If we have 1 million of list, then 
getting one link status may be take too long to expect.

* Solution
- solution is pretty straightforward. We need to introduce threading
because using it we can take list of user data and among data will 
