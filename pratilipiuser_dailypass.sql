-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: pratilipiuser
-- ------------------------------------------------------
-- Server version	5.7.34-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dailypass`
--

DROP TABLE IF EXISTS `dailypass`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dailypass` (
  `dailyPassId` int(11) NOT NULL AUTO_INCREMENT,
  `userId` int(11) DEFAULT NULL,
  `unlockedChapters` json DEFAULT NULL,
  PRIMARY KEY (`dailyPassId`),
  UNIQUE KEY `iddailyPass_UNIQUE` (`dailyPassId`),
  UNIQUE KEY `userId_UNIQUE` (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dailypass`
--

LOCK TABLES `dailypass` WRITE;
/*!40000 ALTER TABLE `dailypass` DISABLE KEYS */;
INSERT INTO `dailypass` VALUES (1,1,'{\"1\": \"2\", \"87\": \"3\", \"88\": \"3\"}'),(2,138,'{\"2\": \"3\", \"87\": \"3\", \"88\": \"3\"}'),(5,142,'{\"2\": \"3\", \"87\": \"3\", \"88\": \"3\"}'),(6,143,'{\"1\": 4, \"2\": 4, \"3\": 4, \"4\": 4, \"5\": 4, \"6\": 4, \"7\": 4, \"8\": 4, \"9\": 4, \"10\": 4, \"11\": 4, \"12\": 4, \"13\": 4, \"14\": 4, \"15\": 4, \"16\": 4, \"17\": 4, \"18\": 4, \"19\": 4, \"20\": 4, \"21\": 4, \"22\": 4, \"23\": 4, \"24\": 4, \"25\": 4, \"26\": 4, \"27\": 4, \"28\": 4, \"29\": 4, \"30\": 4, \"31\": 4, \"32\": 4, \"33\": 4, \"34\": 4, \"35\": 4, \"36\": 4, \"37\": 4, \"38\": 4, \"39\": 4, \"40\": 4, \"41\": 4, \"42\": 4, \"43\": 4, \"44\": 4, \"45\": 4, \"46\": 4, \"47\": 4, \"48\": 4, \"49\": 4, \"50\": 4, \"51\": 4, \"52\": 4, \"53\": 4, \"54\": 4, \"55\": 4, \"56\": 4, \"57\": 4, \"58\": 4, \"59\": 4, \"60\": 4, \"61\": 4, \"62\": 4, \"63\": 4, \"64\": 4, \"65\": 4, \"66\": 4, \"67\": 4, \"68\": 4, \"69\": 4, \"70\": 4, \"71\": 4, \"72\": 4, \"73\": 4, \"74\": 4, \"75\": 4, \"76\": 4, \"77\": 4, \"78\": 4, \"79\": 4, \"80\": 4, \"81\": 4, \"87\": \"3\", \"88\": \"3\"}'),(7,144,'{\"1\": 4, \"2\": 4, \"3\": 4, \"4\": 4, \"5\": 4, \"6\": 4, \"7\": 4, \"8\": 4, \"9\": 4, \"10\": 4, \"11\": 4, \"12\": 4, \"13\": 4, \"14\": 4, \"15\": 4, \"16\": 4, \"17\": 4, \"18\": 4, \"19\": 4, \"20\": 4, \"21\": 4, \"22\": 4, \"23\": 4, \"24\": 4, \"25\": 4, \"26\": 4, \"27\": 4, \"28\": 4, \"29\": 4, \"30\": 4, \"31\": 4, \"32\": 4, \"33\": 4, \"34\": 4, \"35\": 4, \"36\": 4, \"37\": 4, \"38\": 4, \"39\": 4, \"40\": 4, \"41\": 4, \"42\": 4, \"43\": 4, \"44\": 4, \"45\": 4, \"46\": 4, \"47\": 4, \"48\": 4, \"49\": 4, \"50\": 4, \"51\": 4, \"52\": 4, \"53\": 4, \"54\": 4, \"55\": 4, \"56\": 4, \"57\": 4, \"58\": 4, \"59\": 4, \"60\": 4, \"61\": 4, \"62\": 4, \"63\": 4, \"64\": 4, \"65\": 4, \"66\": 4, \"67\": 4, \"68\": 4, \"69\": 4, \"70\": 4, \"71\": 4, \"72\": 4, \"73\": 4, \"74\": 4, \"75\": 4, \"76\": 4, \"77\": 4, \"78\": 4, \"79\": 4, \"80\": 4, \"81\": 4, \"82\": 4, \"83\": 4, \"84\": 4, \"87\": \"3\", \"88\": \"3\"}'),(8,145,'{\"1\": 4, \"2\": 4, \"3\": 4, \"4\": 4, \"5\": 4, \"6\": 4, \"7\": 4, \"8\": 4, \"9\": 4, \"10\": 4, \"11\": 4, \"12\": 4, \"13\": 4, \"14\": 4, \"15\": 4, \"16\": 4, \"17\": 4, \"18\": 4, \"19\": 4, \"20\": 4, \"21\": 4, \"22\": 4, \"23\": 4, \"24\": 4, \"25\": 4, \"26\": 4, \"27\": 4, \"28\": 4, \"29\": 4, \"30\": 4, \"31\": 4, \"32\": 4, \"33\": 4, \"34\": 4, \"35\": 4, \"36\": 4, \"37\": 4, \"38\": 4, \"39\": 4, \"40\": 4, \"41\": 4, \"42\": 4, \"43\": 4, \"44\": 4, \"45\": 4, \"46\": 4, \"47\": 4, \"48\": 4, \"49\": 4, \"50\": 4, \"51\": 4, \"52\": 4, \"53\": 4, \"54\": 4, \"55\": 4, \"56\": 4, \"57\": 4, \"58\": 4, \"59\": 4, \"60\": 4, \"61\": 4, \"62\": 4, \"63\": 4, \"64\": 4, \"65\": 4, \"66\": 4, \"67\": 4, \"68\": 4, \"69\": 4, \"70\": 4, \"71\": 4, \"72\": 4, \"73\": 4, \"74\": 4, \"75\": 4, \"76\": 4, \"77\": 4, \"78\": 4, \"79\": 4, \"80\": 4, \"81\": 4, \"82\": 4, \"83\": 4, \"84\": 4, \"85\": 4, \"86\": 4, \"87\": 4, \"88\": 4}');
/*!40000 ALTER TABLE `dailypass` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-05 10:06:27