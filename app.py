from flask import Flask
from flask_pymongo import PyMongo
from random import choice

app = Flask(name)
app.config["MONGO_URI"] = "mongodb://mongo1:27017,stajdb"
mongo = PyMongo(app)

@app.route('/')
def hello_world():
    return 'Merhaba Python!'

@app.route('/staj')
def random_city():
    cities = mongo.db.iller.find()
    random_city = choice(list(cities))
    return str(random_city)

if name == 'main':
    app.run(host='0.0.0.0', port=4444)