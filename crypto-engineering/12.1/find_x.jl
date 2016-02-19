# Let p = 89, q = 107, n = pq, a = 3, and b = 5. Find x in Zn such that
# a = x (mod p) and b = x (mod q).

p=89
q=107
n=p*q
a=3
b=5

for x in 1:n
    if a == x % p && b == x % q
        println("x=", x)
        exit(0)
    end
end

println("No x found.")
