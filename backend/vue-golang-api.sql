-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: vue-golang-api
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

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
-- Table structure for table `best_produks`
--

DROP TABLE IF EXISTS `best_produks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `best_produks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama` char(100) DEFAULT NULL,
  `harga` bigint unsigned DEFAULT NULL,
  `is_ready` tinyint(1) DEFAULT '1',
  `gambar` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gambar` (`gambar`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `best_produks`
--

LOCK TABLES `best_produks` WRITE;
/*!40000 ALTER TABLE `best_produks` DISABLE KEYS */;
INSERT INTO `best_produks` VALUES (1,'Nasi Ayam Geprek',15000,1,'nasi-ayam-geprek.jpg'),(2,'Nasi Rames',12000,1,'nasi-rames.jpg'),(3,'Sate Ayam',20000,1,'sate-ayam.jpg');
/*!40000 ALTER TABLE `best_produks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keranjangs`
--

DROP TABLE IF EXISTS `keranjangs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `keranjangs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nama` char(100) DEFAULT NULL,
  `harga` bigint DEFAULT NULL,
  `jumlah` bigint DEFAULT NULL,
  `deskripsi` longtext,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_keranjangs` (`user_id`),
  CONSTRAINT `fk_users_keranjangs` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keranjangs`
--

LOCK TABLES `keranjangs` WRITE;
/*!40000 ALTER TABLE `keranjangs` DISABLE KEYS */;
/*!40000 ALTER TABLE `keranjangs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `produks`
--

DROP TABLE IF EXISTS `produks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `produks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama` char(100) DEFAULT NULL,
  `harga` bigint unsigned DEFAULT NULL,
  `is_ready` tinyint(1) DEFAULT '1',
  `gambar` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gambar` (`gambar`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `produks`
--

LOCK TABLES `produks` WRITE;
/*!40000 ALTER TABLE `produks` DISABLE KEYS */;
/*!40000 ALTER TABLE `produks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` longtext NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','admin@gmail.com','$2a$04$xg4Glp.trmhqEJ3dobQXl.CrVi1p5UkaJMNs3miFTGc1uIozlJWYO');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-09 11:20:22
