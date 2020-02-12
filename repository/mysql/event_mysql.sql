USE demo;

DROP TABLE IF EXISTS `event`;

CREATE TABLE `event` (
  `id` varchar(30) NOT NULL,
  `title` varchar(100) NOT NULL,
  `description` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO event (`id`, `title`, `description`) VALUES ("id_go", "Go", "https://golang.org/");

INSERT INTO event (`id`, `title`, `description`) VALUES ("id_gin", "Gin", "https://github.com/gin-gonic/gin");
