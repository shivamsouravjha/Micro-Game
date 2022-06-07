# Micro-Game
An online self-publishingportal to letusers read thecontent. Theyare meant tohelp readers beconsistent inreading anddevelop aregularuserbase.Chapters areunlocked dailyper series fornew users.New User andNew Contentare managedto make surefree 4 chaptersforeveryone(everychapter addedis madeavailable toeveryone).Independent Micro-Servicesmade to inter-connectServices.MetaData isuploaded inbulk andstored inRelationalDatabase. Thisalso has thefunction ofunlocking achapter everyday for theuser(regularsigning meansone chapterunlockeddaily).

## Explaining the Architecture 
This application contains three services,
### User Service
- Service description : Contains user details like name, email etc
- Apis :
  * Fetch users
  * Create a new user

### Content Service
- Service description : Contains meta of readable content that we serve to the users
- Each content is structured as a series which can contain multiple chapters.
- Example: consider a series called “Harry Potter” which has 7 chapters
- Apis required :
  * Fetch content for a user
  * Input : userid , multiple series ids
  * Output : content meta with only unlocked chapters per series
  * Api for bulk upload of the content

### Daily Pass
- Service description : Contains the details of how many chapters per series is unlocked for a
particular user
  * When a user installs the Pratilipi Application (i.e. day1 of user creation) , 4 chapters are
unlocked by default on the day of installation
  * Rest of the chapters are released on a daily basis; that is one chapter per day
  * Any new content that is uploaded, again by default has 4 unlocked chapters for all the
users

- So if a series is uploaded after user creation then the existing users should also see that
series with only 4 chapters unlocked

- Apis
  * Api to unlock one chapters for the given user and series
  * For testing purpose : this api should not be idempotent, if I hit the api twice, it
should unlock two episodes for the requested user and series
  * Number of unlocked chapters = 4 + no. of times unlock api is hit
  * Fetch daily pass data by user
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

## How did I deploy things?
* [Heroku(to deploy three independent microservices)](https://dashboard.heroku.com/apps)
* [RabbitMQ](https://lionfish.rmq.cloudamqp.com/)
* [SQL Server](https://www.freesqldatabase.com/)

## Overall Schematic 
![Group Schema (5)](https://user-images.githubusercontent.com/60891544/161864560-2e77405a-282d-47ed-9764-9808f189b6a0.png)

## User Service Schematic
![System Flow (2)](https://user-images.githubusercontent.com/60891544/161865343-a28fd6e8-391a-44d1-83c9-d7b3c2afa5cc.png)

## Content Service Schematic
![System Flow (11)](https://user-images.githubusercontent.com/60891544/161868158-7b49b558-a787-4bf1-8ec5-4a69abb12bd9.png)

## DailyPass Service Schematic
![System Flow (10)](https://user-images.githubusercontent.com/60891544/161868089-8a9ad952-acd4-4d8f-b93e-986ec9ee88f1.png)
