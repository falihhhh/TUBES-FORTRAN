import random, time, platform, statistics, json, sys

def bubble(a):
    n = len(a)
    for i in range(n - 1):
        swapped = False
        for j in range(n - 1 - i):
            if a[j] > a[j + 1]:
                a[j], a[j + 1] = a[j + 1], a[j]
                swapped = True
        if not swapped:
            break

def selection(a):
    n = len(a)
    for i in range(n - 1):
        m = i
        for j in range(i + 1, n):
            if a[j] < a[m]:
                m = j
        if m != i:
            a[i], a[m] = a[m], a[i]

def insertion(a):
    for i in range(1, len(a)):
        key = a[i]
        j = i - 1
        while j >= 0 and a[j] > key:
            a[j + 1] = a[j]
            j -= 1
        a[j + 1] = key

def bubble_c(a):
    c = 0; n = len(a)
    for i in range(n - 1):
        sw = False
        for j in range(n - 1 - i):
            c += 1
            if a[j] > a[j + 1]:
                a[j], a[j + 1] = a[j + 1], a[j]; sw = True
        if not sw: break
    return c

def selection_c(a):
    c = 0; n = len(a)
    for i in range(n - 1):
        m = i
        for j in range(i + 1, n):
            c += 1
            if a[j] < a[m]: m = j
        if m != i: a[i], a[m] = a[m], a[i]
    return c

def insertion_c(a):
    c = 0
    for i in range(1, len(a)):
        key = a[i]; j = i - 1
        while j >= 0:
            c += 1
            if a[j] > key:
                a[j + 1] = a[j]; j -= 1
            else:
                break
        a[j + 1] = key
    return c

ALGS = {"Bubble Sort": bubble, "Selection Sort": selection, "Insertion Sort": insertion}
ALGS_C = {"Bubble Sort": bubble_c, "Selection Sort": selection_c, "Insertion Sort": insertion_c}
REPS = 5
random.seed(2026)

def timeit(fn, data):
    ts = []
    for _ in range(REPS):
        a = data[:]
        t0 = time.perf_counter(); fn(a); t1 = time.perf_counter()
        assert a == sorted(data)
        ts.append((t1 - t0) * 1000)
    return statistics.mean(ts), statistics.stdev(ts)

res = {"env": {"python": platform.python_version(), "platform": platform.platform(), "processor": platform.processor() or platform.machine()}}
res["random"] = {}
for n in [100, 500, 1000, 2500, 5000]:
    data = [random.randint(0, 100000) for _ in range(n)]
    res["random"][n] = {}
    for name, fn in ALGS.items():
        m, s = timeit(fn, data)
        res["random"][n][name] = [round(m, 3), round(s, 3)]
    print("random", n, res["random"][n], flush=True)

res["cases"] = {}
n = 2000
base = [random.randint(0, 100000) for _ in range(n)]
cases = {"Terurut (best case)": sorted(base), "Acak (average case)": base, "Terbalik (worst case)": sorted(base, reverse=True)}
for cname, data in cases.items():
    res["cases"][cname] = {}
    for name, fn in ALGS.items():
        m, s = timeit(fn, data)
        res["cases"][cname][name] = [round(m, 3), round(s, 3)]
    print(cname, res["cases"][cname], flush=True)

res["ops"] = {}
n = 1000
data = [random.randint(0, 100000) for _ in range(n)]
for name, fn in ALGS_C.items():
    res["ops"][name] = fn(data[:])
print("comparisons n=1000 random", res["ops"])
json.dump(res, open("results.json", "w"), indent=1)
print(res["env"])