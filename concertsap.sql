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
) ENGINE=InnoDB AUTO_INCREMENT=335 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band`
--

LOCK TABLES `band` WRITE;
/*!40000 ALTER TABLE `band` DISABLE KEYS */;
INSERT INTO `band` VALUES (1,'EMINEM',''),(2,'OUTKAST',''),(3,'KINGS OF LEON',''),(4,'ARCTIC MONKEYS',''),(5,'SKRILLEX',''),(6,'CALVIN HARRIS',''),(7,'LORDE',''),(8,'THE AVETT BROTHERS',''),(9,'FOSTER THE PEOPLE',''),(10,'ZEDD',''),(11,'SEBASTIAN INGROSSO',''),(12,'KREWELLA',''),(13,'NAS',''),(14,'ABOVE & BEYOND',''),(15,'CHILDISH GAMBINO',''),(16,'BROKEN BELLS',''),(17,'SPOON',''),(18,'THE HEAD AND THE HEART',''),(19,'CAGE THE ELEPHANT',''),(20,'YOUNG THE GIANT',''),(21,'CHASE & STATUS',''),(22,'INTERPOL',''),(23,'LYKKE LI',''),(24,'CHANCE THE RAPPER',''),(25,'PHANTOGRAM',''),(26,'CUT / COPY',''),(27,'FLOSSTRADAMUS',''),(28,'FITZ & THE TANTRUMS',''),(29,'GROUPLOVE',''),(30,'CHVRCHES',''),(31,'THE GLITCH MOB',''),(32,'PORTUGAL.THE MAN',''),(33,'GRAMATIK',''),(34,'CHROMEO',''),(35,'DARKSIDE',''),(36,'THE 1975',''),(37,'DIMITRI VEGAS & LIKE MIKE',''),(38,'REBELUTION',''),(39,'GLEN HANSARD',''),(40,'JENNY LEWIS',''),(41,'MANCHESTER ORCHESTRA',''),(42,'GESAFFELSTEIN',''),(43,'JOHN BUTLER TRIO',''),(44,'FLUME',''),(45,'MARTIN GARRIX',''),(46,'AFI',''),(47,'J. RODDY WALSTON & THE BUSINESS',''),(48,'THE KOOKS',''),(49,'TROMBONE SHORTY & ORLEANS AVENUE',''),(50,'RUN THE JEWELS',''),(51,'THE AIRBORNE TOXIC EVENT',''),(52,'RUDIMENTAL',''),(53,'PHOSPHORESCENT',''),(54,'DUKE DUMONT',''),(55,'LONDON GRAMMAR',''),(56,'THE TEMPER TRAP',''),(57,'IGGY AZALEA',''),(58,'JOACHIM GARRAUD',''),(59,'BOMBAY BICYCLE CLUB',''),(60,'WHITE DENIM',''),(61,'TYPHOON',''),(62,'Z-TRIP',''),(63,'WARPAINT',''),(64,'KATE NASH',''),(65,'KODALINE',''),(66,'SANDER KLEINENBERG',''),(67,'VANCE JOY',''),(68,'LUCIUS',''),(69,'BLOOD ORANGE',''),(70,'RICH HOMIE QUAN',''),(71,'PERRY/ETTY VS JOACHIM GARRAUD',''),(72,'DELTA RAE',''),(73,'RAC',''),(74,'JAGWAR MA',''),(75,'PARQUET COURTS',''),(76,'GTA',''),(77,'JHENÉ AIKO',''),(78,'TEMPLES',''),(79,'SMALLPOOLS',''),(80,'JUNGLE',''),(81,'KONGOS',''),(82,'NONONO',''),(83,'WILDCAT! WILDCAT!',''),(84,'PAPA',''),(85,'GEMINI CLUB',''),(86,'THE DISTRICTS',''),(87,'VIC MENSA',''),(88,'CRIZZLY',''),(89,'BOMBA ESTÉREO',''),(90,'FRANCISCA VALENZUELA',''),(91,'BLEACHERS',''),(92,'INTO IT. OVER IT.',''),(93,'HEROBUST',''),(94,'COURTNEY BARNETT',''),(95,'SAN FERMIN',''),(96,'ROYAL BLOOD',''),(97,'RATKING',''),(98,'BEAR HANDS',''),(99,'ROADKILL GHOST CHOIR',''),(100,'MEG MYERS',''),(101,'KAUSEA',''),(102,'BRILLZ',''),(103,'BRONZE RADIO RETURN',''),(104,'THE SO SO GLOS',''),(105,'CASH CASH',''),(106,'JACOB PLANT',''),(107,'LINDSAY LOWEND',''),(108,'JOYWAVE',''),(109,'BEBE REXHA',''),(110,'DESERT NOISES',''),(111,'BETTY WHO',''),(112,'JON BATISTE AND STAY HUMAN',''),(113,'BENJAMIN BOOKER',''),(114,'FLY GOLDEN EAGLE',''),(115,'DUGAS',''),(116,'CRASS MAMMOTH',''),(117,'THE LAST INTERNATIONALE',''),(118,'WALLPAPER.',''),(119,'JOHNNYSWIM',''),(120,'CARDIKNOX',''),(121,'CHARLIE HIRSCH',''),(122,'BAGHEERA',''),(123,'SPACE CAPONE',''),(124,'OF VERONA',''),(125,'ANNA LUNOE',''),(126,'HIGHLY SUSPECT',''),(127,'PLASTIC VISIONS',''),(128,'OYINDA',''),(129,'ROCKY BUSINESS',''),(130,'The Knife',''),(131,'The Replacements',''),(132,'Girl Talk',''),(133,'Ellie Goulding',''),(134,'HAIM',''),(135,'Neko Case',''),(136,'Bonobo',''),(137,'Bryan Ferry',''),(138,'The Afghan Whigs',''),(139,'The Cult',''),(140,'Bastille',''),(141,'Aloe Blacc',''),(142,'A$AP Ferg',''),(143,'Woodkid',''),(144,'Carnage',''),(145,'Shlomo',''),(146,'Gareth Emery',''),(147,'Michael Brun',''),(148,'MS MR',''),(149,'Hot Since 82',''),(150,'Damian Lazarus',''),(151,'GOAT',''),(152,'Nina Kraviz',''),(153,'Anthony Green',''),(154,'The Jon Spencer Blues Explosion',''),(155,'Solomun',''),(156,'ZZ Ward',''),(157,'Anti-Flag',''),(158,'Caravan Palace',''),(159,'Flatbush Zombies',''),(160,'Deorro',''),(161,'Waxahatchee',''),(162,'Title Fight',''),(163,'Davide Squillace',''),(164,'DJ Falcon',''),(165,'Dum Dum Girls',''),(166,'Austra',''),(167,'Tom Odell',''),(168,'Dixon',''),(169,'Wye Oak',''),(170,'Crosses',''),(171,'Mako',''),(172,'The Preatures',''),(173,'The Bots',''),(174,'Gabba Gabba Heys',''),(175,'Muse',''),(176,'Queens Of The Stone Age',''),(177,'Pharrell Williams',''),(178,'Pet Shop Boys',''),(179,'MGMT',''),(180,'Empire Of The Sun',''),(181,'Fatboy Slim',''),(182,'Kid Cudi',''),(183,'Sleigh Bells',''),(184,'City And Colour',''),(185,'Dillon Francis',''),(186,'Capital Cities',''),(187,'The Naked And Famous',''),(188,'Mogwai',''),(189,'Solange',''),(190,'Washed Out',''),(191,'Future Islands',''),(192,'Ty Segal',''),(193,'Banks',''),(194,'Tiga',''),(195,'Holy Ghost!',''),(196,'Netsky',''),(197,'RL Grime',''),(198,'Galantis',''),(199,'Foxygen',''),(200,'White Lies',''),(201,'Graveyard',''),(202,'The Internet',''),(203,'Laura Mvula',''),(204,'The Dismemberment Plan',''),(205,'Headhunterz',''),(206,'TJR',''),(207,'Cajmere',''),(208,'Guy Gerber',''),(209,'Nicole Moudaber',''),(210,'MAKJ',''),(211,'The Magician',''),(212,'Young & Sick',''),(213,'Unlocking The Truth',''),(214,'Saints Of Valory',''),(215,'Carbon Airways',''),(216,'UZ',''),(217,'Syd Arthur',''),(218,'Bicep',''),(219,'Drowners',''),(220,'Arcade Fire',''),(221,'Beck',''),(222,'Neutral Milk Hotel',''),(223,'Disclosure',''),(224,'Lana Del Rey',''),(225,'Motorhead',''),(226,'Alesso',''),(227,'Duck Sauce',''),(228,'Little Dragon',''),(229,'Beady Eye',''),(230,'The Toy Dolls',''),(231,'Adventure Club',''),(232,'Big Gigantic',''),(233,'Laurent Garnier',''),(234,'STRFKR',''),(235,'Fishbone',''),(236,'Trombone Shorty',''),(237,'AlunaGeorge',''),(238,'Art Department',''),(239,'Flight Facilities',''),(240,'Frank Turner',''),(241,'John Newman',''),(242,'Maceo Plex',''),(243,'Superchunk',''),(244,'Bombino',''),(245,'Daughter',''),(246,'Bad Manners',''),(247,'Surfer Blood',''),(248,'Lee Burridge',''),(249,'Poolside',''),(250,'Classixx',''),(251,'Showtek',''),(252,'James Vincent McMorrow',''),(253,'Bo Ningen',''),(254,'Aeroplane',''),(255,'Factory Floor',''),(256,'Preservation Hall Jazz Band',''),(257,'The Martinez Brothers',''),(258,'Scuba',''),(259,'John Beaver',''),(260,'12TH PLANET',''),(261,'2MANYDJS',''),(262,'AFROBETA',''),(263,'AFROJACK',''),(264,'ALVIN RISK',''),(265,'ANDY C',''),(266,'ARMIN VAN BUUREN',''),(267,'AVICII',''),(268,'BASEMENT JAXX',''),(269,'BEEN TRILL',''),(270,'BORGORE',''),(271,'BRO SAFARI',''),(272,'CARL COX',''),(273,'CHASE AND STATUS',''),(274,'CEDRIC GERVAIS',''),(275,'CLAUDE VONSTROKE',''),(276,'COSMIC GATE',''),(277,'CRAZE',''),(278,'CUT COPY',''),(279,'CYRIL HAHN',''),(280,'DANNY AVILA',''),(281,'DATSIK',''),(282,'DAVID GUETTA',''),(283,'DIPLO',''),(284,'DIRTYPHONICS',''),(285,'DIZZEE RASCAL',''),(286,'DJ ICEY',''),(287,'DON DIABLO',''),(288,'DUBFIRE',''),(289,'ERIC PRYDZ',''),(290,'EXAMPLE',''),(291,'FEDDE LE GRAND',''),(292,'FIGURE',''),(293,'GOLDFISH',''),(294,'HARDWELL',''),(295,'JAMIE JONES',''),(296,'JOHN O\'CALLAGHAN',''),(297,'JULIO BASHMORE',''),(298,'JUST BLAZE',''),(299,'JUSTIN MARTIN',''),(300,'KASKADE',''),(301,'KILL FRENZY',''),(302,'KILL THE NOISE',''),(303,'LAIDBACK LUKE',''),(304,'LOCO DICE',''),(305,'LUCIANO',''),(306,'MARCO CAROLA',''),(307,'MATT ZO',''),(308,'MIA',''),(309,'MOON BOOTS',''),(310,'NERVO',''),(311,'NEW WORLD PUNX',''),(312,'NICKY ROMERO',''),(313,'OLIVER',''),(314,'PAPER DIAMOND',''),(315,'PAUL KALKBRENNER',''),(316,'PAUL VAN DYK',''),(317,'PETE TONG',''),(318,'PUSHA T',''),(319,'RIFF RAFF',''),(320,'SANDER VAN DOOM',''),(321,'SOUL CLAP',''),(322,'STAFFORD BROTHERS',''),(323,'STATON WARRIORS',''),(324,'STEVE ANGELLO',''),(325,'SUB FOCUS',''),(326,'SUNNERY JAMES & RYAN MARCIANO',''),(327,'TIESTO',''),(328,'TOMMIE SUNSHINE',''),(329,'TOURIST',''),(330,'TRENTEMOLLER',''),(331,'UMEK',''),(332,'VICEROY',''),(333,'WAKA FLOCKA FLAME',''),(334,'ZOMBOY','');
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
) ENGINE=InnoDB AUTO_INCREMENT=397 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band_concert`
--

