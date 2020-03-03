# GoLangSampleAPI

To use this code make sure you have your golang environment, file path, etc configured
Also I recommend using the RESTClient browser extension for implementing 

The comments referring to Dbs are what the code would look like if I hadn't used mock data

Things to improve:
-I'd probably change the data types to something that could be queried since dbs wouldnt 
have the dates as strings for example
-Create Mongo database with actual collections to take it a step further

Things I like:
-Full CRUD
-Topic is more interesting (rapper)
-I have all albums in there even as mock data
___________________________________________________________________________

Updated file is a GraphQL version of the early RESTful code that is the original upload in this repo.
I converted the main chunk of that code to work with GraphQL in order to better query the information as opposed to have to hit multiple endpoints in the RESTful version. I know it's not extremely complex but the purpose of this exercise was to get my hands dirty and mess around with GraphQL
