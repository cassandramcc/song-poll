import spotipy
from spotipy.oauth2 import SpotifyClientCredentials

import config.keys as keys
from datetime import datetime

sp = spotipy.Spotify(auth_manager=SpotifyClientCredentials(client_id=keys.clientID,
                                                           client_secret=keys.clientSecret))

class Album:
    def __init__(self, name, releaseDate):
        self.name = name
        self.releaseDate = releaseDate

    def __str__(self):
        return f"{self.name}-{self.releaseDate}"

def getArtistID(artist):
    searchResult = sp.search(artist,1,type='artist')
    id = searchResult['artists']['items'][0]['id']
    return id

def getArtistAlbums(id):
    moreAlbums = True
    offset = 0
    albums = []
    while moreAlbums:
        albumsResult = sp.artist_albums(id,limit=50,offset=offset)
        if len(albumsResult['items']) == 0:
            break
        for _, album in enumerate(albumsResult['items']):
            albums.append(Album(album['name'],album['release_date']))
        offset += 50
    return albums

def sortAlbums(albums):
    sortedAlbums = sorted(albums, key=lambda a: a.releaseDate, reverse=True)
    return sortedAlbums

def getAlbumsAfterVisitDate(albums,visitDate):
    newAlbums = []
    for album in albums:
        if datetime.strptime(album.releaseDate,'%Y-%m-%d') >= visitDate:
            newAlbums.append(album)
    return newAlbums

artistID = getArtistID("IID")
albums =  getArtistAlbums("3GZ8Rfap7VxAOzABiZEXCL")
newAlbums = getAlbumsAfterVisitDate(albums,datetime.strptime('2023-01-01','%Y-%m-%d'))
for a in newAlbums:
    print(a)    