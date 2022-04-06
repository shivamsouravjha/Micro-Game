# Micro-Game
So how does things work?
* Import Postman Requests [Curl](https://go.postman.co/workspace/My-Workspace~6a5d2256-e31b-4175-8e51-f49f4bdb27d1/collection/14182772-67a55a70-150d-4fa3-b93b-84c479113a1b?action=share&creator=14182772)

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
