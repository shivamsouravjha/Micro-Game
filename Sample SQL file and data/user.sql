SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";
CREATE TABLE `user` (
  `userId` int(11) NOT NULL,
  `firstName` varchar(45) NOT NULL,
  `lastName` varchar(20) DEFAULT NULL,
  `penName` varchar(45) NOT NULL,
  `userEmail` varchar(50) NOT NULL,
  `bio` longtext,
  `number` varchar(10) DEFAULT NULL,
  `password` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `user` (`userId`, `firstName`, `lastName`, `penName`, `userEmail`, `bio`, `number`, `password`) VALUES
(7, 'Sdfd', 'ddsffgjh', 'reggreg', 'gregreg', 's', 'fd    ', '$2a$10$PK7Ly2O1vtZIfQWeu5IpCOXHX241Z9WAaI4rgItX0xM'),
(8, 'Sdfd', 'ddsffgjh', 'reggdsreg', 'grdsegreg', 's', 'fd    ', '$2a$10$Yma7Q7yztagTWXbCyYCfIedTS5TTYN8zhlOOZ8ZotFM'),
(9, 'Sdfd', 'ddsffgjh', 'mm', 'gmrdsegreg', 's', 'fd    ', '$2a$10$fItPIsU5bewyJymrEWv9ueQ/Ugjuhi.idOMLzNuxzFYlPEh79LCgu'),
(10, 'Sdfd', 'ddsffgjh', 'mmfdggf', 'gmrgfddsegreg', 's', 'fd    ', '$2a$10$OSkgaSngwPhW99ONrlmYC.PAIDbCjYcEKfI4D57IOmir6Za93hdWu');
ALTER TABLE `user`
  ADD PRIMARY KEY (`userId`),
  ADD UNIQUE KEY `id_UNIQUE` (`userId`),
  ADD UNIQUE KEY `penName_UNIQUE` (`penName`),
  ADD UNIQUE KEY `userEmail_UNIQUE` (`userEmail`);
ALTER TABLE `user`
  MODIFY `userId` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;COMMIT;