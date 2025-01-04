CREATE DATABASE IF NOT EXISTS `WebPractice1`;
CREATE TABLE IF NOT EXISTS "AbuseEntity"(
    "IPAddress" VARCHAR(15),            
    "IsPublic" BOOLEAN,                 
    "IPVersion" INT,                    
    "IsWhitelisted" BOOLEAN,            
    "AbuseConfidenceScore" INT,        
    "CountryCode" VARCHAR(5),             
    "CountryName" VARCHAR(100),         
    "UsageType" VARCHAR(100),           
    "ISP" VARCHAR(200),                 
    "IsTor" BOOLEAN,                    
    "IsFromDB" BOOLEAN 
);