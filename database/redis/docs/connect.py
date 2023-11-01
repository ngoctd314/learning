import redis

r = redis.Redis(host='192.168.49.2', port=30301, decode_responses=True)

set1 = r.sadd("bikes:racing:france", "bike:1")
print(set1)
set1 = r.sadd("bikes:racing:france", "bike:1")
print(set1)
set1 = r.sadd("bikes:racing:france", "bike:2")
print(set1)
set1 = r.sismember("bikes:racing:france", "bike:1")
print(set1)
set1 = r.sismember("bikes:racing:france", "bike:2")
print(set1)
set1 = r.sismember("bikes:racing:france", "bike:3")
print(set1)
