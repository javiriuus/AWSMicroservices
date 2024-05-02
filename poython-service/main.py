from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from marshmallow import Schema, fields

app = Flask(__name__)

# Configura la conexión a PostgreSQL
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://username:password@hostname:5432/database_name'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

# Define un modelo de base de datos para la tabla "products"
class Product(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    price = db.Column(db.Float, nullable=False)

# Define un esquema de serialización para Product
class ProductSchema(Schema):
    id = fields.Int()
    name = fields.Str()
    price = fields.Float()

product_schema = ProductSchema()
products_schema = ProductSchema(many=True)

# Ruta para obtener la lista de productos
@app.route('/products', methods=['GET'])
def get_products():
    products = Product.query.all()
    return products_schema.jsonify(products)

# Inicia el servidor HTTP
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
