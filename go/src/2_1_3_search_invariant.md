### Initialization
Before entering the loop. i := -1. A[-1..i] does not contain the value we want to search for. A[-1..i] is not part of the array.

### Maintenance
At each step A[-1..i] where A[i] != value, this set will not contain the value.

### Termination
When A[-1..i] contains the value, we terminate the loop.
