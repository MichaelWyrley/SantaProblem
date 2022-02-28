# SantaProblem
My Solution to the common santa problem in GO

# Problem
This problem demonstrates the use of semaphores to coordinate three types of threads: Santa Claus sleeps in his shop at the North Pole and can only be wakened by either

a. all nine reindeer being back from their vacation in the South Pacific, or

b. by some of the elves having difficulties making toys.

To allow Santa to get some sleep, the elves can only wake him when three of them have problems. While three elves are having their problems solved, any other elves wishing to visit Santa must wait for those elves to return. If Santa wakes up to find three elves waiting at his shops door, along with the last reindeer having come back from the tropics, Santa has decided that the elves can wait until after Christmas, because it is more important to get his sleigh ready. (It is assumed that the reindeer dont want to leave the tropics, and therefore they stay there until the last possible moment.) The last reindeer to arrive must get Santa while the others wait in a warming hut before being harnessed to the sleigh.