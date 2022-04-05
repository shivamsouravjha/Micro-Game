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
-- Table structure for table `content`
--

DROP TABLE IF EXISTS `content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `content` (
  `chapterId` int(11) NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `story` longtext,
  `seriesId` varchar(45) NOT NULL,
  PRIMARY KEY (`chapterId`),
  UNIQUE KEY `contentId_UNIQUE` (`chapterId`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content`
--

LOCK TABLES `content` WRITE;
/*!40000 ALTER TABLE `content` DISABLE KEYS */;
INSERT INTO `content` VALUES (1,'cadf','fasf','14'),(2,'Sfsssssdsf','sdfdfdfds','14'),(3,'Sfdsf','sdfdfds','14'),(4,'Sfsssssdsf','sdfdfdfds','14'),(5,'Sfdsf','sdfdfds','14'),(6,'Sfsssssdsf','sdfdfdfds','14'),(7,'Sfdsf','sdfdfds','14'),(8,'Sfsssssdsf','sdfdfdfds','14'),(9,'Sfdsf','sdfdfds','14'),(10,'Sfsssssdsf','sdfdfdfds','14'),(11,'Sfdsf','sdfdfds','14'),(12,'Sfsssssdsf','sdfdfdfds','14'),(13,'Sfdsf','sdfdfds','14'),(14,'Sfsssssdsf','sdfdfdfds','14'),(15,'Sfdsf','sdfdfds','2'),(16,'Sfsssssdsf','sdfdfdfds','14'),(17,'Sfdsf','sdfdfds','2'),(18,'Sfdsf','sdfdfds','2'),(19,'Sfdsf','sdfdfds','2'),(20,'Sfdsf','sdfdfds','2'),(21,'Sfdsf','sdfdfds','2'),(22,'Sfdsf','sdfdfds','2'),(23,'Sfdsf','sdfdfds','2'),(24,'Sfdsf','sdfdfds','2');
/*!40000 ALTER TABLE `content` ENABLE KEYS */;
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
