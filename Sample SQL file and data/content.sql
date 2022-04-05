SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";
CREATE TABLE `content` (
  `chapterId` int(11) NOT NULL,
  `title` longtext,
  `story` longtext,
  `seriesId` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `content` (`chapterId`, `title`, `story`, `seriesId`) VALUES
(1, 'Sfdsf', 'sdfdfds', '1'),
(2, 'Sfdsf', 'sdfdfds', '2'),
(3, 'Sfdsf', 'sdfdfds', '1'),
(4, 'Sfdsf', 'sdfdfds', '1'),
(5, 'Sfdsf', 'sdfdfds', '1'),
(6, 'Sfdsf', 'sdfdfds', '1'),
(7, 'Sfdsf', 'sdfdfds', '1'),
(8, 'Sfdsf', 'sdfdfds', '1'),
(9, 'Sfdsf', 'sdfdfds', '1'),
(10, 'Sfdsf', 'sdfdfds', '1'),
(11, 'Sfdsf', 'sdfdfds', '1'),
(12, 'Sfdsf', 'sdfdfds', '1'),
(13, 'Sfdsf', 'sdfdfds', '1'),
(14, 'Sfdsf', 'sdfdfds', '1'),
(15, 'Sfdsf', 'sdfdfds', '1'),
(16, 'Sfdsf', 'sdfdfds', '1');

ALTER TABLE `content`
  ADD PRIMARY KEY (`chapterId`),
  ADD UNIQUE KEY `contentId_UNIQUE` (`chapterId`);
ALTER TABLE `content`
  MODIFY `chapterId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;COMMIT;