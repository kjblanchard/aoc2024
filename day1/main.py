list1 = []
list2 = []
with open('day1/input.txt', 'r') as distances:
    for line in distances.readlines():
        split = line.split()
        list1.append(int(split[0]))
        list2.append(int(split[1]))
list1.sort()
list2.sort()
# part1
distanceTotal = 0
for i in range(len(list1)):
    distanceTotal += abs(list1[i] - list2[i])
print(distanceTotal)
# part2
similarity = 0
for lhs in list1:
    count = 0
    for rhs in list2:
        if lhs == rhs:
            count += 1
        elif rhs > lhs:
            similarity += (lhs * count)
            break
print(similarity)
