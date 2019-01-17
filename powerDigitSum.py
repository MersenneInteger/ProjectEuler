#problem-16 power digit sum

pow = str(2 ** 1000)
sum = 0
for i in range(len(pow)):
    sum += int(pow[i])
print(sum)