LOCK TABLES `band_concert` WRITE;
/*!40000 ALTER TABLE `band_concert` DISABLE KEYS */;
INSERT INTO `band_concert` VALUES (1,1,2),(2,2,2),(130,2,3),(3,3,2),(4,4,2),(5,5,2),(189,5,3),(6,6,2),(248,6,3),(7,7,2),(191,7,3),(8,8,2),(9,9,2),(192,9,3),(10,10,2),(134,10,3),(395,10,4),(11,11,2),(12,12,2),(264,12,3),(354,12,4),(13,13,2),(197,13,3),(14,14,2),(298,14,4),(15,15,2),(16,16,2),(133,16,3),(17,17,2),(18,18,2),(199,18,3),(19,19,2),(201,19,3),(20,20,2),(21,21,2),(22,22,2),(23,23,2),(24,24,2),(262,24,3),(316,24,4),(25,25,2),(26,26,2),(27,27,2),(257,27,3),(339,27,4),(28,28,2),(29,29,2),(152,29,3),(30,30,2),(203,30,3),(31,31,2),(144,31,3),(387,31,4),(32,32,2),(33,33,2),(343,33,4),(34,34,2),(137,34,3),(35,35,2),(214,35,3),(36,36,2),(259,36,3),(37,37,2),(38,38,2),(39,39,2),(40,40,2),(41,41,2),(42,42,2),(341,42,4),(43,43,2),(44,44,2),(148,44,3),(45,45,2),(141,45,3),(46,46,2),(140,46,3),(47,47,2),(289,47,3),(48,48,2),(49,49,2),(50,50,2),(51,51,2),(52,52,2),(265,52,3),(53,53,2),(54,54,2),(165,54,3),(55,55,2),(56,56,2),(57,57,2),(58,58,2),(59,59,2),(217,59,3),(60,60,2),(61,61,2),(62,62,2),(63,63,2),(209,63,3),(64,64,2),(159,64,3),(65,65,2),(66,66,2),(67,67,2),(68,68,2),(69,69,2),(229,69,3),(70,70,2),(71,71,2),(72,72,2),(73,73,2),(375,73,4),(74,74,2),(150,74,3),(75,75,2),(76,76,2),(230,76,3),(344,76,4),(77,77,2),(288,77,3),(78,78,2),(207,78,3),(79,79,2),(80,80,2),(81,81,2),(82,82,2),(83,83,2),(84,84,2),(85,85,2),(86,86,2),(87,87,2),(88,88,2),(320,88,4),(89,89,2),(90,90,2),(91,91,2),(92,92,2),(93,93,2),(94,94,2),(95,95,2),(96,96,2),(97,97,2),(287,97,3),(98,98,2),(236,98,3),(99,99,2),(100,100,2),(101,101,2),(102,102,2),(103,103,2),(104,104,2),(105,105,2),(106,106,2),(107,107,2),(108,108,2),(109,109,2),(110,110,2),(111,111,2),(112,112,2),(113,113,2),(114,114,2),(115,115,2),(116,116,2),(117,117,2),(118,118,2),(119,119,2),(120,120,2),(121,121,2),(122,122,2),(123,123,2),(124,124,2),(125,125,2),(292,125,3),(126,126,2),(127,127,2),(128,128,2),(129,129,2),(131,130,3),(132,131,3),(135,132,3),(136,133,3),(138,134,3),(139,135,3),(142,136,3),(143,137,3),(145,138,3),(146,139,3),(147,140,3),(149,141,3),(151,142,3),(153,143,3),(154,144,3),(313,144,4),(155,145,3),(156,146,3),(340,146,4),(157,147,3),(158,148,3),(160,149,3),(161,150,3),(323,150,4),(162,151,3),(163,152,3),(164,153,3),(166,154,3),(167,155,3),(380,155,4),(168,156,3),(169,157,3),(170,158,3),(171,159,3),(172,160,3),(173,161,3),(174,162,3),(175,163,3),(176,164,3),(177,165,3),(178,166,3),(179,167,3),(180,168,3),(181,169,3),(182,170,3),(183,171,3),(184,172,3),(185,173,3),(186,174,3),(187,175,3),(188,176,3),(190,177,3),(193,178,3),(194,179,3),(362,179,4),(195,180,3),(334,180,4),(196,181,3),(198,182,3),(200,183,3),(202,184,3),(204,185,3),(327,185,4),(205,186,3),(206,187,3),(208,188,3),(210,189,3),(211,190,3),(212,191,3),(213,192,3),(215,193,3),(216,194,3),(218,195,3),(219,196,3),(366,196,4),(220,197,3),(376,197,4),(221,198,3),(222,199,3),(223,200,3),(224,201,3),(225,202,3),(226,203,3),(227,204,3),(228,205,3),(231,206,3),(232,207,3),(233,208,3),(234,209,3),(235,210,3),(237,211,3),(238,212,3),(239,213,3),(240,214,3),(241,215,3),(242,216,3),(243,217,3),(244,218,3),(245,219,3),(246,220,3),(247,221,3),(249,222,3),(250,223,3),(251,224,3),(252,225,3),(253,226,3),(302,226,4),(254,227,3),(255,228,3),(256,229,3),(258,230,3),(260,231,3),(299,231,4),(261,232,3),(263,233,3),(266,234,3),(267,235,3),(268,236,3),(269,237,3),(270,238,3),(306,238,4),(271,239,3),(272,240,3),(273,241,3),(274,242,3),(359,242,4),(275,243,3),(276,244,3),(277,245,3),(278,246,3),(279,247,3),(280,248,3),(356,248,4),(281,249,3),(282,250,3),(283,251,3),(379,251,4),(284,252,3),(285,253,3),(286,254,3),(290,255,3),(291,256,3),(293,257,3),(294,258,3),(295,259,3),(296,260,4),(297,261,4),(300,262,4),(301,263,4),(303,264,4),(304,265,4),(305,266,4),(307,267,4),(308,268,4),(309,269,4),(310,270,4),(311,271,4),(312,272,4),(314,273,4),(315,274,4),(317,275,4),(318,276,4),(319,277,4),(321,278,4),(322,279,4),(324,280,4),(325,281,4),(326,282,4),(328,283,4),(329,284,4),(330,285,4),(331,286,4),(332,287,4),(333,288,4),(335,289,4),(336,290,4),(337,291,4),(338,292,4),(342,293,4),(345,294,4),(346,295,4),(347,296,4),(348,297,4),(349,298,4),(350,299,4),(351,300,4),(352,301,4),(353,302,4),(355,303,4),(357,304,4),(358,305,4),(360,306,4),(361,307,4),(363,308,4),(364,309,4),(365,310,4),(367,311,4),(368,312,4),(369,313,4),(370,314,4),(371,315,4),(372,316,4),(373,317,4),(374,318,4),(377,319,4),(378,320,4),(381,321,4),(382,322,4),(383,323,4),(384,324,4),(385,325,4),(386,326,4),(388,327,4),(389,328,4),(390,329,4),(391,330,4),(392,331,4),(393,332,4),(394,333,4),(396,334,4);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `band_concert_record`
--

