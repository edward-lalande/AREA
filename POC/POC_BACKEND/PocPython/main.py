from flask import Flask, jsonify, request 
import sqlite3

app = Flask(__name__)

def init_db():
    conn = sqlite3.connect('database/user.db')
    cursor = conn.cursor()
    cursor.execute('''
        CREATE TABLE IF NOT EXISTS user (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            mail TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            name TEXT NOT NULL,
            lastname TEXT NOT NULL
        )
    ''')
    conn.commit()
    conn.close()

@app.route('/user', methods=['POST'])
def create_user():
    data = request.json
    mail = data.get('mail')
    password = data.get('password')
    name = data.get('name')
    lastname = data.get('lastname')

    if not (mail and password and name and lastname):
        return jsonify({'error': 'Tous les champs sont requis'}), 400

    try:
        conn = sqlite3.connect('database/user.db')
        cursor = conn.cursor()
        cursor.execute('''
            INSERT INTO user (mail, password, name, lastname)
            VALUES (?, ?, ?, ?)
        ''', (mail, password, name, lastname))
        conn.commit()
        conn.close()
        return jsonify({'message': 'Utilisateur créé avec succès'}), 201
    except sqlite3.IntegrityError:
        return jsonify({'error': 'Un utilisateur avec cet email existe déjà'}), 400

@app.route('/user', methods=['GET'])
def get_users():
    conn = sqlite3.connect('database/user.db')
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM user')
    users = cursor.fetchall()
    conn.close()

    users_list = [{'id': user[0], 'mail': user[1], 'name': user[3], 'lastname': user[4]} for user in users]
    return jsonify(users_list)

@app.route('/user/<int:user_id>', methods=['GET'])
def get_user(user_id):
    conn = sqlite3.connect('database/user.db')
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM user WHERE id = ?', (user_id,))
    user = cursor.fetchone()
    conn.close()

    if user:
        user_data = {'id': user[0], 'mail': user[1], 'name': user[3], 'lastname': user[4]}
        return jsonify(user_data)
    else:
        return jsonify({'error': 'Utilisateur non trouvé'}), 404

@app.route('/user/<int:user_id>', methods=['PUT'])
def update_user(user_id):
    data = request.json
    mail = data.get('mail')
    password = data.get('password')
    name = data.get('name')
    lastname = data.get('lastname')

    if not (mail and password and name and lastname):
        return jsonify({'error': 'Tous les champs sont requis'}), 400

    conn = sqlite3.connect('database/user.db')
    cursor = conn.cursor()
    cursor.execute('''
        UPDATE user
        SET mail = ?, password = ?, name = ?, lastname = ?
        WHERE id = ?
    ''', (mail, password, name, lastname, user_id))
    conn.commit()
    updated_rows = cursor.rowcount
    conn.close()

    if updated_rows > 0:
        return jsonify({'message': 'Utilisateur mis à jour avec succès'})
    else:
        return jsonify({'error': 'Utilisateur non trouvé'}), 404

@app.route('/user/<int:user_id>', methods=['DELETE'])
def delete_user(user_id):
    conn = sqlite3.connect('database/user.db')
    cursor = conn.cursor()
    cursor.execute('DELETE FROM user WHERE id = ?', (user_id,))
    conn.commit()
    deleted_rows = cursor.rowcount
    conn.close()

    if deleted_rows > 0:
        return jsonify({'message': 'Utilisateur supprimé avec succès'})
    else:
        return jsonify({'error': 'Utilisateur non trouvé'}), 404


@app.route('/', methods = ['GET', 'POST'])
def home(): 
	if(request.method == 'GET'):

		data = "hello world"
		return jsonify({'data': data}) 


@app.route('/home/<int:num>', methods = ['GET'])
def disp(num): 

	return jsonify({'data': num**2})

if __name__ == '__main__':
	app.run(debug = True)
