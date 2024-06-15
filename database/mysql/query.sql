SELECT
  `users`.`id`,
  `users`.`username`,
  `Order`.`id` AS `Order__id`,
  `Order`.`user_id` AS `Order__user_id`,
  `Order`.`price` AS `Order__price`
FROM
  `users`
  LEFT JOIN `orders` `Order` ON `users`.`id` = `Order`.`user_id`
  AND price = 1;

SELECT
  `users`.`id`,
  `users`.`name`,
  `users`.`age`,
  `Company`.`id` AS `Company__id`,
  `Company`.`name` AS `Company__name`
FROM
  `users`
  LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`
  AND `Company`.`alive` = true;
