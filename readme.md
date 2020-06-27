

### deploy to heroku 

   1) heroku conatiner:login 
   2) navigate to project root where Docker file exists 
    
     "heroku create" 
     which will create us a app name to deploy , lets say abc-abc-12345 is app name

   3) now , just push to heroku registry using 
       heroku container:push web --app <appname>

   4) final step top verify app , 
   
       heroku open --app <appname>