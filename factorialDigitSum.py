#problem-20 factorial digit sum

def factorial(n):
    return (n * factorial(n-1)) if n != 1 else n

pow = str(factorial(100))
sum = 0
for i in range(len(pow)):
    sum += int(pow[i])
print(sum)