import redis

r = redis.Redis(host='192.168.49.2', port=30301, decode_responses=True)

res37 = r.set("new_bikes", "bike:1")
print(res37)


res38 = r.lpop("list39")
print(res38)

