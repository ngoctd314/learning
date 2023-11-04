import redis

r = redis.Redis(host='192.168.49.2', port=30301, decode_responses=True)


r.hincrby("bike:1", "price", 100)
res4 = r.hget("bike:1", "price")
print(res4)
