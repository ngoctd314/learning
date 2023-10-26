import redis

r = redis.Redis(host='192.168.49.2', port=30301, decode_responses=True)

# res1 = r.lpush('list', 'key1')
# print(res1)

# list = r.rpush('list', 'key1')
# r.rpush('list', 'key2')
# r.rpush('list', 'key3')

res = r.rpush("bikes:repairs", "list1")
print(res)


# res32 = r.brpop("bikes:repairs")
# print(res32)
