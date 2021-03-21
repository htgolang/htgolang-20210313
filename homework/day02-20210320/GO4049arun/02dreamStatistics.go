package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	var database = make(map[string]int)
	// read content from a file
	txt, err := ioutil.ReadFile(`02IhaveADream.txt`)
	if err != nil {
		panic(err)
	}
	//(1)Main function realization
	for i := 0; i < len(txt); i++ {
		//filter a character in A~Z||a~z
		if (byte(txt[i]) >= 65 && byte(txt[i]) <= 90) || (byte(txt[i]) >= 97 && byte(txt[i]) <= 122) {
			//translate the type of a character to that of the string
			var index string = string(byte(txt[i]))
			//initialize the index and value for the database of map
			value, ok := database[index]
			if ok {
				//add one if it exist
				database[index] = value + 1
			} else {
				//initialize to one if it's not existed
				database[index] = 1
			}
		}
	}
	/* for k, v := range database {
		fmt.Println(k, v)
	} */
	//(2)sorted the database of map by key
	//(2-1)To store the keys in slice in sorted order
	i := 0
	keys := make([]string, len(database))
	for k := range database {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	//(2-2)read the values from the database map by key
	for _, k := range keys {
		fmt.Println(k, database[k])
	}
} /*
> go run .\02dreamStatistics.go
A 18
B 4
C 6
D 1
E 1
F 3
G 9
H 1
I 23
J 1
L 12
M 7
N 23
O 3
P 3
R 1
S 8
T 13
W 16
Y 3
a 518
b 103
c 171
d 247
e 860
f 214
g 159
h 375
i 524
j 20
k 51
l 317
m 177
n 438
o 585
p 91
q 6
r 399
s 400
t 637
u 174
v 80
w 144
x 5
y 118
z 5
*/
