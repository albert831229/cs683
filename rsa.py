import random

def get_primes(start, stop):
    if start >= stop:
        return []
    primes = [2]
    for n in range(3, stop + 1, 2):
        for p in primes:
            if n % p == 0:
                break
        else:
            primes.append(n)

    while primes and primes[0] < start:
        del primes[0]

    return primes


def no_gcd(a, b):
    for n in range(2, min(a, b) + 1):
        if a % n == b % n == 0:
            return False
    return True


def make_key_pair(length):
    # The key is stronger if ``p`` and ``q`` have similar bit length. We
    # choose two prime numbers in ``range(start, stop)`` so that the
    # difference of bit lengths is at most 2.
    start = 1 << (length // 2 - 1)
    stop = 1 << (length // 2 + 1)
    primes = get_primes(start, stop)

    p = random.choice(primes)
    primes.remove(p)
    q = random.choice(primes)

    fn = (p - 1) * (q - 1)
    for e in range(3, fn, 2):
        if no_gcd(e, fn):
            break

    for d in range(3, fn, 2):
        if d * e % fn == 1:
            break

    return PublicKey(p * q, e), PrivateKey(p * q, d)


class PublicKey:

    def __init__(self, n, e):
        self.e = e
        self.n = n

    def encrypt(self, x):
        return pow(x, self.e, self.n)


class PrivateKey:
    def __init__(self, n, d):
        self.d = d
        self.n = n

    def decrypt(self, x):
        return pow(x, self.d, self.n)


if __name__ == '__main__':
    public = PublicKey(n=2534665157, e=7)
    private = PrivateKey(n=2534665157, d=1810402843)

    assert public.encrypt(123) == 2463995467
    assert public.encrypt(456) == 2022084991
    assert public.encrypt(123456) == 1299565302

    assert private.decrypt(2463995467) == 123
    assert private.decrypt(2022084991) == 456
    assert private.decrypt(1299565302) == 123456

    for length in range(10, 17):
        public, private = make_key_pair(length)

        assert public.n == private.n

        x = random.randrange(public.n - 2)
        y = public.encrypt(x)
        assert private.decrypt(y) == x
