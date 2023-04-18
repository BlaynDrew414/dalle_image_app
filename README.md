DALL-E Image App
This app uses the DALL-E neural network with a Go backend and Mongo database to generate and store images from text descriptions. It has a web interface where users can input a text description, and the DALL-E model will generate an image based on that description. 

Installation
To install and run the app, follow these steps:

Clone this repository to your local machine.
Install the required go packages by running the following command in your terminal:

Copy code and import all necessary packages. 

Ensure the directoy is set to /backend.

Start the web server by running the following command in your terminal:

go run main.go

Open a web browser and navigate to http://localhost:8080 to access the app. Generate and view images from their respective Get and Post request handlers, utilizing an API testing tool, such as Postman.

Usage

 The DALL-E model will generate an image based on the text description, which will be displayed on the screen.

Limitations
Please note that the DALL-E model is still experimental and may not always generate accurate or relevant images. Additionally, the app's performance may be affected by the computational resources available on the machine running the app.
Also note that this app was made with a free trial subscription to both OpenAI and Mongodb Atlas, which may have set data limits that may or may not be maxed. It may be necessary to create your own api key and database URI.

Credits
This app was created by Blayn Drew and is based on the DALL-E model by OpenAI.

License
This project is licensed under the MIT License - see the LICENSE file for details.
