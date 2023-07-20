CREATE TABLE
    `event_logs` (
        `id` varchar(36) NOT NULL,
        `event_name` varchar(60) NOT NULL,
        `event_time` datetime NOT NULL,
        `description` text,
        `subject` varchar(60) DEFAULT NULL,
        `object` varchar(60) DEFAULT NULL,
        `data` json DEFAULT NULL,
        `created_at` datetime DEFAULT NULL,
        `client_id` varchar(60) DEFAULT NULL,
        PRIMARY KEY (`id`, `event_time`) USING BTREE
    ) ENGINE = InnoDB DEFAULT CHARSET = latin1;

CREATE TABLE
    `event_logs` (
        `id` VARCHAR(36) NOT NULL,
        `event_name` VARCHAR(60) NOT NULL,
        `event_time` DATETIME NOT NULL,
        `description` TEXT,
        `subject` VARCHAR(60) DEFAULT NULL,
        `object` VARCHAR(60) DEFAULT NULL,
        `data` JSON DEFAULT NULL,
        `created_at` DATETIME DEFAULT NULL,
        `client_id` VARCHAR(60) DEFAULT NULL,
        PRIMARY KEY (`id`, `event_time`) USING BTREE
    ) ENGINE = InnoDB DEFAULT CHARSET = latin1;

CREATE TABLE
    `event_hooks` (
        `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
        `client_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'client_id ref to client_id of iam_clients',
        `name` varchar(180) COLLATE utf8mb4_unicode_ci NOT NULL,
        `verified` tinyint(4) NOT NULL COMMENT 'webhook is verified or not',
        `status` int(11) NOT NULL COMMENT 'webhook is active/inactive for receive incoming webhook',
        `hook_type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'HTTP support only',
        `config` json NOT NULL COMMENT 'config for webhook request constructor',
        `events` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'list of subscribe events that client_id want to monitor',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci
CREATE TABLE
    `event_hooks` (
        `id` VARCHAR(36) COLLATE utf8mb4_unicode_ci NOT NULL,
        `client_id` VARCHAR(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'client_id ref to client_id of iam_clients' `name` VARCHAR(180) COLLATE utf8mb4_unicode_ci NOT NULL,
        `name` VARCHAR(180) COLLATE utf8mb4_unicode_ci NOT NULL,
        `verified` TINYINT(4) NOT NULL COMMENT 'webhook is verified or not',
        `status` INT(11) NOT NULL COMMENT 'webhook is active/inactive for receive incoming webhook'
    )