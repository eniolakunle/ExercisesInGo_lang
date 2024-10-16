## Exercise 1
Create a struct named Person with three fields: FirstName and LastName of
type string and Age of type int. Write a function called MakePerson that
takes in firstName, lastName, and age and returns a Person. Write a second
function MakePersonPointer that takes in firstName, lastName, and age and
returns a *Person. Call both from main. Compile your program with go build
-gcflags="-m". This both compiles your code and prints out which values
escape to the heap. Are you surprised about what escapes?
## Exercise 2
Write two functions. The UpdateSlice function takes in a []string and a
string. It sets the last position in the passed-in slice to the passed-in string. At
the end of UpdateSlice, print the slice after making the change. The GrowSlice
function also takes in a []string and a string. It appends the string onto the
slice. At the end of GrowSlice, print the slice after making the change. Call these
functions from main. Print out the slice before each function is called and after
each function is called. Do you understand why some changes are visible in main
and why some changes are not?
## Exercise 3
Write a program that builds a []Person with 10,000,000 entries (they could all
be the same names and ages). See how long it takes to run. Change the value
of GOGC and see how that affects the time it takes for the program to complete.
Set the environment variable GODEBUG=gctrace=1 to see when garbage collections
happen and see how changing GOGC changes the number of garbage collections.
What happens if you create the slice with a capacity of 10,000,000?

### Exercise 3 Results:
#### Populating splice without a capacity 🐌
```
Time Started:  2024-10-16 01:53:51.293392 -0400 EDT m=+0.000150215
gc 1 @0.002s 8%: 0.020+1.3+0.004 ms clock, 0.082+0/1.2/0.13+0.019 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 4 P
gc 2 @0.004s 5%: 0.008+6.7+0.003 ms clock, 0.032+0/1.0/5.7+0.014 ms cpu, 5->5->1 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 4 P
                    .
                    .
        19 Garbage Collections later
                    .
                    .

gc 21 @0.837s 11%: 0.032+272+0.003 ms clock, 0.12+22/84/345+0.012 ms cpu, 466->466->385 MB, 466 MB goal, 0 MB stacks, 0 MB globals, 4 P
gc 22 @1.227s 11%: 0.062+357+0.003 ms clock, 0.25+67/114/471+0.014 ms cpu, 829->829->602 MB, 829 MB goal, 0 MB stacks, 0 MB globals, 4 P
Duration:  2.086162013s
```

#### Populating splice w/ a 10 million entry capacity 🐎

```Time Started:  2024-10-16 01:54:05.447712 -0400 EDT m=+0.000139864
gc 1 @0.001s 22%: 0.022+244+0.026 ms clock, 0.088+0/220/441+0.10 ms cpu, 381->381->381 MB, 382 MB goal, 0 MB stacks, 0 MB globals, 4 P
Duration:  300.642322ms
```

