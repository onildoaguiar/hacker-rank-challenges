# Left Rotation

A left rotation operation on an array of size `n` shifts each of the array's elements `1` unit to the left. Given an integer, `d`, rotate the array that many steps left and return the result.

### Example
`d` = 2

`arr` = [1, 3, 4, 4, 5]

After `2` rotations, `arr` = [3, 4, 5, 1, 2].

### Function Description

Complete the rotateLeft function in the editor below.

rotateLeft has the following parameters:

- int d: the amount to rotate by
- int arr[n]: the array to rotate

### Returns

- int[n]: the rotated array

### Input Format

The first line contains two space-separated integers that denote , the number of integers, and , the number of left rotations to perform.
The second line contains  space-separated integers that describe .

### Sample Input

5 4
1 2 3 4 5

### Sample Output

5 1 2 3 4