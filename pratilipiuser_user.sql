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
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `userId` int(11) NOT NULL AUTO_INCREMENT,
  `firstName` varchar(45) NOT NULL,
  `lastName` varchar(20) DEFAULT NULL,
  `penName` varchar(45) NOT NULL,
  `userEmail` varchar(50) NOT NULL,
  `bio` longtext,
  `number` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`userId`),
  UNIQUE KEY `id_UNIQUE` (`userId`),
  UNIQUE KEY `penName_UNIQUE` (`penName`),
  UNIQUE KEY `userEmail_UNIQUE` (`userEmail`)
) ENGINE=InnoDB AUTO_INCREMENT=146 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'Sdfd','dsf','sfd','FSDf','s','fd'),(12,'Sdfd','dsf','sfdu','FSDkf','s','fd    '),(13,'Sdfd','dsf','sfcd','FcSDf','s','fd    '),(14,'Sdfd','dsf','sfcdds','FcSDsdf','s','fd    '),(15,'Sdfd','ddsf','sfdd','dFSDf','s','fd    '),(16,'Sdfd','ddsf','sfddd','dFdSDf','s','fd    '),(17,'Sdfd','ddsf','sfgfddd','dFdSgfDf','s','fd    '),(18,'Sdfd','ddsf','sfgfdodd','dFdSugfDf','s','fd    '),(19,'Sdfd','ddsf','sfgfdodfchd','dFdSughjgfDf','s','fd    '),(20,'Sdfd','ddsf','sfgfdon,dfchd','dFdSbughjgfDf','s','fd    '),(21,'Sdfd','ddsf','sfgfdofchd','dFdughjgfDf','s','fd    '),(22,'Sdfd','ddsf','sfdofchd','dFdugfDf','s','fd    '),(23,'Sdfd','ddsf','sfdofchdui','dFdugifDf','s','fd    '),(24,'Sdfd','ddsf','sfdofchdudcsi','dFdugidcfDf','s','fd    '),(25,'Sdfd','ddsf','sfdofchdsi','dFdugisccfDf','s','fd    '),(26,'Sdfd','ddsf','sfdofchi','dFdugcfDf','s','fd    '),(27,'Sdfd','ddsf','sfdofcsdhi','dFdugcsdfDf','s','fd    '),(28,'Sdfd','ddsf','sfdofcsdhisa','dFdugcsdsfDf','s','fd    '),(29,'Sdfd','ddsf','sfdofcsdhisasc','dFdugcsdsfDscf','s','fd    '),(30,'Sdfd','ddsf','sfdofsdhisasc','dFducsdsfDscf','s','fd    '),(31,'Sdfd','ddsf','sfdofaxsdhsc','dFduccf','s','fd    '),(32,'Sdfd','ddsf','sfdofaxsdhsdsc','dFduccs','s','fd    '),(33,'Sdfd','ddsf','sfdofadhsdsc','dFdus','s','fd    '),(34,'Sdfd','ddsf','sfdadhsdsc','dFds','s','fd    '),(35,'Sdfd','ddsf','sfdsdadhsdsc','dFs','s','fd    '),(36,'Sdfd','ddsf','sfdsdacsdhsdsc','dFscs','s','fd    '),(37,'Sdfd','ddsf','sfdsdacsxdhsdsc','dFscsxz','s','fd    '),(38,'Sdfd','ddsf','sfdsdacsxdshsdsc','dFscsxzc','s','fd    '),(39,'Sdfd','ddsf','sfdsdacshsdsc','dFscassxzc','s','fd    '),(40,'Sdfd','ddsf','sfdsdacsdshsdsc','dFscasssdxzc','s','fd    '),(41,'Sdfd','ddsf','sfacsdshsdsc','dFsxzc','s','fd    '),(42,'Sdfd','ddsf','sfsdshsdsc','adFsxzc','s','fd    '),(43,'Sdfd','ddsf','sfshsdsc','Fsxzc','s','fd    '),(44,'Sdfd','ddsf','sfshssdsc','sxzc','s','fd    '),(45,'Sdfd','ddsf','sfshsscsdsc','sxzcs','s','fd    '),(46,'Sdfd','ddsf','sfshssscsdsc','sxszcs','s','fd    '),(47,'Sdfd','ddsf','sfshscsdsc','sxzcssds','s','fd    '),(48,'Sdfd','ddsf','sfshscsds','sxzcssdass','s','fd    '),(49,'Sdfd','ddsf','sfshscsdss','sxzcssdas','s','fd    '),(50,'Sdfd','ddsf','sfshscsads','sxzcsdas','s','fd    '),(51,'Sdfd','ddsf','sfshscassds','sxzsdas','s','fd    '),(52,'Sdfd','ddsf','sfshscascssds','sxzsdsas','s','fd    '),(53,'Sdfd','ddsf','sfshssds','sdfdzsdsas','s','fd    '),(54,'Sdfd','ddsf','sfshssdscs','sdfdzsds','s','fd    '),(55,'Sdfd','ddsf','sfshssdscxc','sdfdzsdcxs','s','fd    '),(56,'Sdfd','ddsf','sfshssdsdscxc','sdfdzss','s','fd    '),(57,'Sdfd','ddsf','sfshshsdscxc','sdfdzsthfds','s','fd    '),(58,'Sdfd','ddsf','sfshshsdsgucxc','sdfdzsttfhfds','s','fd    '),(59,'Sdfd','ddsf','sfshshsdsgufgjcxc','sdfdzstfds','s','fd    '),(60,'Sdfd','ddsfh','sfshshsdsgfghufgjcxc','sdfdzfhgstfds','s','fd    '),(61,'Sdfd','ddsfh','sfshsgjcxc','sdfdzftfds','s','fd    '),(62,'Sdfd','ddsffgjh','sfshsggfhjcxc','sdfdzfjfgtfds','s','fd    '),(63,'Sdfd','ddsffgjh','sfshsggfhdgjcxc','sdfdzfjdfgfgtfds','s','fd    '),(64,'Sdfd','ddsffgjh','sfshsggffdhdgjcxc','sdfdzfjdfgfddffgtfds','s','fd    '),(65,'Sdfd','ddsffgjh','sfshsggffdhgfhdgjcxc','sdfdzfjhgfdfgfddffgtfds','s','fd    '),(66,'Sdfd','ddsffgjh','sfshsggffdhgfhdgjctuyxc','sdfdzfjhgfrytdfgfddffgtfds','s','fd    '),(67,'Sdfd','ddsffgjh','sfshsggffdhgfhdgjctuydrtxc','sdfdzfjhgfryffgtfds','s','fd    '),(68,'Sdfd','ddsffgjh','sfshsggffdhgfhdgjctuydrc','sdfdzfjhgffgtfds','s','fd    '),(69,'Sdfd','ddsffgjh','sfshsggffdhygfhdgjctuydrtrc','sdfdzfjhgdtrffgtfds','s','fd    '),(70,'Sdfd','ddsffgjh','sfshsggffdhygfhdgjcrtrc','sdfdzfjhgdtrtyrffgtfds','s','fd    '),(71,'Sdfd','ddsffgjh','sfshsggffdhygfhdgjcvbcrtrc','sdfdzfjhgdtrtyrffgtbvfds','s','fd    '),(72,'Sdfd','ddsffgjh','sfshsggffdhygfhdgsdfjcvbcrtrc','sdfdzfjhgdtfdsrtyrffgtbvfds','s','fd    '),(73,'Sdfd','ddsffgjh','sfjcvbcrtrc','sdfdzfjhggtbvfds','s','fd    '),(74,'Sdfd','ddsffgjh','sfjcvbcrtrcyg','sdfdzfjhgyggtbvfds','s','fd    '),(75,'Sdfd','ddsffgjh','sfjcvbcrtrcygdrt','sdfdzfjhgyggtbvdrtfds','s','fd    '),(76,'Sdfd','ddsffgjh','sfjcvbcrtrcyfytgdrt','sdfdzfjhgygfytgtbvdrtfds','s','fd    '),(77,'Sdfd','ddsffgjh','sfjcvbcrtrcyfytgy8gdrt','sdfdzfjhgygfytgtbv8gydrtfds','s','fd    '),(78,'Sdfd','ddsffgjh','sfjcvbcrtrcyfytgy878ygdrt','sdfdzfjhgygfytgtbv8gyd7yrtfds','s','fd    '),(79,'Sdfd','ddsffgjh','sfjcvbcrtrcyfytgyiu878ygdrt','sdfdzfjhgygfytgtbv8gyd7yiuirtfds','s','fd    '),(80,'Sdfd','ddsffgjh','sfjcvbcrtrcyfytgyiuok878ygdrt','sdfdzfjhgygyd7yiuirtfds','s','fd    '),(81,'Sdfd','ddsffgjh','sfjcvbcrtrcyf8ygdrt','sdfdzfjhgydgyd7yiuirtfds','s','fd    '),(82,'Sdfd','ddsffgjh','sfjcvbcr8ygdrt','sdfdzfjhgirtfds','s','fd    '),(83,'Sdfd','ddsffgjh','sfygdrt','sdfdzfjirtfds','s','fd    '),(84,'Sdfd','ddsffgjh','sfffghygdrt','sdfdzfjirfghtfds','s','fd    '),(85,'Sdfd','ddsffgjh','sfffghygdhjrt','sdfdzfjirfghjtfds','s','fd    '),(86,'Sdfd','ddsffgjh','sfffghygdhjgjrt','sdfdzfjirfggjh]hjtfds','s','fd    '),(87,'Sdfd','ddsffgjh','sfffghygdhjgjrthij','sdfdzfjirfggjh]tfds','s','fd    '),(88,'Sdfd','ddsffgjh','sfffghygdthij','sdfdzfjifds','s','fd    '),(89,'Sdfd','ddsffgjh','sfffghygdthiergj','sdfdzfjiergfds','s','fd    '),(90,'Sdfd','ddsffgjh','sfffghygdfsdthiergj','sdfdfsdzfjiergfds','s','fd    '),(91,'Sdfd','ddsffgjh','sfffghygdfsfsddthiergj','sdfdfsdzfjiedsfrgfds','s','fd    '),(92,'Sdfd','ddsffgjh','sfffghygdfsfsddthdfsiergj','sdfdfsdzfjiedsffdsrgfds','s','fd    '),(93,'Sdfd','ddsffgjh','sfffghygdfsfsddth98dfsiergj','sdfdfsdzfjieds89ffdsrgfds','s','fd    '),(94,'Sdfd','ddsffgjh','sfffghygdfsfsddth98hgjdfsiergj','sdfdfsdzfjieds89ffdghjsrgfds','s','fd    '),(95,'Sdfd','ddsffgjh','sfffghygdfsfsiddth98hgjdfsiergj','sdfdfsihudzfjieds89ffdghjsrgfds','s','fd    '),(96,'Sdfd','ddsffgjh','sfffghygdfsfsiyuoddth98hgjdfsiergj',':jhgjgh','s','fd    '),(97,'Sdfd','ddsffgjh','sffth98hgjdfsiergj',':jhgjghhui','s','fd    '),(98,'Sdfd','ddsffgjh','sfgjdfsiergj',':jhgghhui','s','fd    '),(99,'Sdfd','ddsffgjh','siergj',':jhgyuighhui','s','fd    '),(100,'Sdfd','ddsffgjh','sil;iergj',':jhgyuig;khhui','s','fd    '),(101,'Sdfd','ddsffgjh','sil;ierkgj',':jhgyuig;kpohhui','s','fd    '),(102,'Sdfd','ddsffgjh','sil;ierkigj',':jhgyuig;kpiuohhui','s','fd    '),(103,'Sdfd','ddsffgjh','sil;iedsfrkigj',':jhgyuig;kpiufsdohhui','s','fd    '),(104,'Sdfd','ddsffgjh','sil;iedsfgfhrkigj',':jhgyuig;kpiufsfghdohhui','s','fd    '),(105,'Sdfd','ddsffgjh','sil;iedsfgfghfhrkigj',':jhgyuig;kpiufhgfsfghdohhui','s','fd    '),(106,'Sdfd','ddsffgjh','sil;iedsfgfghrytfhrkigj',':jhgyuig;yr','s','fd    '),(107,'Sdfd','ddsffgjh','sil;iedsfgfghuiorytfhrkigj',':jhgyuig;uioyr','s','fd    '),(108,'Sdfd','ddsffgjh','sil;iedsfgfghugdfgdfiorytfhrkigj',':jhgyuig;uiogfdyr','s','fd    '),(109,'Sdfd','ddsffgjh','sil;iedsgdfgfghugdfgdfiorytfhrkigj',':jhgyuigdgf;uiogfdyr','s','fd    '),(110,'Sdfd','ddsffgjh','sil;iedsgdfggdffghugdfgdfiorytfhrkigj',':jhgyuigdgfdgf;uiogfdyr','s','fd    '),(111,'Sdfd','ddsffgjh','sil;ifghugdfgdfiorytfhrkigj',':jhgf;uiogfdyr','s','fd    '),(112,'Sdfd','ddsffgjh','sil;ifghugdfgtredfiorytfhrkigj',':jhtregf;uiogfdyr','s','fd    '),(113,'Sdfd','ddsffgjh','sil;ifghugdfgtrefgdfiorytfhrkigj',':jhtregf;uiofggfdyr','s','fd    '),(114,'Sdfd','ddsffgjh','sil;ifghugdfgtrefgrytfhrkigj',':jhtregf;uiyr','s','fd    '),(115,'Sdfd','ddsffgjh','sil;ifghugdfgtrefglk;rytfhrkigj',':jhtregf;uik;llk;yr','s','fd    '),(116,'Sdfd','ddsffgjh','sil;ifghugdfgtrefglk;hjgrytfhrkigj',':jhtregf;uik;llk;yhgjr','s','fd    '),(117,'Sdfd','ddsffgjh','sil;ifghulk;hjgrytfhrkigj',':jf;uik;llk;yhgjr','s','fd    '),(118,'Sdfd','ddsffgjh','silk;hjgrytfhrkigj','f;uik;llk;yhgjr','s','fd    '),(119,'Sdfd','ddsffgjh','sielk;hjgrytfhrkigj','fde;uik;llk;yhgjr','s','fd    '),(120,'Sdfd','ddsffgjh','sielk;hsdjgrytfhrkigj','fdeds;uik;llk;yhgjr','s','fd    '),(121,'Sdfd','ddsffgjh','sielk;hsdjgrytfkigj','fdeds;uik;yhgjr','s','fd    '),(122,'Sdfd','ddsffgjh','sielk;hsdjgryhgtfkigj','fdeds;uikgh;yhgjr','s','fd    '),(123,'Sdfd','ddsffgjh','sielk;hsdjgryhgdctfkigj','fdeds;uikgh;yhgdjr','s','fd    '),(124,'Sdfd','ddsffgjh','sielk;yhgdctfkigj','fdedikgh;yhgdjr','s','fd    '),(125,'Sdfd','ddsffgjh','sie;yhgdctfkigj','fdeikgh;yhgdjr','s','fd    '),(126,'Sdfd','ddsffgjh','sie;yhgdctgj','fdeikyhgdjr','s','fd    '),(127,'Sdfd','ddsffgjh','sie;yhgdcscgj','fdeiksdyhgdjr','s','fd    '),(128,'Sdfd','ddsffgjh','sie;yhgdcscdszxgj','fdeiksdyhgxzdjr','s','fd    '),(129,'Sdfd','ddsffgjh','sie;yhgdcscdcszxgj','fdeiksdyhcxgxzdjr','s','fd    '),(130,'Sdfd','ddsffgjh','sie;yhgdcscdcdsszxgj','fdeiksdyhcxgxsdzdjr','s','fd    '),(131,'Sdfd','ddsffgjh','sie;yhgdcscdcdsszxgjsc','fdeiksdyhcxgxsdsddszdjr','s','fd    '),(132,'Sdfd','ddsffgjh','sie;yhgdcscdcdsszxgjsdc','fdeiksdyhcxgxsdsddszsdjr','s','fd    '),(133,'Sdfd','ddsffgjh','sie;yhgdcdcdsszxgjsdc','fdeiksdyhcxgxsdjr','s','fd    '),(134,'Sdfd','ddsffgjh','sie;yhgcdcdcdsszxgjsdc','fdeikssfdyhcxgxsdjr','s','fd    '),(135,'Sdfd','ddsffgjh','sie;yhgcdcdacdsszxgjsdc','fdeikssfddyhcxgxsdjr','s','fd    '),(136,'Sdfd','ddsffgjh','sie;ycdsszxgjsdc','fdikssfddyhcxgxsdjr','s','fd    '),(137,'Sdfd','ddsffgjh','siycdsszxgjsdc','fdsfddyhcxgxsdjr','s','fd    '),(138,'Sdfd','ddsffgjh','sidscdsszxgjsdc','fdsdsfddyhcxgxsdjr','s','fd    '),(139,'Sdfd','ddsffgjh','sidscdsssdzxgjsdc','ds','s','fd    '),(140,'Sdfd','ddsffgjh','siddsssdzxgjsdc','dss','s','fd    '),(141,'Sdfd','ddsffgjh','siddsdc','dssc','s','fd    '),(142,'Sdfd','ddsffgjh','siddsdhc','dssckn','s','fd    '),(143,'Sdfd','ddsffgjh','siddsdshc','dssccsdkn','s','fd    '),(144,'Sdfd','ddsffgjh','siddsdshcs','dssccssdkn','s','fd    '),(145,'Sdfd','ddsffgjh','vsiddsdshcs','vdssccssdkn','s','fd    ');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
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
