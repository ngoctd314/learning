import requests

url = "http://localhost:8000/snippets/"


data = {
    "title": "title1",
    "code": "123",
    "linenos ": True,
}
# response = requests.post(url, json=data)

response = requests.get(url)

response_json = response.json()
print(response_json)
