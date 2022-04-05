SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";
CREATE TABLE `dailypass` (
  `dailyPassId` int(11) NOT NULL,
  `userId` int(11) DEFAULT NULL,
  `unlockedChapters` longtext
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `dailypass` (`dailyPassId`, `userId`, `unlockedChapters`) VALUES
(11, 1, '{\"\":\"\",\"1\":\"4\",\"3\":\"4\",\"4\":\"4\",\"87\":\"3\",\"88\":\"3\",\"89\":\"3\",\"90\":\"4\"}'),
(12, 4, '{\"1\":4,\"3\":\"4\",\"4\":\"4\"}'),
(13, 5, '{\"1\":4,\"3\":\"4\",\"4\":\"4\"}'),
(14, 7, '{\"1\":4,\"2\":4,\"3\":\"4\",\"4\":\"4\"}'),
(15, 8, '{\"1\":4,\"2\":4,\"3\":\"4\",\"4\":\"4\"}'),
(16, 9, '{\"1\":4,\"2\":4,\"3\":\"4\",\"4\":\"4\"}'),
(17, 10, '{\"1\":4,\"2\":4,\"3\":4,\"4\":\"6\"}');
ALTER TABLE `dailypass`
  ADD PRIMARY KEY (`dailyPassId`),
  ADD UNIQUE KEY `iddailyPass_UNIQUE` (`dailyPassId`),
  ADD UNIQUE KEY `userId_UNIQUE` (`userId`);
ALTER TABLE `dailypass`
  MODIFY `dailyPassId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;COMMIT;