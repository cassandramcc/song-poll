import spotipy
from spotipy.oauth2 import SpotifyClientCredentials
from spotipy.oauth2 import SpotifyOAuth
from music import Album, Artist, Track

import config.keys as keys
from datetime import datetime

class SpotifyAPI():

    #sp = spotipy.Spotify(auth_manager=SpotifyClientCredentials(client_id=keys.clientID,client_secret=keys.clientSecret))
    sp = spotipy.Spotify(auth_manager=SpotifyOAuth(client_id=keys.clientID, client_secret=keys.clientSecret, redirect_uri='http://localhost:5000/callback', scope='user-library-read playlist-modify-private playlist-modify-public'))

    def getAuthManager(self,auth_manager):
        auth = spotipy.Spotify(auth_manager=auth_manager)
        return auth

    def getArtistID(self, artist):
        searchResult = self.sp.search(artist,1,type='artist')
        id = searchResult['artists']['items'][0]['id']
        return id

    def getArtist(self, id):
        artistResult =  self.sp.artist(id)
        return Artist(artistResult['name'],artistResult['id'],None)

    def getArtistAlbums(self, id):
        offset = 0
        albums = []
        while True:
            albumsResult =  self.sp.artist_albums(id,limit=50,offset=offset)
            if len(albumsResult['items']) == 0:
                break
            for _, album in enumerate(albumsResult['items']):
                albums.append(Album(album['name'],[self.getArtist(id)],album['release_date'],album['id']))
            offset += 50
        return albums

    def sortAlbums(albums):
        sortedAlbums = sorted(albums, key=lambda a: a.releaseDate, reverse=True)
        return sortedAlbums

    def getAlbumTracks(self, album):
        tracks = []
        result = self.sp.album_tracks(album.id)
        for track in result['items']:
            artists = track['artists']
            artistsObj = []
            for artist in artists:
                artistsObj.append(Artist(artist['name'],artist['id'],'2023-01-01'))
            tracks.append(Track(track['name'],artistsObj,track['id']))
        return tracks

    def getAlbumsAfterVisitDate(albums,visitDate):
        newAlbums = []
        for album in albums:
            try:
                if datetime.strptime(album.releaseDate,'%Y-%m-%d') >= datetime.strptime(visitDate,'%Y-%m-%d'):
                    newAlbums.append(album)
            except:
                pass
        return newAlbums

    def getPlaylistTracks(self, id):
        offset = 0
        tracks = []
        while True:
            playlistResult = self.sp.playlist_items(id, limit=100,offset=offset)
            if len(playlistResult['items']) == 0:
                break
            for _, track in enumerate(playlistResult['items']):
                artists = track['track']['artists']
                artistsObj = []
                for artist in artists:
                    artistsObj.append(Artist(artist['name'],artist['id'],'2023-01-01'))
                tracks.append(Track(track['track']['name'],artistsObj, track['track']['id']))
            offset += 100
        return tracks

    def getArtistAlbumsAfterVisitDate(self,id,visitDate):
        albums =  self.getArtistAlbums(id)
        newAlbums = self.getAlbumsAfterVisitDate(albums,visitDate)
        return newAlbums    

    def artistExists(artist, allArtists):
        for a in allArtists:
            if a.name == artist.name:
                return True
        return False
