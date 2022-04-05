SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

CREATE TABLE `series` (
  `seriesId` int(11) NOT NULL,
  `author` varchar(45) DEFAULT NULL,
  `chapters` longtext,
  `name` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `series` (`seriesId`, `author`, `chapters`, `name`) VALUES
(1, '14', '[\"1\",\"2\",\"3\",\"4\",\"5\",\"16\"]', '12'),
(2, '14', '[\"1\",\"2\",\"3\",\"4\",\"5\"]', '12'),
(3, '14', '[\"1\",\"2\",\"3\",\"4\",\"5\"]', '12'),
(4, '1jhg4', '[\"1\",\"2\",\"3\",\"4\",\"5\"]', '12'),
(5, '1jhg4', '[\"1\",\"2\",\"3\",\"4\",\"5\"]', '12');

ALTER TABLE `series`
  ADD PRIMARY KEY (`seriesId`),
  ADD UNIQUE KEY `idseries_UNIQUE` (`seriesId`);

ALTER TABLE `series`
  MODIFY `seriesId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;COMMIT;
