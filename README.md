
### launch mysql
`docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql `

Use mysql client to connect to mysql, then create database publicX
`mysql --port=3306 -h172.17.0.1 -uroot -p123456` 

`create database publicX;`

The tables will be created after the CampaignMsgs start, detail schemas is in pkg/dao/

### build

`go build`

### create campaign

`./publicX campaignMgrOrm`

It will create a campaign and import recipients from ./data.csv,
the schedual time is a fixed value of time.Now()+30 seconds.
