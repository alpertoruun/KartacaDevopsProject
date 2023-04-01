#!/usr/bin/env bash

echo "Sleeping while waiting for MongoDB to start…"
sleep 20

echo "Setting MongoDB replica set…"
mongo --host mongo1 --port 27017 --eval 'rs.initiate({_id: “staj”, members: [{_id: 0, host: “mongo1:27017”}, {_id: 1, host: “mongo2:27017”}, {_id: 2, host: “mongo3:27017”}]})'

echo "Creating stajdb database…"
mongo --host mongo1 --port 27017 --authenticationDatabase admin --username admin --password admin --eval 'db.adminCommand({create: “stajdb”})'

echo "Creating collection 'iller' and inserting data…"
mongo --host mongo1 --port 27017 --authenticationDatabase admin --username admin --password admin stajdb --eval 'db.iller.insertMany(["Malatya","Adiyaman","Urfa","Maraş","Elazig","Ankara","Kayseri","Van","Bursa","Bingol"])'

echo "Creating collection 'ulkeler' and inserting data…"
mongo --host mongo1 --port 27017 --authenticationDatabase admin --username admin --password admin stajdb --eval 'db.ulkeler.insertMany(["Türkiye","ABD","Ingiltere","Danimarka","Misir","Irak","Iran","Almanya","Norveç",])'