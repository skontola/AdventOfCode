# Solution for day one of advent of code. Pretty straight forward.
# Part 1
intList = []
total = 0

f = open("input1.txt", "r")

inputList = f.read().splitlines()

for item in inputList:
    tmpList = []
    for x in item:
        if x.isdigit():
            tmpList.append(int(x))
    intList.append(tmpList)

for item in intList:
    total = total + 10 * item[0] + item[-1]

print("Part 1 total:", total)

# Part 2
total = 0
tmpList = []

def notMinusOne(x):
    if (x == -1):
        return False
    else:
        return True

def finder(lst, str, i):
    x = lst.find(str, i)
    if (x != -1):
        #print(x)
        tmpList.append(x)
        #print(tmpList)
        finder(lst, str, x+1)

# Fan vad ful men orkar inte fixa den
for item in inputList:
    rowList = [-1 for i in range(len(item))]

    finder(item, "one", 0)
    finder(item, "1", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 1
    tmpList.clear()
    finder(item, "two", 0)
    finder(item, "2", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 2
    tmpList.clear()
    finder(item, "three", 0)
    finder(item, "3", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 3
    tmpList.clear()
    finder(item, "four", 0)
    finder(item, "4", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 4
    tmpList.clear()
    finder(item, "five", 0)
    finder(item, "5", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 5
    tmpList.clear()
    finder(item, "six", 0)
    finder(item, "6", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 6
    tmpList.clear()
    finder(item, "seven", 0)
    finder(item, "7", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 7
    tmpList.clear()
    finder(item, "eight", 0)
    finder(item, "8", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 8
    tmpList.clear()
    finder(item, "nine", 0)
    finder(item, "9", 0)
    for i in range(len(tmpList)):
        rowList[tmpList[i]] = 9
    tmpList.clear()
    
    rowList = filter(notMinusOne, rowList)
    rowList = list(rowList)

    total = total + 10 * rowList[0] + rowList[-1]

print("Part 2 total:", total)