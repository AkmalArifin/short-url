CREATE TABLE `shorten_url` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `url` varchar(256) UNIQUE NOT NULL,
  `short_code` varchar(256),
  `access_count` int DEFAULT 0,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE INDEX `shorten_url_index_0` ON `shorten_url` (`id`, `url`);
