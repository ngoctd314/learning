# Tech1001 Kafka 

**Features**

- Reliability
- Durability
- Scalability
- High performance

**Application**

Messaging system, application logs, streaming, activity tracking, metrics, microservice pub/sub, bigdata platform: spark, hahoop.

**Kafka message**

+ Key-binary (can be null)
+ Value-binary (can be null)
+ Compress Type (none, gzip, snappy...)
+ Headers
+ Partition + Offset
+ Timestamp (system or user set)

**Topic**

Kafka có topic __consumer_offsets lưu thông tin về offset đã được read trong từng consumer => consumer bị lỗi, có thể đọc lại data từ offset được commit trước đó.

**Broker and cluster**

Kafka client kết nối đến 1 broker (bootstrap broker) => có đủ thông tin để kết nối đến toàn cluster (gồm brokers, topics, partitions, thông qua metadata khi kết nối)

**Topic replication**

**Confluent Schema registry**

- Mô tả các trường dữ liệu trong data record
- Trường dữ liệu bắt buộc/option
- Giá trị mặc định của các trường dữ liệu
