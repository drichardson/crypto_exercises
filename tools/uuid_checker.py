#!/usr/bin/python3

# Python version of uuid_checker.go

import uuid
import time

start_time = time.time()

collisions = {}
collisionCount = 0
summaryEvery = 1000000
summaryCounter = 0
errors = 0
count = 0

while True:
    u = uuid.uuid4()

    if u in collisions:
        uCount = collisions[u] + 1
        collisions[u] = uCount
        print(f"{uCount} collisions for UUID {u}")
    else:
        collisions[u] = 0

    count += 1
    summaryCounter += 1
    if summaryCounter >= summaryEvery:
        summaryCounter = 0
        runningTime = time.time() - start_time
        print(f"Summary: count={count}, errors={errors}, collisions={collisionCount}, runningTime={runningTime}s")



