CREATE TABLE `posts` (
  `id` varchar(50) NOT NULL DEFAULT uuid(),
  `title` text DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `thumbnail` text DEFAULT NULL,
  `author_id` varchar(50) DEFAULT NULL,
  `reviewer_id` varchar(50) DEFAULT NULL,
  `publish_at` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp(),
  `status` enum('draft','published','awaiting_review','approved') DEFAULT 'draft',
  `is_deleted` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
