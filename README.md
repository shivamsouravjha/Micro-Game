# Micro-Game
So how does things work?
* Import Postman Requests [Curl](https://www.getpostman.com/collections/b406514faba24cd2fd1e)
(Import using JSON link)
## Explaining APIs:-
* Get Username
  - Input:Username
  - Output:userId,firstName,lastName,penName,userEmail,bio,number
* Create User
  - Input:userId,firstName,lastName,penName,userEmail,bio,number,password
  - Output:token
* Login User
  - Input:penName,password
  - Output:token
* Fetch Unlocked Chapters
  - Input:UserID,SeriesID
  - Output:[]{seriesId,title,story}
* Bulk Upload Chapters
  - Input:[]{seriesId,title,story}
  - Output:Success Response
* Get Series
  - Input:seriesId
  - Output:seriesId,author,name,chapters
* Create Series
  - Input:author,name,chapters
  - Output:Success Response
* Unlock Chapters
  - Input:UserID,SeriesID
  - Output:Success Response
* Get Unlock Chapters
  - Input:UserID
  - Output:Chapter and Unlocked List for given user

## Inter-Connected Microservices:-
* CreateUser
  - New User is created and userID is sent to get all available chapter.
* UserSeries
  - All available chapters with userID is stored in DailyPass Service Database.
* SeriesContent
  - New Series is added,it's added to all existing users.

## Things I used:-
* Go(lang)
* SQL
* RabbitMQ(to connect microservices)
* JWT(for unlocking chapters)

## Overall Schematic 
![Group Schema (5)](https://user-images.githubusercontent.com/60891544/161864560-2e77405a-282d-47ed-9764-9808f189b6a0.png)

## User Service Schematic
![System Flow (2)](https://user-images.githubusercontent.com/60891544/161865343-a28fd6e8-391a-44d1-83c9-d7b3c2afa5cc.png)

## Content Service Schematic
![System Flow (11)](https://user-images.githubusercontent.com/60891544/161868158-7b49b558-a787-4bf1-8ec5-4a69abb12bd9.png)

## DailyPass Service Schematic
![System Flow (10)](https://user-images.githubusercontent.com/60891544/161868089-8a9ad952-acd4-4d8f-b93e-986ec9ee88f1.png)
