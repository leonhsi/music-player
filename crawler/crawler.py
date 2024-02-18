import csv
import urllib
import json
import requests
import ssl
import string
from pytube import YouTube
from tqdm import tqdm

ssl._create_default_https_context = ssl._create_stdlib_context

class Song:
    def __init__(self, name, creator, url):
        self.name = name
        self.creator = creator
        self.url = url

def read_csv_file(file_path):
    data = []
    with open(file_path, newline='') as csvfile:
        csv_reader = csv.reader(csvfile)
        for row in csv_reader:
            data.append(row)
    return data

def process_track_name(track_name):
    track_name = track_name.replace('+', ' ')
    track_name = urllib.parse.unquote(track_name)
    track_name = track_name.replace('/_/', '/')

    track_list = track_name.split('/')

    for idx in range(len(track_list)):
        track_list[idx] = track_list[idx].translate(str.maketrans('', '', string.punctuation))

    return track_list

def read_csv_data(csv_data, limit):
    songs = []

    for idx, _ in zip(range(limit), csv_data):
        creator, name = process_track_name(csv_data[idx + 1][1])
        url = csv_data[idx + 1][10]

        song = Song(name, creator, url)
        songs.append(song)
        
    return songs 

def song_to_dict(song):
    return {
        "name": song.name,
        "creator": song.creator,
        "url": song.url,
        "music_path": './audios/' + song.name.replace(' ', '-') + ".mp3", 
        "cover_path": './covers/' + song.name.replace(' ', '-') + ".jpeg" 
    }

if __name__ == "__main__":
    file_path = 'tracks.csv'
    mp3_download_path = '../frontend/src/audios/'
    img_download_path = '../frontend/src/covers/'
    metadata_path = '../backend/'
    track_limit = 70
    
    # parse data from csv file
    csv_data = read_csv_file(file_path)
    song_data = read_csv_data(csv_data, track_limit)

    # for idx in range(track_limit):
    #     print(idx, ":", song_data[idx].name, "/", song_data[idx].creator, "/", song_data[idx].url)

    # dump metadata to json file
    song_dict = [song_to_dict(song) for song in song_data]
    with open(metadata_path + 'metadata.json', 'w') as json_file:
        json.dump(song_dict, json_file, indent=4)

    # download mp3 file and cover image from YouTube
    progress = tqdm(total = track_limit)

    for song in song_data:
        yt = YouTube(song.url)

        # mp3
        yt.streams.filter().get_audio_only().download(mp3_download_path, song.name.replace(' ', '-') + ".mp3")

        # image
        resp = requests.get(yt.thumbnail_url)

        if resp.status_code == 200:
            extesion = resp.headers['Content-Type'].split('/')[1]
            with open(img_download_path + song.name.replace(' ', '-') + "." + extesion, 'wb') as img_file:
                img_file.write(resp.content)
        else:
            print("Failed to download thumbnail for", song.name, "by", song.creator)

        progress.update(1)
