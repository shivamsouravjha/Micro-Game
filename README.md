# Micro-Game
So how does things work?
* Import Postman Requests [Curl](https://www.getpostman.com/collections/b406514faba24cd2fd1e)
(Import using JSON link)
## Explaining APIs:-
* Get Username
  - Input:Username
  - Output:userId,firstName,lastName,penName,userEmail,bio,number
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/getUser/:username
* Create User
  - Input:userId,firstName,lastName,penName,userEmail,bio,number,password
  - Output:token
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/createUser
* Login User
  - Input:penName,password
  - Output:token
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/loginUser
* Fetch Unlocked Chapters
  - Input:UserID,SeriesID
  - Output:[]{seriesId,title,story}
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/fetchContent/:userID/:seriesID
* Bulk Upload Chapters
  - Input:[]{seriesId,title,story}
  - Output:Success Response
  - Ssample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/uploadContent
* Get Series
  - Input:seriesId
  - Output:seriesId,author,name,chapters
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/getSeries/:seriesID
* Create Series
  - Input:author,name,chapters
  - Output:Success Response
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/uploadSeries
* Unlock Chapters
  - Input:UserID,SeriesID
  - Output:Success Response
  - Sample URL:-https://dailypassserviceshivam.herokuapp.com/api/v0/dailypass/unlockContent
* Get Unlock Chapters
  - Input:UserID
  - Output:Chapter and Unlocked List for given user
  - Sample URL:-https://dailypassserviceshivam.herokuapp.com/api/v0/dailypass/getUnlockedContent/:userID

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
