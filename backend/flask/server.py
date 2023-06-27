from flask import Flask, redirect, request
import spotipy
from spotify import SpotifyAPI
from spotipy.oauth2 import SpotifyOAuth
import config.keys as keys


app = Flask(__name__)

@app.route('/')
def index():
    sp = SpotifyAPI()
    
    
    return response

@app.route('/callback')
def callback():
    sp = SpotifyAPI()
    auth_manager = sp.getAuthManager()
    code = request.args.get('code')
    token_info = auth_manager.get_access_token(code)
    sp = spotipy.Spotify(auth_manager=auth_manager)
    # do something with the Spotify API
    return "Logged in successfully"

if __name__ == '__main__':
    app.run(debug=True)