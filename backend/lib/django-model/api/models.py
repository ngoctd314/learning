from django.db import models


class Person(models.Model):
    first_name = models.CharField(max_length=30)
    last_name = models.CharField(max_length=30)

#
# CREATE TABLE `api_person` (
# `id` bigint NOT NULL AUTO_INCREMENT,
# `first_name` varchar(30) NOT NULL,
# `last_name` varchar(30) NOT NULL,
# PRIMARY KEY (`id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
#
