-- MySQL dump 10.13  Distrib 5.5.34, for osx10.6 (i386)
--
-- Host: localhost    Database: concertsap
-- ------------------------------------------------------
-- Server version	5.5.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `band`
--

DROP TABLE IF EXISTS `band`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `band` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(25555) DEFAULT NULL,
  `website` varchar(2555) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band`
--

LOCK TABLES `band` WRITE;
/*!40000 ALTER TABLE `band` DISABLE KEYS */;
INSERT INTO `band` VALUES (1,'Red Hot Chilli Peppers','boom.com'),(2,'Marcus Band','marcus.com'),(3,'TEST','bam'),(4,'yo yo','test');
/*!40000 ALTER TABLE `band` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `band_concert`
--

DROP TABLE IF EXISTS `band_concert`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `band_concert` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `band_id` int(11) DEFAULT NULL,
  `concert_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `band_concert` (`band_id`,`concert_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band_concert`
--

LOCK TABLES `band_concert` WRITE;
/*!40000 ALTER TABLE `band_concert` DISABLE KEYS */;
/*!40000 ALTER TABLE `band_concert` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `band_concert_record`
--

DROP TABLE IF EXISTS `band_concert_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `band_concert_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `band_id` int(11) DEFAULT NULL,
  `concert_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `time` time DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `band_concert_date_time` (`band_id`,`concert_id`,`date`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band_concert_record`
--

LOCK TABLES `band_concert_record` WRITE;
/*!40000 ALTER TABLE `band_concert_record` DISABLE KEYS */;
/*!40000 ALTER TABLE `band_concert_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `concert`
--

DROP TABLE IF EXISTS `concert`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `concert` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(25555) DEFAULT NULL,
  `state_id` int(11) DEFAULT NULL,
  `website` varchar(2555) DEFAULT NULL,
  `start` datetime DEFAULT NULL,
  `end` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `concert`
--

LOCK TABLES `concert` WRITE;
/*!40000 ALTER TABLE `concert` DISABLE KEYS */;
INSERT INTO `concert` VALUES (1,'coachella','boom',1,'bam','2014-03-11 00:00:00','2014-03-20 00:00:00'),(2,'test','123 test street',1,'www.test.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(3,'bam','123 bam street',1,'www.bam.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(4,'test','123 test street',1,'www.test.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(5,'BOOM','123 bam street',1,'www.bam.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(6,'test','123 test street',1,'www.test.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(7,'BOOM','123 bam street',1,'www.bam.com','0000-00-00 00:00:00','0000-00-00 00:00:00'),(8,'BAM','',13,'','0000-00-00 00:00:00','0000-00-00 00:00:00'),(9,'TEST RIGHT NOW','',13,'','0000-00-00 00:00:00','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `concert` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `retailer`
--

DROP TABLE IF EXISTS `retailer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `retailer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `website` varchar(2555) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `retailer`
--

LOCK TABLES `retailer` WRITE;
/*!40000 ALTER TABLE `retailer` DISABLE KEYS */;
INSERT INTO `retailer` VALUES (1,'stubhub','www.stubhub.com');
/*!40000 ALTER TABLE `retailer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `state`
--

DROP TABLE IF EXISTS `state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `state` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `acronym` varchar(2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `state` (`acronym`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `state`
--

LOCK TABLES `state` WRITE;
/*!40000 ALTER TABLE `state` DISABLE KEYS */;
INSERT INTO `state` VALUES (1,'North Carolina','NC'),(2,'South Carolina','SC'),(3,'Florida','FL'),(4,'Georgia','GA'),(5,'California','CA'),(11,'Virginia','VA'),(12,'Washington','WA'),(13,'Arizona','AZ'),(14,'Colorado','CO'),(15,'BAM','BA');
/*!40000 ALTER TABLE `state` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ticket_record`
--

DROP TABLE IF EXISTS `ticket_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ticket_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `price` double DEFAULT NULL,
  `concert_id` int(11) DEFAULT NULL,
  `retailer_id` int(11) DEFAULT NULL,
  `timestamp` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ticket_record`
--

LOCK TABLES `ticket_record` WRITE;
/*!40000 ALTER TABLE `ticket_record` DISABLE KEYS */;
INSERT INTO `ticket_record` VALUES (1,420,1,1,'2014-03-11 00:00:00'),(2,470,1,1,'2014-03-12 00:00:00'),(3,495,1,1,'2014-03-13 00:00:00'),(4,450,1,1,'2014-03-11 00:00:00'),(5,505,1,1,'2014-03-11 00:00:00'),(6,535,1,1,'2014-03-11 00:00:00'),(7,505,1,1,'2014-03-12 00:00:00'),(8,535,1,1,'2014-03-12 00:00:00'),(9,450,1,1,'2014-03-13 00:00:00'),(10,505,1,1,'2014-03-13 00:00:00'),(11,535,1,1,'2014-03-13 00:00:00'),(12,123321,1,1,'2014-03-14 21:07:43'),(13,32,1,1,'2014-03-14 21:32:40'),(14,2323,1,1,'0000-00-00 00:00:00');
/*!40000 ALTER TABLE `ticket_record` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-03-29 10:00:20
