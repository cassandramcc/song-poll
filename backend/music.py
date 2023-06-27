from typing import List

class Artist:
    def __init__(self, name, id, lastVisited):
        self.name = name
        self.id = id
        self.lastVisited = lastVisited

    def __str__(self):
        return f"{self.name} - {self.id}"


class Album:
    def __init__(self, name, artists: List[Artist], releaseDate, id):
        self.name = name
        self.artists = artists
        self.releaseDate = releaseDate
        self.id = id

    def __str__(self):
        artists = [artist.name for artist in self.artists]
        return f"{self.name} - {artists} - {self.releaseDate}"

    
class Track:
    def __init__(self, name, artists: List[Artist], id):
        self.name = name
        self.artists = artists
        self.id = id

    def __str__(self):
        artists = [artist.name for artist in self.artists]
        return f"{self.name} - {artists} - {self.id}"