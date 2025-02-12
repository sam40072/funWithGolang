# fun with golang
## info
this is just a project I've made while trying to learn Golang
## features
# command line arguments 
takes arguments from the console
# fibonacci number
generates up to the nth number of the fibonacci sequence and prints the slice and final number
# Collatz conjecture
shows steps taken to get to one
# Collatz conjecture with range
takes range [x, y] and prints the number with the largest amount of steps needed to get to one, the number of steps and average amount of steps of all the numbers.
# coin flips
flips a coin n number of times and prints the largest consecutive tails and heads in the sequence. Also optionally prints the sequence   

## usage:
<executable> <module> <aditional options>
<executable> f <number of steps> (gets fibonacci up to <number of steps> steps)
<executable> c <number to test> (performs all collatz conjecture on the <number to test> until 1 is reached)
<executable> cs <lower bound> <upper bound> (performs the collatz conjecture on all the numbers in the range)
<executable> cf <num of coins> (flips the coin <num of coins> times)
