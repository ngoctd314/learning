package main

/*
Solution hiện tại là dùng poller để check QR code status.
+ Lợi ích: dễ dàng implement
+ Nhược điểm: call nhiều

WebSocket cho phép two-way communication giữa client và web server.
+ Nhược điểm: khó implement hơn poller, phải maintain connection logic

Chú ý: đóng ws connection khi QR code status success, client close tab.
*/