LOCK TABLES `band_concert_record` WRITE;
/*!40000 ALTER TABLE `band_concert_record` DISABLE KEYS */;
INSERT INTO `band_concert_record` VALUES (1,38,3,'0000-00-00','00:00:00'),(2,39,3,'2014-08-25','14:30:00'),(3,40,3,'2014-03-26','02:18:16');
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
  `start` date DEFAULT NULL,
  `end` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `concert`
--

LOCK TABLES `concert` WRITE;
/*!40000 ALTER TABLE `concert` DISABLE KEYS */;
INSERT INTO `concert` VALUES (2,'Lollapalooza','',1,'http://www.lollapalooza.com/','2014-08-01','2014-08-03'),(3,'Coachella','',2,'http://www.coachella.com/','2014-04-11','0000-00-00'),(4,'Ultra','',3,'http://www.ultramusicfestival.com/','2014-03-28','2014-03-30');
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `state`
--

LOCK TABLES `state` WRITE;
/*!40000 ALTER TABLE `state` DISABLE KEYS */;
INSERT INTO `state` VALUES (1,'Illinois','IL'),(2,'California','CA'),(3,'Florida','FL');
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ticket_record`
--

LOCK TABLES `ticket_record` WRITE;
/*!40000 ALTER TABLE `ticket_record` DISABLE KEYS */;
INSERT INTO `ticket_record` VALUES (1,12.34,3,1,'0000-00-00 00:00:00'),(2,12.34,3,1,'0000-00-00 00:00:00'),(3,0,3,1,'0000-00-00 00:00:00'),(4,0,3,1,'0000-00-00 00:00:00'),(5,0,3,1,'0000-00-00 00:00:00'),(6,0,3,1,'0000-00-00 00:00:00'),(7,0,3,1,'0000-00-00 00:00:00'),(8,0,3,1,'0000-00-00 00:00:00'),(9,0,3,1,'0000-00-00 00:00:00'),(10,0,3,1,'0000-00-00 00:00:00'),(11,0,3,1,'0000-00-00 00:00:00'),(12,0,3,1,'2014-03-31 22:45:14'),(13,123,1,1,'2014-04-07 09:49:26');
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

-- Dump completed on 2014-04-13 12:56:44
