# Namefinder
A REST service to find about dog name preferences
## TODO
1. Properly set up REST service
 - Router
 - Decouple Routes from Router from Logic
2. Calculate Association Rules based on Transaction DB
 - Randomly generate transaction DB
 - Refactor AssociationRules into own file
 - Use goroutines
3. Refactor App into Packages and Comment
 - Decouple Nameserver from rest, Server as lib?
    - Start Server own App
    - Fill DB own app
 - Hard coded URLS and DB names..
4. Allow multiple association items
5. Decouple REST API from Webservice
 - Web Interface?
 - Curl Interface



